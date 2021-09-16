package main

import (
	"go-examples/go-postgres/internal/pkg/storage"
	"log"
	"time"
)

func main() {
	s, err := storage.New("127.0.0.1", 5432, "postgres", "12345", "postgres", time.Second)
	if err != nil {
		log.Fatal(err)
	}

	err = s.AddStudent("Victor", "--------------", 24)
	if err != nil {
		log.Fatal(err)
	}

	students, err := s.SelectStudentsByName("Victor")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(students)
}
