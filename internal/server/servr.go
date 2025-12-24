package server

import (
	"net/http"

	"github.com/b612lpp/goprj002/internal/config"
	"github.com/b612lpp/goprj002/internal/router"
)

type Server interface {
	Run() error
	//Shutdown(ctx context.Context) error
}

type MyServer struct {
	httpServer *http.Server
}

func NewMyServer(c config.ServerConf, r router.Router) *MyServer {

	return &MyServer{httpServer: &http.Server{Addr: c.Port, Handler: r.Handler()}}

}

func (ms *MyServer) Run() error {

	return ms.httpServer.ListenAndServe()

}
