package list_ints

import "sync"

type List struct {
	items []int

	ch chan int

	mu sync.Mutex
}

func New() *List {
	l := &List{
		ch: make(chan int, 50),
	}

	go l.process()

	return l
}

func (l *List) Push(value int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.items = append(l.items, value)
}

func (l *List) PushChannel(value int) {
	l.ch <- value
}

func (l *List) process() {
	for value := range l.ch {
		l.items = append(l.items, value)
	}
}
