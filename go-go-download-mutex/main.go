package main // go-go-download-mutex

import (
	"fmt"
	"strconv"
	"sync"
)

var content = map[int]string{} // содержимое файла по его номеру
var mutex = sync.RWMutex{}

func main() {
	content = make(map[int]string) // Инициализируем мапу
	wgDownload := sync.WaitGroup{} // Для ожилания горутин, скачивающих файлы

	for i := 1; i <= 10; i++ {
		wgDownload.Add(1) // перед запуском новой горутины учитываем её в счетчике
		go downloadBigFile(i, &wgDownload)
	}

	wgDownload.Wait() // Ждем загрузки файлов

	// Выводим содержимое полученных файлов
	mutex.RLock() // лочим для чтения
	fmt.Println(content)
	mutex.RUnlock() // освобождаем лок
}

func downloadBigFile(fileNum int, wg *sync.WaitGroup) {
	defer wg.Done() // при завершении горутины обязательно помечаем, что она завершилась
	fmt.Println("Файл", fileNum, "загружен")

	mutex.Lock()                                          // Лочим мьютекс
	content[fileNum] = "content-" + strconv.Itoa(fileNum) // Записываем в мапу
	mutex.Unlock()                                        // Обязательно освобождаем мьютекс
}
