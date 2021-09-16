package model

import "go-examples/go-postgres/internal/pkg/storage"

type Logic struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) *Logic {
	return &Logic{storage: storage}
}

func (l *Logic) CreateAndSelectStudent(name, surname string, age int) ([]storage.Student, error) {
	err := l.storage.AddStudent(name, surname, age)
	if err != nil {
		return nil, err
	}

	return l.storage.SelectStudentsByName(name)
}
