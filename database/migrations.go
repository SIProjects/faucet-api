package database

import (
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func migrateDatabase(url string) error {
	dir := os.Getenv("CWD")

	if len(dir) == 0 {
		dir = "./"
	}

	m, err := migrate.New(
		"file://"+dir+"/res/migrations",
		url+"?sslmode=disable",
	)

	if err != nil {
		return err
	}

	defer m.Close()

	err = m.Up()

	if err != nil && err.Error() != "no change" {
		return err
	}

	return nil
}
