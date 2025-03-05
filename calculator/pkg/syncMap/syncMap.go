package syncMap

import (
	"calculator/internal/orchestrator/transport/http/models"
	"fmt"
	"sync"
)

type SyncMap struct {
	mu sync.Mutex
	m  map[int]models.Expression
}

func (m *SyncMap) Add(key int, value models.Expression) {
	m.mu.Lock()
	m.m[key] = value
	m.mu.Unlock()
}

func (m *SyncMap) Delete(key int) {
	m.mu.Lock()
	delete(m.m, key)
	m.mu.Unlock()
}

func (m *SyncMap) GetKeys() []int {
	var array []int
	m.mu.Lock()
	for value := range m.m {
		array = append(array, value)
	}
	m.mu.Unlock()
	return array
}

func (m *SyncMap) GetValues() []models.Expression {
	var array []models.Expression
	m.mu.Lock()
	for key := range m.m {
		array = append(array, m.m[key])
	}
	m.mu.Unlock()
	return array
}

func (m *SyncMap) Get(key int) (models.Expression, error) {
	m.mu.Lock()

	if value, ok := m.m[key]; ok {
		m.mu.Unlock()
		return value, nil
	}
	m.mu.Unlock()
	return models.Expression{}, fmt.Errorf("unknown key")
}

func NewSafeMap() *SyncMap {
	return &SyncMap{m: make(map[int]models.Expression), mu: sync.Mutex{}}
}
