package rest

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (r *Rest) get(w http.ResponseWriter, request *http.Request) {
	// Отрезаем от строки запроса первые 7 символов ("/phone/") - остается только номер телефона
	phoneNumber := strings.TrimPrefix(request.URL.Path, "/phone/")

	log.Println(phoneNumber)

	phone, err := r.storage.GetByNumber(phoneNumber)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Not found"))
		return
	}

	w.WriteHeader(http.StatusOK) // Отдаем ответ 200 OK

	_, _ = fmt.Fprintf(w, "name: %s, phone: %s", phone.Name, phone.Number)
}
