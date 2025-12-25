package config

import (
	"log/slog"
	"os"

	"github.com/b612lpp/goprj002/repository"
)

type ServerConf struct {
	Port   string
	Env    string
	Logger *slog.Logger
	Db     repository.Repo
}

// читаем переменные окружения чтобы собрать экземпляр конфигурации по инфраструктуре
func NewServerConf() ServerConf {
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db := repository.NemIMDB()
	return ServerConf{GetEnv("APP_PORT", ":8081"), GetEnv("APP_ENV", "local"), l, db}

}

// вызов ОС функции на чтение переменных
func GetEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return def
}
