package storage

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetByNumber_NotFound(t *testing.T) {
	storage := New()
	_, err := storage.GetByNumber("111")
	assert.Error(t, err)
}

func TestGetByNumber_Success(t *testing.T) {
	storage := New()
	phone, err := storage.GetByNumber("74212340077")
	require.NoError(t, err)

	assert.Equal(t, "74212340077", phone.Number)
	assert.Equal(t, "Ivan", phone.Name)
}
