package repository

import (
	"sync"

	"github.com/arhefr/Yandex-Go/orch/internal/model"
)

var Exprs = NewSafeMap()

type repository struct {
	mu sync.Mutex
	m  map[string]model.Request
}

func NewSafeMap() *repository {
	return &repository{m: make(map[string]model.Request), mu: sync.Mutex{}}
}

func (r *repository) Add(key string, value model.Request) {
	r.mu.Lock()
	r.m[key] = value
	r.mu.Unlock()
}

func (r *repository) Get(key string) (model.Request, bool) {
	r.mu.Lock()

	if value, ok := r.m[key]; ok {
		r.mu.Unlock()
		return value, ok
	}
	r.mu.Unlock()
	return model.Request{}, false
}

func (r *repository) Delete(key string) {
	r.mu.Lock()
	delete(r.m, key)
	r.mu.Unlock()
}

func (r *repository) GetKeys() []string {
	var array []string
	r.mu.Lock()
	for value := range r.m {
		array = append(array, value)
	}
	r.mu.Unlock()
	return array
}

func (r *repository) GetValues() []model.Request {
	var array []model.Request
	r.mu.Lock()
	for key := range r.m {
		array = append(array, r.m[key])
	}
	r.mu.Unlock()
	return array
}
