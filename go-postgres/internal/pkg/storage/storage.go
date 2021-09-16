package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Storage struct {
	ctx     context.Context
	db      *sql.DB
	timeout time.Duration
}

type Student struct {
	Name    string
	Surname string
	Age     int
}

func New(host string, port int, user, password, dbname string, timeout time.Duration) (*Storage, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	return &Storage{
		ctx:     context.Background(),
		db:      db,
		timeout: timeout,
	}, err
}

func NewFromConnection(db *sql.DB, timeout time.Duration) *Storage {
	return &Storage{
		ctx:     context.Background(),
		db:      db,
		timeout: timeout,
	}
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) AddStudent(name, surname string, age int) error {
	ctx, cancel := context.WithTimeout(s.ctx, s.timeout)
	defer cancel()

	_, err := s.db.ExecContext(ctx, "INSERT INTO school.students(name, surname, age) VALUES ($1, $2, $3)", name, surname, age)
	return err
}

func (s *Storage) SelectStudentsByName(name string) ([]Student, error) {
	ctx, cancel := context.WithTimeout(s.ctx, s.timeout)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, "SELECT age, name, surname FROM school.students WHERE name=$1", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	students := make([]Student, 0)

	for rows.Next() {
		student := Student{}

		if err := rows.Scan(&student.Age, &student.Name, &student.Surname); err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}
