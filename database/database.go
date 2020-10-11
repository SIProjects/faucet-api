package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
)

type Executable interface {
	Exec(ctx context.Context, sql string, arg ...interface{}) error
}

type QueryRowable interface {
	QueryRow(
		ctx context.Context, sql string, args ...interface{},
	) interface{ Scannable }
}

type Scannable interface {
	Scan(dest ...interface{}) error
}

type Database interface {
	Executable
	QueryRowable
	Close()
}

type PGDatabase struct {
	Conn *pgx.Conn
}

func New(url string) (*PGDatabase, error) {
	conn, err := pgx.Connect(context.Background(), url)

	if err != nil {
		return nil, err
	}

	db := PGDatabase{
		Conn: conn,
	}

	err = db.Conn.Ping(context.Background())

	if err != nil {
		return nil, err
	}

	err = migrateDatabase(url)

	if err != nil {
		return nil, err
	}

	return &db, nil
}

func (d *PGDatabase) Exec(
	ctx context.Context, sql string, args ...interface{},
) error {
	_, err := d.Conn.Exec(ctx, sql, args...)
	return err
}

func (d *PGDatabase) QueryRow(
	ctx context.Context, sql string, args ...interface{},
) interface{ Scannable } {
	row := d.Conn.QueryRow(ctx, sql, args...)
	return row
}

func (d *PGDatabase) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	d.Conn.Close(ctx)
}
