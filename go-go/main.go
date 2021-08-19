// go-go Показывать
package main

import (
	"fmt"
	"time"
)

func main() {
	go printEverySecond("a")
	printEverySecond("b")
}

func printEverySecond(str string) {
	for {
		fmt.Println(str)
		time.Sleep(time.Millisecond)
	}
}
