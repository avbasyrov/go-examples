package storage

import (
	"errors"
	"go-examples/go-web-server/internal/interfaces"
)

func (h *HardCodedStorage) GetByNumber(phone string) (interfaces.Phone, error) {
	for _, p := range phones {
		if p.Number == phone {
			return p, nil
		}
	}

	return interfaces.Phone{}, errors.New("phone not found")
}
