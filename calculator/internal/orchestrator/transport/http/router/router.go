package router

import (
	"calculator/internal/orchestrator/transport/http/handler"
	"calculator/internal/orchestrator/transport/http/middleware"
	"log"
	"net/http"
)

type Config struct {
	Port string
	Host string

	PathGet  string
	PathAdd  string
	PathTask string
}

type OrchestratorRouter struct {
	Config
	Router *http.ServeMux
}

func NewRouter(config Config) OrchestratorRouter {
	mux := http.NewServeMux()
	mux.Handle(config.PathAdd, middleware.MethodCheck("POST", http.HandlerFunc(handler.AddTask)))
	mux.Handle(config.PathGet, middleware.MethodCheck("GET", http.HandlerFunc(handler.GetIDs)))
	mux.Handle(config.PathGet+"/{id}", middleware.MethodCheck("GET", http.HandlerFunc(handler.GetID)))
	mux.HandleFunc(config.PathTask, handler.GetOperation)

	return OrchestratorRouter{Config: config, Router: mux}
}

func (r *OrchestratorRouter) Run() {
	server := http.Server{
		Addr:    ":" + r.Config.Port,
		Handler: middleware.Logging(r.PathTask, r.Router),
	}

	log.Printf("server starting on %s:%s\n", r.Host, r.Port)
	log.Fatal(server.ListenAndServe())
}
