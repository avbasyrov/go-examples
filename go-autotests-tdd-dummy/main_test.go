package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCostError(t *testing.T) {
	item := NewItem(-1, -1)
	_, err := item.Cost()
	assert.Error(t, err, "некорректные данные")
}

func TestCost(t *testing.T) {
	tests := []struct {
		item PortfolioCost
		want float32
	}{
		{item: NewItem(2, 11), want: 22},
		{item: NewItem(3, 10), want: 30},
		{item: NewItem(10, 2), want: 20},
	}

	for _, test := range tests {
		cost, err := test.item.Cost()
		require.NoError(t, err)
		assert.Equal(t, test.want, cost)
	}
}
