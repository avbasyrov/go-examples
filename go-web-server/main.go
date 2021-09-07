package main

import (
	"fmt"
	"go-examples/go-web-server/internal/rest"
	"go-examples/go-web-server/internal/storage"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	server := rest.New(storage.New())
	fmt.Println("Список номеров - http://localhost" + port + "/list")
	fmt.Println("Проверка по номеру - http://localhost" + port + "/phone/74212340077")
	log.Fatal(http.ListenAndServe(port, server))
}
