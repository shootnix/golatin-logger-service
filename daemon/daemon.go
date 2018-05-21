package daemon

import (
	"github.com/gorilla/mux"
	"github.com/shootnix/golatin-logger-service/config"
	"github.com/shootnix/golatin-logger-service/controllers"
	"log"
	"net/http"
)

type Daemon struct {
	Srv *http.Server
}

func NewDaemon(cfg config.DaemonConfig) *Daemon {
	r := mux.NewRouter()

	r.HandleFunc("/log", controllers.POSTLog).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    cfg.Address,
	}

	d := &Daemon{
		Srv: srv,
	}

	return d
}

func (d *Daemon) Run() {
	log.Fatal(d.Srv.ListenAndServe())
}
