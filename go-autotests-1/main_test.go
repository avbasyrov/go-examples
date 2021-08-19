package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCostError(t *testing.T) {
	item := &PortfolioItem{
		Amount: -1,
		Price:  -1,
	}
	_, err := item.Cost()
	assert.Error(t, err, "некорректные данные")
}

func TestCost(t *testing.T) {
	tests := []struct {
		item *PortfolioItem
		want float32
	}{
		{item: &PortfolioItem{Amount: 2, Price: 11}, want: 22},
		{item: &PortfolioItem{Amount: 3, Price: 10}, want: 30},
		{item: &PortfolioItem{Amount: 10, Price: 2}, want: 20},
	}

	for _, test := range tests {
		cost, err := test.item.Cost()
		require.NoError(t, err)
		assert.Equal(t, test.want, cost)
	}
}
