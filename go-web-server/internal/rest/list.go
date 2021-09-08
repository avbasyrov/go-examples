package rest

import (
	"fmt"
	"net/http"
)

func (r *Rest) list(w http.ResponseWriter, _ *http.Request) {
	// Отдаем ответ 200 OK
	w.WriteHeader(http.StatusOK)

	phones := r.storage.List()
	for _, phone := range phones {
		_, _ = fmt.Fprintf(w, "name: %s, phone: %s; ", phone.Name, phone.Number)
	}
}
