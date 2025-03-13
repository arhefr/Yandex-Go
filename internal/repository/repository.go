package repository

import (
	"calculator/internal/orchestrator/models"
	"sync"
)

var Tasks = NewSafeMap()

type repository struct {
	mu sync.Mutex
	m  map[string]models.Expression
}

func NewSafeMap() *repository {
	return &repository{m: make(map[string]models.Expression), mu: sync.Mutex{}}
}

func (r *repository) Add(key string, value models.Expression) {
	r.mu.Lock()
	r.m[key] = value
	r.mu.Unlock()
}

func (r *repository) Get(key string) (models.Expression, bool) {
	r.mu.Lock()

	if value, ok := r.m[key]; ok {
		r.mu.Unlock()
		return value, ok
	}
	r.mu.Unlock()
	return models.Expression{}, false
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

func (r *repository) GetValues() []models.Expression {
	var array []models.Expression
	r.mu.Lock()
	for key := range r.m {
		array = append(array, r.m[key])
	}
	r.mu.Unlock()
	return array
}
