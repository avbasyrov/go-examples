package list

import "sync"

type Map struct {
	items map[string]string

	mu sync.Mutex
}

func NewMapList() *Map {
	return &Map{items: make(map[string]string)}
}

func (m *Map) Push(key string, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.items[key] = value
}
