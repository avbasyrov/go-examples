package main

import (
	"go-examples/go-benchmarks/internal/pkg/list"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	s := list.NewMapList()

	ticker := time.NewTicker(10 * time.Millisecond)
	var mu sync.Mutex
	wg := sync.WaitGroup{}

	runtime.SetMutexProfileFraction(2)

	for thread := 0; thread < runtime.GOMAXPROCS(0); thread++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			i := 0

			for t := range ticker.C {
				i++
				log.Println(t)

				mu.Lock()
				key := "key_" + strconv.Itoa(i)
				value := "value_" + strconv.Itoa(doSomeJob(i))
				s.Push(key, value)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
}

func doSomeJob(i int) int {
	result := 0
	for a := 0; a < i; a++ {
		result += a
	}
	return result
}
