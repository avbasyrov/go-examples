package list

import (
	"sync"
	"testing"
)

func TestSlice_Push(t *testing.T) {
	s := NewSliceList()

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Push("key-1", "value-1")
		}()
	}
	wg.Wait()
}
