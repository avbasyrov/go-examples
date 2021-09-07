package rest

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRest_ServeHTTP(t *testing.T) {
	server := &Rest{storage: newFakeStorage()}

	t.Run("запрос к неправильному урлу", func(t *testing.T) {
		// Создаем структуру HTTP-запроса
		request, err := http.NewRequest("GET", "/wrong_url", nil)
		require.NoError(t, err)

		// Recorder для записи ответа сервера
		recorder := httptest.NewRecorder()

		// Делаем запрос
		server.ServeHTTP(recorder, request)

		// Ожидаем код ответа сервера 400 (Bad request)
		assert.Equal(t, recorder.Code, http.StatusBadRequest)
	})

	t.Run("запрос неправильным методом", func(t *testing.T) {
		// Создаем структуру HTTP-запроса
		request, err := http.NewRequest("POST", "/list", nil)
		require.NoError(t, err)

		// Recorder для записи ответа сервера
		recorder := httptest.NewRecorder()

		// Делаем запрос
		server.ServeHTTP(recorder, request)

		// Ожидаем код ответа сервера 405 (Method not allowed)
		assert.Equal(t, recorder.Code, http.StatusMethodNotAllowed)
	})

	t.Run("успешный запрос списка", func(t *testing.T) {
		// Создаем структуру HTTP-запроса
		request, err := http.NewRequest("GET", "/list", nil)
		require.NoError(t, err)

		// Recorder для записи ответа сервера
		recorder := httptest.NewRecorder()

		// Делаем запрос
		server.ServeHTTP(recorder, request)

		// Ожидаем код ответа сервера 200 (OK)
		assert.Equal(t, recorder.Code, http.StatusOK)
	})
}
