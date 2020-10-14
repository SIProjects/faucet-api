package model

import "time"

type Payout struct {
	TxID       string    `json:"txid"`
	Amount     float64   `json:"amount"`
	Address    string    `json:"address"`
	IsMined    bool      `json:"is_mined"`
	InsertedAt time.Time `json:"inserted_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
