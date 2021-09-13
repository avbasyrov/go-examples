package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите текст: ")
	text, _ := reader.ReadString('\n')

	_, _ = os.Stdout.Write([]byte("Введенный текст: "))
	_, _ = os.Stdout.Write([]byte(text))
}
