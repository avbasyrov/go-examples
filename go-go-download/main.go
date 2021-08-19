package main // go-go-download

import (
	"fmt"
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
	time.Sleep(time.Second)
	fmt.Println("Файл", fileNum, "загружен")
}
