package main // go-go-download-waitgroup

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	started := time.Now()

	// WaitGroup используется для ожидания всех горутин
	wg := sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		wg.Add(1) // перед запуском новой горутины учитываем её в счетчике
		go downloadBigFile(i, &wg)
	}

	// Wait будет ожидать до тех пор, пока счетчик запущенных горутин в WaitGroup не обнулится
	wg.Wait()

	fmt.Println("Длительность:", time.Now().Sub(started).String())
}

func downloadBigFile(fileNum int, wg *sync.WaitGroup) {
	defer wg.Done() // при завершении горутины обязательно помечаем, что она завершилась

	// Случайная задержка от 1 до 2 секунд
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	fmt.Println("Файл", fileNum, "загружен")
}
