package rest

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-examples/go-web-server/internal/interfaces"
	"testing"
)

type fakeStorage struct{}

func newFakeStorage() *fakeStorage {
	return &fakeStorage{}
}

func (f *fakeStorage) GetByNumber(phone string) (interfaces.Phone, error) {
	if phone == "111" {
		return interfaces.Phone{
			Number: "111",
			Name:   "Alex",
		}, nil
	}

	return interfaces.Phone{}, errors.New("not found")
}

func (f *fakeStorage) List() []interfaces.Phone {
	return []interfaces.Phone{
		{Number: "111", Name: "Alex"},
		{Number: "222", Name: "Bro"},
	}
}

func TestNew(t *testing.T) {
	storage := newFakeStorage()
	rest := New(storage)

	assert.Equal(t, storage, rest.storage)
}
