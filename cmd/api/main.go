package main

import (
	"log/slog"

	"github.com/b612lpp/goprj002/internal/config"
	"github.com/b612lpp/goprj002/internal/delivery/http/auth"
	"github.com/b612lpp/goprj002/internal/delivery/http/health"
	"github.com/b612lpp/goprj002/internal/router"
	"github.com/b612lpp/goprj002/internal/server"
)

func main() {

	c := config.NewServerConf()
	slog.SetDefault(c.Logger)
	r := router.NewMyRouter()

	m := health.NewHealthHandler()
	r.AddRout("/health", m.ResponsOK)

	a := auth.NewAuth()
	r.AddRout("/auth/registry", a.Registry)
	r.AddRout("/auth/authenticate", a.Authenticate)

	s := server.NewMyServer(c, r)
	s.Run()

}
