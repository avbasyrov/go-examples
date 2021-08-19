package main // go-go-download-channel-buffered

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	// wgDownload используется для ожидания горутин скачивающих файлы
	wgDownload := sync.WaitGroup{}

	// Создаем буферизированный канал для строк, будем писать в него результаты каждой загрузки файлов
	ch := make(chan string, 5)

	for i := 1; i <= 10; i++ {
		wgDownload.Add(1) // перед запуском новой горутины учитываем её в счетчике
		go downloadBigFile(i, &wgDownload, ch)
	}

	var wgPrintContent sync.WaitGroup
	wgPrintContent.Add(1)
	go func() {
		defer wgPrintContent.Done()

		// Получаем строки из канала, пока он не закроется
		for {
			content, ok := <-ch
			if !ok {
				fmt.Println("Канал закрыт, больше нечего из него читать")
				break
			}

			fmt.Println("Содержимое файла:", content)
		}
	}()

	// Wait будет ожидать до тех пор, пока счетчик запущенных горутин в WaitGroup не обнулится
	wgDownload.Wait()
	close(ch)

	// Ждем пока выведутся все строки из канала
	wgPrintContent.Wait()
}

func downloadBigFile(fileNum int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done() // при завершении горутины обязательно помечаем, что она завершилась

	// Случайная задержка от 0 до 1 секунды
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	fmt.Println("Файл", fileNum, "загружен")

	ch <- "content-" + strconv.Itoa(fileNum)
}
