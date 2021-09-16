package model

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go-examples/go-postgres/internal/pkg/storage"
	"log"
	"regexp"
	"testing"
	"time"
)

func TestLogic_CreateAndSelectStudent(t *testing.T) {
	db, mock := NewMock()
	s := storage.NewFromConnection(db, time.Second)
	defer s.Close()

	l := New(s)

	queryInsert := regexp.QuoteMeta("INSERT INTO school.students(name, surname, age) VALUES ($1, $2, $3)")
	mock.ExpectExec(queryInsert).
		WithArgs("Ivan", "Sidorov", 18).
		WillReturnResult(sqlmock.NewResult(0, 0))

	querySelect := regexp.QuoteMeta("SELECT name, surname, age FROM school.students WHERE name=$1")
	rows := sqlmock.NewRows([]string{"name", "surname", "age"}).
		AddRow("Ivan", "Sidorov", 18).
		AddRow("Ivan", "Petrov", 23)

	mock.ExpectQuery(querySelect).WithArgs("Ivan").WillReturnRows(rows)

	students, err := l.CreateAndSelectStudent("Ivan", "Sidorov", 18)
	require.NoError(t, err)
	assert.Equal(t, 2, len(students))
	assert.Equal(t, "Ivan", students[0].Name)
	assert.Equal(t, "Sidorov", students[0].Surname)
	assert.Equal(t, 18, students[0].Age)
	assert.Equal(t, "Ivan", students[1].Name)
	assert.Equal(t, "Petrov", students[1].Surname)
	assert.Equal(t, 23, students[1].Age)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}
