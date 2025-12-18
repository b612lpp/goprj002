package config

import "os"

type ServerConf struct {
	Port string
	Env  string
}

func NewServerConf() ServerConf {
	return ServerConf{GetEnv("APP_PORT", ":8081"), GetEnv("APP_ENV", "local")}

}

func GetEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return def
}
