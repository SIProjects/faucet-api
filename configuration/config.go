package configuration

type Config struct {
	MinPayout int
	MaxPayout int
}

func New(min, max int) *Config {
	return &Config{
		MinPayout: min,
		MaxPayout: max,
	}
}
