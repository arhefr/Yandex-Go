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
	g := e.Group("/api/v1")
	{
		g.POST(ENDPOINT_ADD, h.AddExpr)
		g.GET(ENDPOINT_GET, h.GetIDs)
		g.GET(ENDPOINT_GET+"/:id", h.GetID)

		g.POST(ENDPOINT_SIGN_IN, h.SignIn)
		g.POST(ENDPOINT_LOG_IN, h.LogIn)
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
