package rest

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSuccess(t *testing.T) {
	// Создаем наш веб-сервер с фейковым хранилищем
	storage := newFakeStorage()
	server := New(storage)

	// Создаем новый HTTP-запрос с методом GET к урлу /phone/111 с пустым телом запроса
	req, err := http.NewRequest("GET", "/phone/111", nil)
	require.NoError(t, err)

	// Создаем ResponseRecorder для записи ответа сервера
	rr := httptest.NewRecorder()

	// Делаем запрос к нашему серверу
	server.ServeHTTP(rr, req)

	// Проверяем код ответа сервера
	assert.Equal(t, http.StatusOK, rr.Code)

	// Проверяем содержимое ответа
	expected := "name: Alex, phone: 111"
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetNotFound(t *testing.T) {
	// Создаем наш веб-сервер с фейковым хранилищем
	storage := newFakeStorage()
	server := New(storage)

	// Создаем новый HTTP-запрос с методом GET к урлу /phone/3333 с пустым телом запроса
	req, err := http.NewRequest("GET", "/phone/3333", nil)
	require.NoError(t, err)

	// Создаем ResponseRecorder для записи ответа сервера
	rr := httptest.NewRecorder()

	// Делаем запрос к нашему серверу
	server.ServeHTTP(rr, req)

	// Проверяем код ответа сервера
	assert.Equal(t, http.StatusNotFound, rr.Code)

	// Проверяем содержимое ответа
	expected := "Not found"
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetWrongMethod(t *testing.T) {
	// Создаем наш веб-сервер с фейковым хранилищем
	storage := newFakeStorage()
	server := New(storage)

	// Создаем новый HTTP-запрос с методом GET к урлу /phone с пустым телом запроса
	req, err := http.NewRequest("POST", "/phone/111", nil)
	require.NoError(t, err)

	// Создаем ResponseRecorder для записи ответа сервера
	rr := httptest.NewRecorder()

	// Делаем запрос к нашему серверу
	server.ServeHTTP(rr, req)

	// Проверяем код ответа сервера
	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
}
