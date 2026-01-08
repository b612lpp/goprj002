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
	WH     repository.ReadingStorage
}

// читаем переменные окружения чтобы собрать экземпляр конфигурации по инфраструктуре
func NewServerConf() ServerConf {
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))
	wh := repository.NewWareHouse(repository.NemIMDB(), repository.NewEventDb())

	return ServerConf{GetEnv("APP_PORT", ":8081"), GetEnv("APP_ENV", "local"), l, wh}

}

// вызов ОС функции на чтение переменных
func GetEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return def
}
