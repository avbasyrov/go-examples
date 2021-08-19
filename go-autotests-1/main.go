package main // go-autotests-1

import (
	"errors"
	"fmt"
)

type PortfolioItem struct {
	Amount int
	Price  float32
}

func (p *PortfolioItem) Cost() (float32, error) {
	if p.Amount >= 0 && p.Price > 0 {
		return float32(p.Amount) * p.Price, nil
	}

	return 0, errors.New("некорректные данные")
}

func main() {
	item := &PortfolioItem{
		Amount: 2,
		Price:  11,
	}
	fmt.Println(item.Cost())
}
