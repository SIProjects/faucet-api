package fixture

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
	Status int `yaml:"status"`
}

type CacheFixtures struct {
	ExistsFixtures []AddressExistsFixture `yaml:address-exists`
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

type Fixture struct {
	Name           string          `yaml:"name"`
	Description    string          `yaml:"description"`
	Request        FixtureRequest  `yaml:"request"`
	Response       FixtureResponse `yaml:"response"`
	DatabaseChecks []DatabaseCheck `yaml:"database"`
	Cache          CacheFixtures   `yaml:"cache"`
}
