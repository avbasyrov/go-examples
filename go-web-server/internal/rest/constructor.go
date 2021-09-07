package rest

import (
	"go-examples/go-web-server/internal/interfaces"
)

type Rest struct {
	storage interfaces.Storage // наше хранилище списка телефонов
}

func New(storage interfaces.Storage) *Rest {
	return &Rest{
		storage: storage,
	}
}
