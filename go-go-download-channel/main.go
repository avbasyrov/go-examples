package main // go-go-download-channel

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	started := time.Now()

	// WaitGroup используется для ожидания всех горутин
	wg := sync.WaitGroup{}

	// Создаем канал для строк, будем писать в него результаты каждой загрузки файлов
	ch := make(chan string)

	// Получаем строки из канала
	go func() {
		for content := range ch {
			fmt.Println("Содержимое файла:", content)
		}
	}()

	for i := 1; i <= 10; i++ {
		wg.Add(1) // перед запуском новой горутины учитываем её в счетчике
		go downloadBigFile(i, &wg, ch)
	}

	// Wait будет ожидать до тех пор, пока счетчик запущенных горутин в WaitGroup не обнулится
	wg.Wait()

	fmt.Println("Длительность:", time.Now().Sub(started).String())
}

func downloadBigFile(fileNum int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done() // при завершении горутины обязательно помечаем, что она завершилась

	// Случайная задержка от 0 до 1 секунды
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	fmt.Println("Файл", fileNum, "загружен")

	ch <- "content-" + strconv.Itoa(fileNum)
}
