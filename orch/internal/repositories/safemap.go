package repository

import (
	"sync"

	"github.com/arhefr/Yandex-Go/orch/internal/model"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[string]model.Request
}

func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[string]model.Request), mu: sync.Mutex{}}
}

func (r *SafeMap) Add(req model.Request, key string) {
	r.mu.Lock()
	r.m[key] = req
	r.mu.Unlock()
}

func (r *SafeMap) Replace(key string, req model.Request) {
	r.mu.Lock()
	r.m[key] = req
	r.mu.Unlock()
}

func (r *SafeMap) GetByID(key string) (*model.Request, bool) {
	r.mu.Lock()

	if value, ok := r.m[key]; ok {
		r.mu.Unlock()
		return &value, ok
	}
	r.mu.Unlock()
	return nil, false
}

func (r *SafeMap) Delete(key string) {
	r.mu.Lock()
	delete(r.m, key)
	r.mu.Unlock()
}

func (r *SafeMap) GetKeys() (arr []string) {

	r.mu.Lock()
	for value := range r.m {
		arr = append(arr, value)
	}
	r.mu.Unlock()
	return arr
}

func (r *SafeMap) Get() (arr []model.Request) {

	r.mu.Lock()
	for key := range r.m {
		arr = append(arr, r.m[key])
	}
	r.mu.Unlock()
	return arr
}
