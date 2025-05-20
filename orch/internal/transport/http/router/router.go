package router

import (
	"net/http"

	"github.com/arhefr/Yandex-Go/orch/internal/transport/http/handlers"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Port string
}

type Router struct {
	*Config
	Echo *echo.Echo
}

func NewRouter(config *Config, h *handlers.Handler) Router {

	e := echo.New()
	reqAuth := e.Group("/api/v1", h.AuthRequired)
	{
		reqAuth.POST(ENDPOINT_ADD, h.AddExpr)
		reqAuth.GET(ENDPOINT_GET, h.GetIDs)
		reqAuth.GET(ENDPOINT_GET+"/:id", h.GetID)
	}
	auth := e.Group("/api/v1")
	{
		auth.POST(ENDPOINT_SIGN_IN, h.SignIn)
		auth.POST(ENDPOINT_LOG_IN, h.LogIn)
	}

	e.GET(ENDPOINT_TASK, h.SendTask)
	e.POST(ENDPOINT_TASK, h.CatchTask)

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
