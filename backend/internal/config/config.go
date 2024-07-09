package config

type Config struct {
	Port         string `env:"REST_PORT" default:"8082"`
	DataBasePath string `env:"DATA_BASE" default:"postgres://postgres:postgres@postgres:5432/postgres"`
}
