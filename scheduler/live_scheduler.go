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
	Interval  time.Duration
	Chain     *chain.Chain
	Cache     cache.Cache
	DB        database.Database
	Logger    *log.Logger
	MinPayout int
	MaxPayout int
}

func New(
	interval time.Duration,
	db database.Database,
	c cache.Cache,
	ch *chain.Chain,
	l *log.Logger,
	min int,
	max int,
) (*LiveScheduler, error) {

	l.Println(
		"Scheduler set up to payout between", min, "-", max, "every",
		interval.String(),
	)

	return &LiveScheduler{
		Interval:  interval,
		Chain:     ch,
		DB:        db,
		Cache:     c,
		Logger:    l,
		MinPayout: min,
		MaxPayout: max,
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
			go s.CreatePayouts()
			go s.CheckMined()
		}
	}
}

func (s *LiveScheduler) CreatePayouts() {
	addresses, err := s.Cache.GetNextAddresses(20)
	if err != nil {
		s.Logger.Println("Obtain Cache Error:", err.Error())
		return
	}

	if len(addresses) == 0 {
		return
	}

	payments := make([]node.Payment, 0)

	for _, x := range addresses {
		amount, err := s.randomAmount()
		if err != nil {
			s.Logger.Println("Random Amount Error:", err.Error())
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
			continue
		}

		s.Logger.Println("Paying out", p.Amount, "to", p.Address)
	}

	s.Logger.Println("Paid out:", txid)

}

func (s *LiveScheduler) randomAmount() (btcutil.Amount, error) {
	val := rand.Intn(s.MaxPayout-s.MinPayout) + s.MinPayout
	return btcutil.NewAmount(float64(val))
}

func (s *LiveScheduler) CheckMined() {
	now := time.Now()
	txids, err := payout.UnminedTxIDs(s.DB)

	if err != nil {
		s.Logger.Println("Error:", err.Error())
		return
	}

	for _, txid := range txids {
		txn, err := s.Chain.Node.GetTransaction(txid)

		if err != nil {
			s.Logger.Println("Error:", err.Error())
			continue
		}

		if txn.Confirmations == 0 {
			continue
		}

		err = payout.SetMined(txid, now, s.DB)

		if err != nil {
			s.Logger.Println("Error:", err.Error())
			continue
		}

		s.Logger.Println("Setting mined", txid)
	}
}
