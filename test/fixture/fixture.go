package fixture

import "github.com/SIProjects/faucet-api/configuration"

type FixtureHeader struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type FixtureRequest struct {
	Path    string          `yaml:"path"`
	Method  string          `yaml:"method"`
	Headers []FixtureHeader `yaml:"headers"`
	Body    string          `yaml:"body"`
}

type FixtureResponse struct {
	Status int     `yaml:"status"`
	Body   *string `yaml:"body"`
}

type SetupFixtures struct {
	DB []string `yaml:"db"`
}

type FixtureConfig struct {
	MinPayout int `yaml:"min_payout"`
	MaxPayout int `yaml:"max_payout"`
}

func (f *FixtureConfig) ToConfig() configuration.Config {
	minPayout := 10
	maxPayout := 100
	if f != nil {
		minPayout = f.MinPayout
		maxPayout = f.MaxPayout
	}

	return configuration.Config{
		MinPayout: minPayout,
		MaxPayout: maxPayout,
	}
}

type CacheResults struct {
	PendingResults []string `yaml:"pending"`
}

type AddressExistsFixture struct {
	Address string `yaml:"address"`
	Exists  bool   `yaml:"exists"`
	Error   bool   `yaml:"error"`
}

type DatabaseCheck struct {
	Query string `yaml:"query"`
	Rows  int    `yaml:"rows"`
}

type CacheFixtures struct {
	GetQueuedCount int                    `yaml:"get_queued_count"`
	ExistsFixtures []AddressExistsFixture `yaml:"address_exists"`
}

type NodeFixtures struct {
	GetBalance float64 `yaml:"get_balance"`
}

type MockFixtures struct {
	Cache *CacheFixtures `yaml:"cache"`
	Node  *NodeFixtures  `yaml:"node"`
}

type Fixture struct {
	Name           string          `yaml:"name"`
	Description    string          `yaml:"description"`
	Mocks          MockFixtures    `yaml:"mocks"`
	Config         FixtureConfig   `yaml:"config"`
	Request        FixtureRequest  `yaml:"request"`
	Response       FixtureResponse `yaml:"response"`
	DatabaseChecks []DatabaseCheck `yaml:"database"`
	Cache          CacheResults    `yaml:"cache"`
	Setup          SetupFixtures   `yaml:"setup"`
}
