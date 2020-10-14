package payout

import (
	"context"
	"time"

	"github.com/SIProjects/faucet-api/database"
	"github.com/SIProjects/faucet-api/model"
)

func Recent(db database.Database) (res []model.Payout, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `
	SELECT txid, address, amount, is_mined, inserted_at, updated_at
	FROM payout ORDER BY updated_at DESC LIMIT(50)
	`
	rows, err := db.Query(ctx, query)

	if err != nil {
		return res, err
	}

	res = make([]model.Payout, 0)

	for rows.Next() {
		var payout model.Payout
		rows.Scan(
			&payout.TxID, &payout.Address, &payout.Amount,
			&payout.IsMined, &payout.InsertedAt, &payout.UpdatedAt,
		)
		res = append(res, payout)
	}

	return res, rows.Err()
}

func Insert(payout model.Payout, db database.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `
	INSERT INTO payout (
		txid, address, amount, is_mined, inserted_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6)
	`

	err := db.Exec(
		ctx, query, payout.TxID, payout.Address, payout.Amount, payout.IsMined,
		payout.InsertedAt, payout.UpdatedAt,
	)

	return err
}
