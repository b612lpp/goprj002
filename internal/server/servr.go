package server

import (
	"net/http"

	"github.com/b612lpp/goprj002/application"
	"github.com/b612lpp/goprj002/internal/config"
	"github.com/b612lpp/goprj002/internal/router"
	"github.com/b612lpp/goprj002/repository"
)

type Server interface {
	Run() error
	//Shutdown(ctx context.Context) error
}

type MyServer struct {
	httpServer *http.Server
	db         repository.Repo
}

func NewMyServer(c config.ServerConf, r router.Router) *MyServer {

	return &MyServer{httpServer: &http.Server{Addr: c.Port, Handler: r.Handler()}, db: repository.NemIMDB()}

}

func (ms *MyServer) Run() error {
	x := application.NewSubmitReading(ms.db)
	return ms.httpServer.ListenAndServe()

}
