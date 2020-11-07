package server

// Config ...
type Config struct {
	dbURL string
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		dbURL: "postgres://postgres:postgres@localhost:5432/postgres", // FIXME
	}
}
