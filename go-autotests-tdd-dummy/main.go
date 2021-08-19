package main // go-autotests-tdd-dummy
// Показывать в IDE

type PortfolioItem struct {
	Amount int
	Price  float32
}

type PortfolioCost interface {
	Cost() (float32, error)
}

func (p *PortfolioItem) Cost() (float32, error) {
	// TODO: реализовать код метода Cost() после тестов
	return 0, nil
}

func NewItem(amount int, price float32) PortfolioCost {
	return &PortfolioItem{Amount: amount, Price: price}
}

func main() {
}
