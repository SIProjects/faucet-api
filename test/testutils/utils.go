package testutils

import (
	"context"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/SIProjects/faucet-api/app"
	"github.com/SIProjects/faucet-api/chain"
	"github.com/SIProjects/faucet-api/database"
	"github.com/SIProjects/faucet-api/test/fixture"
	"github.com/jackc/pgx/v4"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func randomName(n uint) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type SandboxContext struct {
	App           *app.App
	Name          string
	ConnectionURL string
	Cache         *MockCache
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func databaseUrl(dbName string) string {
	dbUrl := getenv("TEST_DATABASE_URL", "postgres://postgres:postgres@localhost:5432")

	return dbUrl + "/" + dbName
}

func connectAdmin() (*pgx.Conn, error) {
	dbUrl := databaseUrl("postgres")
	return pgx.Connect(context.Background(), dbUrl)
}

func (s *SandboxContext) Close() {
	s.App.Close()

	conn, err := connectAdmin()

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close(context.Background())

	_, err = conn.Exec(
		context.Background(),
		"DROP DATABASE "+s.Name,
	)

	if err != nil {
		log.Fatalln(err)
	}
}

func NewSandbox(fx *fixture.Fixture) (*SandboxContext, error) {
	os.Setenv("CWD", "../../")

	conn, err := connectAdmin()

	if err != nil {
		return nil, err
	}

	name := randomName(20)

	_, err = conn.Exec(
		context.Background(),
		"CREATE DATABASE "+name,
	)

	if err != nil {
		return nil, err
	}

	err = conn.Close(context.Background())

	if err != nil {
		return nil, err
	}

	dbUrl := databaseUrl(name)

	db, err := database.New(dbUrl)

	if err != nil {
		return nil, err
	}

	cache := NewMockCache(fx.Mocks.Cache)

	n := NewMockNode(fx.Mocks.Node)

	ch := chain.New(n, chain.Testnet)

	l := log.New(os.Stdout, "", 0)

	config := fx.Config.ToConfig()
	a, err := app.New(db, cache, ch, NewMockScheduler(), l, &config)

	if err != nil {
		return nil, err
	}

	sb := SandboxContext{
		App:           a,
		Name:          name,
		ConnectionURL: dbUrl,
		Cache:         cache,
	}

	return &sb, nil
}
