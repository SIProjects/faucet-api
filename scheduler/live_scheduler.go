package scheduler

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/SIProjects/faucet-api/cache"
	"github.com/SIProjects/faucet-api/chain"
	"github.com/SIProjects/faucet-api/database"
	"github.com/SIProjects/faucet-api/model"
	"github.com/SIProjects/faucet-api/node"
	"github.com/SIProjects/faucet-api/resources/payout"
	"github.com/btcsuite/btcutil"
)

type LiveScheduler struct {
	Interval time.Duration
	Chain    *chain.Chain
	Cache    cache.Cache
	DB       database.Database
}

func New(
	interval time.Duration,
	db database.Database,
	c cache.Cache,
	ch *chain.Chain,
) (*LiveScheduler, error) {
	return &LiveScheduler{
		Interval: interval,
		Chain:    ch,
		DB:       db,
		Cache:    c,
	}, nil
}

func (s *LiveScheduler) Start(done chan struct{}) {
	timer := time.NewTicker(s.Interval)
	defer timer.Stop()

	for {
		select {
		case <-done:
			return
		case <-timer.C:
			s.CreatePayouts()
		}
	}
}

func (s *LiveScheduler) CreatePayouts() {
	addresses, err := s.Cache.GetNextAddresses(20)
	if err != nil {
		log.Println("Obtain Cache Error:", err.Error())
		return
	}

	if len(addresses) == 0 {
		return
	}

	payments := make([]node.Payment, 0)

	for _, x := range addresses {
		amount, err := randomAmount()
		if err != nil {
			log.Println("Random Amount Error:", err.Error())
			return
		}
		payments = append(payments, node.Payment{
			Address: x,
			Amount:  amount,
		})
	}

	txid, amounts, err := s.Chain.Node.PayToAddresses(payments)

	if err != nil {
		log.Println("RPC Error:", err.Error())
		return
	}

	fmt.Println("Paid out:", txid)

	now := time.Now()
	for address, amount := range amounts {
		p := model.Payout{
			Amount:     amount.ToUnit(btcutil.AmountBTC),
			Address:    address.String(),
			TxID:       txid,
			InsertedAt: now,
			UpdatedAt:  now,
		}

		err := payout.Insert(p, s.DB)
		if err != nil {
			fmt.Println("Error inserting:", err.Error())
		}
	}
}

func randomAmount() (btcutil.Amount, error) {
	val := rand.Int63n(100-10) + 10
	return btcutil.NewAmount(float64(val))
}
