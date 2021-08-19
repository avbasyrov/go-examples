package main // go-autotests-2

type PortfolioItem struct {
	Amount int
	Price  float32
}

type PortfolioCost interface {
	Cost() (float32, error)
}

// TODO: реализовать метод Cost() после тестов

func main() {
}
