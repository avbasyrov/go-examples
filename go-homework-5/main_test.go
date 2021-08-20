package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate_Three(t *testing.T) {
	calculator := NewCalculator()
	result := calculator.Calc(3)

	assert.Len(t, result, 3)
	assert.Equal(t, 1, result[0])
	assert.Equal(t, 2, result[1])
	assert.Equal(t, 6, result[2])
}
