package cache

import (
	"context"
	"time"

	"github.com/SIProjects/faucet-api/chain"
	"github.com/btcsuite/btcutil"
	"github.com/go-redis/redis/v8"
)

const (
	QUEUE_KEY   = "pending:queue"
	PENDING_KEY = "pending"
)

type RedisCache struct {
	Name   string
	Client *redis.Client
	Chain  *chain.Chain
}

func New(url, password, dbname string, ch *chain.Chain) (*RedisCache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		return nil, err
	}

	c := RedisCache{
		Client: rdb,
		Name:   dbname,
		Chain:  ch,
	}

	return &c, nil
}

func (c *RedisCache) Close() {
	c.Client.Close()
}

func (c *RedisCache) Key(key string) string {
	return c.Name + ":" + key
}

func (c *RedisCache) AddressKey(addr btcutil.Address) string {
	return c.Name + ":paid:" + addr.String()
}

func (c *RedisCache) CanAddAddress(addr btcutil.Address) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	pending, err := c.Client.SIsMember(
		ctx, c.Key(PENDING_KEY), addr.String(),
	).Result()

	if err != nil {
		return false, err
	}

	if pending {
		return false, nil
	}

	paid, err := c.Client.Exists(ctx, c.AddressKey(addr)).Result()

	if err != nil {
		return false, err
	}

	return paid == 0, nil
}

func (c *RedisCache) AddAddressToQueue(addr btcutil.Address) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := c.Client.TxPipelined(ctx, func(p redis.Pipeliner) error {
		_, err := p.SAdd(ctx, c.Key(PENDING_KEY), addr.String()).Result()

		if err != nil {
			return err
		}

		_, err = p.RPush(ctx, c.Key(QUEUE_KEY), addr.String()).Result()

		if err != nil {
			return err
		}

		_, err = p.Set(ctx, c.AddressKey(addr), 1, time.Hour*24).Result()

		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (c *RedisCache) AddAddressPayout() error {
	return nil
}

func (c *RedisCache) GetNextAddresses(num int) ([]btcutil.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := make([]btcutil.Address, 0)
	p := c.Client

	for i := 0; i < num; i++ {
		exists, err := c.Client.Exists(ctx, c.Key(QUEUE_KEY)).Result()

		if err != nil {
			return res, err
		}

		if exists == 0 {
			return res, nil
		}

		addrStr, err := p.LPop(ctx, c.Key(QUEUE_KEY)).Result()

		if err != nil {
			return []btcutil.Address{}, err
		}

		if addrStr == "" {
			continue
		}

		addr, err := c.Chain.DecodeAddress(addrStr)

		if err != nil {
			return []btcutil.Address{}, err
		}

		_, err = p.SRem(ctx, c.Key(PENDING_KEY), addrStr).Result()

		if err != nil {
			return []btcutil.Address{}, err
		}

		res = append(res, addr)
	}

	return res, nil
}

func (c *RedisCache) GetQueuedCount() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return c.Client.LLen(ctx, c.Key(QUEUE_KEY)).Result()
}
