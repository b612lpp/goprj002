package main

import (
	"log/slog"

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

	health := health.NewHealthHandler()
	r.AddPublicRout("/public/health", health.ResponsOK)

	a := auth.NewAuth()
	r.AddPublicRout("/public/auth/registry", a.Registry)
	r.AddPublicRout("/public/auth/authenticate", a.Authenticate)

	meter := meter.NewMeter()
	r.AddPrivateRout("/private/meter", meter.TmpMeter)

	r.CompilemmMux()
	s := server.NewMyServer(c, r)
	s.Run()

}
