package model

import "time"

type Payout struct {
	TxID       string
	Amount     uint64
	Address    string
	InsertedAt time.Time
	UpdatedAt  time.Time
}
