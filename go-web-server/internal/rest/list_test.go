package rest

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListSuccess(t *testing.T) {
	// Создаем наш веб-сервер с фейковым хранилищем
	storage := newFakeStorage()
	server := New(storage)

	// Создаем новый HTTP-запрос с методом GET к урлу /list с пустым телом запроса
	req, err := http.NewRequest("GET", "/list", nil)
	require.NoError(t, err)

	// Создаем ResponseRecorder для записи ответа сервера
	rr := httptest.NewRecorder()

	// Делаем запрос к нашему серверу
	server.ServeHTTP(rr, req)

	// Проверяем код ответа сервера
	assert.Equal(t, http.StatusOK, rr.Code)

	// Проверяем содержимое ответа
	expected := "name: Alex, phone: 111; name: Bro, phone: 222; "
	assert.Equal(t, expected, rr.Body.String())
}

func TestListWrongMethod(t *testing.T) {
	// Создаем наш веб-сервер с фейковым хранилищем
	storage := newFakeStorage()
	server := New(storage)

	// Создаем новый HTTP-запрос с методом GET к урлу /list с пустым телом запроса
	req, err := http.NewRequest("POST", "/list", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Создаем ResponseRecorder для записи ответа сервера
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.list)

	// Вызываем напрямую обработчик нашего сервера с запросом req
	handler.ServeHTTP(rr, req)

	// Проверяем код ответа сервера
	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
}
