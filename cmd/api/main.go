package main

import (
	"github.com/b612lpp/goprj002/internal/config"
	"github.com/b612lpp/goprj002/internal/delivery/http/health"
	"github.com/b612lpp/goprj002/internal/router"
	"github.com/b612lpp/goprj002/internal/server"
)

func main() {
	c := config.NewServerConf()
	r := router.NewMyRouter()
	r.AddRout("/", health.NewHealthHandler().ResponsOK)
	s := server.NewMyServer(c, r)
	s.Run()

}
