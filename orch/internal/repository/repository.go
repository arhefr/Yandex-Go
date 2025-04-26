package repository

import (
	"sync"

	"github.com/arhefr/Yandex-Go/orch/internal/model"
	"github.com/arhefr/Yandex-Go/orch/internal/tools"
)

var Exprs = NewRepo()

type Repository struct {
	mu sync.Mutex
	m  map[string]model.Request
}

func NewRepo() *Repository {
	return &Repository{m: make(map[string]model.Request), mu: sync.Mutex{}}
}

func (r *Repository) Add(expr model.Expression) (string, error) {
	r.mu.Lock()
	key := tools.NewCryptoRand()
	req := model.NewExpr(key, &expr)
	r.m[key] = req
	r.mu.Unlock()

	return key, nil
}

func (r *Repository) Replace(key string, req model.Request) {
	r.mu.Lock()
	r.m[key] = req
	r.mu.Unlock()
}

func (r *Repository) GetByID(key string) (model.Request, bool) {
	r.mu.Lock()

	if value, ok := r.m[key]; ok {
		r.mu.Unlock()
		return value, ok
	}
	r.mu.Unlock()
	return model.Request{}, false
}

func (r *Repository) Delete(key string) {
	r.mu.Lock()
	delete(r.m, key)
	r.mu.Unlock()
}

func (r *Repository) GetKeys() []string {
	var array []string
	r.mu.Lock()
	for value := range r.m {
		array = append(array, value)
	}
	r.mu.Unlock()
	return array
}

func (r *Repository) Get() []model.Request {
	var array []model.Request
	r.mu.Lock()
	for key := range r.m {
		array = append(array, r.m[key])
	}
	r.mu.Unlock()
	return array
}
