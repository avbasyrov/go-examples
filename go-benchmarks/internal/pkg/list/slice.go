package list

import "sync"

type itemType struct {
	key   string
	value string
}

type Slice struct {
	items []itemType

	mu sync.Mutex
}

func NewSliceList() *Slice {
	return &Slice{}
}

func (s *Slice) Push(key string, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items = append(s.items, itemType{
		key:   key,
		value: value,
	})
}
