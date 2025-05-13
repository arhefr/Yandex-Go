package handlers

import (
	"sync"

	"github.com/arhefr/Yandex-Go/orch/internal/services"
)

type Handler struct {
	mu       sync.RWMutex
	services *services.Services
}

func NewHandler(services *services.Services) Handler {
	return Handler{services: services, mu: sync.RWMutex{}}
}
