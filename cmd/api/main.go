package main

import (
	"fmt"
	"log/slog"

	"github.com/b612lpp/goprj002/application"
	"github.com/b612lpp/goprj002/internal/config"
	"github.com/b612lpp/goprj002/internal/delivery/http/auth"
	"github.com/b612lpp/goprj002/internal/delivery/http/health"
	"github.com/b612lpp/goprj002/internal/delivery/http/meter"
	"github.com/b612lpp/goprj002/internal/router"
	"github.com/b612lpp/goprj002/internal/server"
)

func main() {

	c := config.NewServerConf()
	slog.SetDefault(c.Logger)
	r := router.NewMyRouter()

	//создаём юз кейсы
	gasMeterUseCase := application.NewSubmitReadingGas(c.Db)
	enMeterUseCase := application.NewSubmitReadingEn(c.Db)

	health := health.NewHealthHandler()
	r.AddPublicRout("/public/health", health.ResponsOK)
	//создаём хэндлер, передаём юз кейс, добавляем маршруты
	a := auth.NewAuth()
	r.AddPublicRout("/public/auth/registry", a.Registry)
	r.AddPublicRout("/public/auth/authenticate", a.Authenticate)

	//создаем хэндлер, передаёмюз кейс
	meterGasHandler := meter.NewGasMeterHandler(*gasMeterUseCase)
	meterEnHandler := meter.NewEnMeterHandler(*enMeterUseCase)

	//Создаем маршруты
	r.AddPrivateRout("/private/metergas", meterGasHandler.GetGasValues)
	r.AddPrivateRout("/private/meteren", meterEnHandler.GetEnValues)

	r.CompilemmMux()
	s := server.NewMyServer(c, r)

	fmt.Printf("Сервер запущен. Порт %s база данных %s", c.Port, c.Db.GetTitle())
	s.Run()

}
