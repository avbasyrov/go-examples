package main // go-go-download

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	started := time.Now()

	for i := 1; i <= 10; i++ {
		downloadBigFile(i)
	}

	fmt.Println("Длительность:", time.Now().Sub(started).String())
}

func downloadBigFile(fileNum int) {
	// Случайная задержка от 1 до 2 секунд
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	fmt.Println("Файл", fileNum, "загружен")
}
