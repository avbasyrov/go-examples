package storage

import "go-examples/go-web-server/internal/interfaces"

func (h *HardCodedStorage) List() []interfaces.Phone {
	return phones
}
