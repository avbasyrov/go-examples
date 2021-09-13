package main // go-cli-args

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args) // выведет слайс из аргументов, нулевым элементом будет имя программы
}
