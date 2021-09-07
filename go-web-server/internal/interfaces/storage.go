package interfaces

type Storage interface {
	GetByNumber(phone string) (Phone, error)
	List() []Phone
}
