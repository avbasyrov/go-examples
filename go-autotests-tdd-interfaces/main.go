package main // go-autotests-tdd-interfaces
// Показывать в IDE

type PortfolioItem struct {
	Amount int
	Price  float32
}

// Вместо реализации делаем сначала только интерфейс
type PortfolioCost interface {
	Cost() (float32, error)
}

// TODO: реализовать метод Cost() после тестов

func main() {
}
