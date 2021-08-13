package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVaccine_DiseaseProbability(t *testing.T) {
	sputnikV := NewSputnikV()
	assert.Equal(t, 0.03, sputnikV.DiseaseProbability())
}

func TestHuman_ApplyVaccine(t *testing.T) {
	phizer := NewPhizer()
	sputnikV := NewSputnikV()
	John := Human{}

	John.ApplyVaccine(phizer)

	assert.Implements(t, (*VaccineInterface)(nil), John.appliedVaccine)
	assert.NotEqual(t, sputnikV, John.appliedVaccine)
	assert.Equal(t, phizer, John.appliedVaccine)
	assert.Equal(t, 0.04, John.appliedVaccine.DiseaseProbability())
}
