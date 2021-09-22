package list_ints

import "sync"

type List struct {
	items []int

	mu sync.Mutex
}

func (l *List) Push(value int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.items = append(l.items, value)
}
