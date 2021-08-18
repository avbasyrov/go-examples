package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestConvertToYaml(t *testing.T) {
	jsonBlob := []byte(`[{"name": "Igor", "info":{"age": 33}}]`)

	output := &bytes.Buffer{}

	err := ConvertToYaml(jsonBlob, output)
	require.NoError(t, err)

	decoder := yaml.NewDecoder(output)

	var result []struct {
		Name string `yaml:"name"`
		Data struct {
			Age int `yaml:"age"`
		} `yaml:"info"`
	}
	err = decoder.Decode(&result)
	require.NoError(t, err)
	require.Len(t, result, 1)
	assert.Equal(t, "Igor", result[0].Name)
	assert.Equal(t, 33, result[0].Data.Age)
}
