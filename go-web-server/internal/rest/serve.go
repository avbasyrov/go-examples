package rest

import (
	"net/http"
	"strings"
)

func (r *Rest) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	// Наш контроллер ожидает только GET-запросы
	if request.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("Method not allowed"))
		return
	}

	if request.URL.Path == "/list" {
		r.list(w, request)
		return
	}

	if strings.HasPrefix(request.URL.Path, "/phone/") {
		r.get(w, request)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte("Bad request"))
	return
}
