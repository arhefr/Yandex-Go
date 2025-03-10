package router

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Host string
	Port string

	PathGet  string
	PathAdd  string
	PathTask string
}

type Router struct {
	Config
	Router *mux.Router
}

func NewRouter(config Config) Router {
	mux := mux.NewRouter()

	mux.HandleFunc(config.PathAdd, AddTask).
		Methods("POST")
	mux.HandleFunc(config.PathGet, GetIDs).
		Methods("GET")
	mux.HandleFunc(config.PathGet+"/{id}", GetID).
		Methods("GET")
	mux.HandleFunc(config.PathTask, GetOperation).
		Methods("POST", "GET")

	return Router{Config: config, Router: mux}
}

func (r *Router) Run() {
	server := http.Server{
		Addr:    ":" + r.Config.Port,
		Handler: Logging(r.Router),
	}

	log.Infof("Server start listening on %s:%s", r.Host, r.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server crush: ", err)
	}
}
