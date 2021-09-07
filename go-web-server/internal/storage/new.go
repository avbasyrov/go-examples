package storage

type HardCodedStorage struct{}

func New() *HardCodedStorage {
	return &HardCodedStorage{}
}
