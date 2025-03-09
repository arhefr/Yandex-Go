package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
		Handler: Logging("/internal/task", r.Router),
	}

	log.Printf("server starting on %s:%s\n", r.Host, r.Port)
	log.Fatal(server.ListenAndServe())
}
