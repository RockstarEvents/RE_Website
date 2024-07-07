package config

type Config struct {
	Port string `env:"REST_PORT" default:"8082"`
}
