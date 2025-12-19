package config

import (
	"log/slog"
	"os"
)

type ServerConf struct {
	Port   string
	Env    string
	Logger *slog.Logger
}

func NewServerConf() ServerConf {
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return ServerConf{GetEnv("APP_PORT", ":8081"), GetEnv("APP_ENV", "local"), l}

}

func GetEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return def
}
