package main // go-cli-signals

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)

	// Перехватываем сигналы
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2)

	// Получаем приходящие сигналы из канала
	for {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
	}
}
