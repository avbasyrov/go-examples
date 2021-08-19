package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type FakeItem struct {
	Amount int
	Price  float32
}

func (f *FakeItem) Cost() (float32, error) {
	return 0, nil
}
func NewTestItem(amount int, price float32) PortfolioCost {
	// Пока нет кода - используем временную заглушку
	return &FakeItem{Amount: amount, Price: price}
}

func TestCostError(t *testing.T) {
	item := NewTestItem(-1, -1)
	_, err := item.Cost()
	assert.Error(t, err, "некорректные данные")
}

func TestCost(t *testing.T) {
	tests := []struct {
		item PortfolioCost
		want float32
	}{
		{item: NewTestItem(2, 11), want: 22},
		{item: NewTestItem(3, 10), want: 30},
		{item: NewTestItem(10, 2), want: 20},
	}

	for _, test := range tests {
		cost, err := test.item.Cost()
		require.NoError(t, err)
		assert.Equal(t, test.want, cost)
	}
}
