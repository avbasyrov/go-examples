package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	storage := New()
	assert.Equal(t, phones, storage.List())
}
