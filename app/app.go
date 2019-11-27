package app

import (
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"langar/app/handler"
	"langar/app/model"
	"langar/config"
	"net/http"
)

type App struct {
	Router    *mux.Router
	Auth      *model.Auth
	DockerCli *client.Client
}

// INIT
func (a *App) Init() {
	a.Router = mux.NewRouter()
	a.setRouters()

	a.Auth = &model.Auth{
		AccessKey: config.Config.AccessKey,
		SecretKey: config.Config.SecretKey,
	}
}

func (a *App) setRouters() {
	APISubRouter := a.Router.PathPrefix("/api").Subrouter()

	APISubRouter.HandleFunc("/network", a.GetNetworks).Methods("GET")
}

// NETWORK
func (a *App) GetNetworks(w http.ResponseWriter, r *http.Request) {
	if handler.IsAuthorized(w, r, a.Auth) {
		handler.GetNetworks(w, r)
	}
}

// RUN
func (a *App) Run(host string) {
	log.Infoln("Starting ...")
	log.Fatal(http.ListenAndServe(host, a.Router))
}