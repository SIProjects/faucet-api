package model

import "time"

type Payout struct {
	TxID       string    `json:"tx_id,omitempty"`
	Amount     float64   `json:"amount,omitempty"`
	Address    string    `json:"address,omitempty"`
	IsMined    bool      `json:"is_mined,omitempty"`
	InsertedAt time.Time `json:"inserted_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
