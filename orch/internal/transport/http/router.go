package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

const (
	PATH_ADD  = "/api/v1/calculate"
	PATH_GET  = "/api/v1/expressions"
	PATH_TASK = "/internal/task"
)

type Config struct {
	Port string
}

type Router struct {
	*Config
	Echo *echo.Echo
}

func NewRouter(config *Config) Router {

	e := echo.New()
	e.POST(PATH_ADD, AddExpr)
	e.GET(PATH_GET, GetIDs)
	e.GET(PATH_GET+"/:id", GetID)
	e.GET(PATH_TASK, SendTask)
	e.POST(PATH_TASK, CatchTask)

	return Router{Config: config, Echo: e}
}

func (r *Router) Run() {
	server := http.Server{
		Addr:    ":" + r.Config.Port,
		Handler: r.Echo,
	}

	log.Infof("Server start on localhost:%s", r.Config.Port)
	log.Fatal(server.ListenAndServe())
}
