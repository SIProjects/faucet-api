package cache

import (
	"context"
	"time"

	"github.com/btcsuite/btcutil"
	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	Name   string
	Client *redis.Client
}

func New(url, password, dbname string) (*RedisCache, error) {
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
	pending, err := c.Client.HExists(ctx, c.Key("pending"), addr.String()).Result()

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

	return paid != 0, nil
}

func (c *RedisCache) AddAddressToQueue(addr btcutil.Address) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := c.Client.HSet(ctx, c.Key("pending"), addr.String()).Result()

	if err != nil {
		return err
	}

	return nil
}

func (c *RedisCache) AddAddressPayout() error {
	return nil
}
