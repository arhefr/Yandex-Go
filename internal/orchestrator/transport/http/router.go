package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Port string

	PathGet  string
	PathAdd  string
	PathTask string
}

type Router struct {
	*Config
	Echo *echo.Echo
}

func NewRouter(config *Config) Router {

	e := echo.New()
	e.POST(config.PathAdd, AddExpr)
	e.GET(config.PathGet, GetIDs)
	e.GET(config.PathGet+"/:id", GetID)
	e.GET(config.PathTask, GetTask)
	e.POST(config.PathTask, FetchTask)

	return Router{Config: config, Echo: e}
}

func (r *Router) Run() {
	server := http.Server{
		Addr:    ":" + r.Config.Port,
		Handler: r.Echo,
	}

	log.Infof("Server start on localhost:%s", r.Port)
	log.Fatal(server.ListenAndServe())
}
