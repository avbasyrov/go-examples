package main // go-go-download-select

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	wgDownload := sync.WaitGroup{} // Для ожилания горутин, скачивающих файлы
	ch := make(chan string, 5)     // Канал для содержимого файлов
	done := make(chan int)         // Канал для сигнала, что мы закончили отображать контент файлов

	for i := 1; i <= 3; i++ {
		wgDownload.Add(1) // перед запуском новой горутины учитываем её в счетчике
		go downloadBigFile(i, &wgDownload, ch)
	}

	go func() {
		wgDownload.Wait() // Ждем загрузки файлов
		// Отправляем в канал done что-нибудь, чтобы просигнализировать, что мы закончили загрузку
		done <- 0
	}()

	// Ждем из каналов содержимое файлов и сигнал об окончании загрузки
	for {
		select {
		case <-done:
			fmt.Println("Пришел сигнал, что загрузка файлов завершена")
			return
		case file := <-ch:
			fmt.Println("Содержимое файла:", file)
		}
	}
}

func downloadBigFile(fileNum int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done() // при завершении горутины обязательно помечаем, что она завершилась
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	fmt.Println("Файл", fileNum, "загружен")
	ch <- "content-" + strconv.Itoa(fileNum)
}
