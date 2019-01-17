package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {

	testConfig, err := NewConfig()

	assert.NoError(t, err)
	assert.Equal(t, "8080", testConfig.Port, "Port should be string of 8080")
	assert.IsType(t, "8080", testConfig.Port)

}

func TestNewConfigWithEnvValues(t *testing.T) {

	port := "9000"
	os.Setenv("SERVICE_PORT", port)

	testConfig, err := NewConfig()

	assert.NoError(t, err)
	assert.Equal(t, port, testConfig.Port)

}

func TestNewConfigWithInvalidPortValue(t *testing.T) {

	port := "22"
	os.Setenv("SERVICE_PORT", port)

	_, err := NewConfig()

	assert.Error(t, err)

	port = "90000"
	os.Setenv("SERVICE_PORT", port)

	_, err = NewConfig()
	assert.Error(t, err)

	port = "ABCDEFG"
	os.Setenv("SERVICE_PORT", port)

	_, err = NewConfig()
	assert.Error(t, err)
}
