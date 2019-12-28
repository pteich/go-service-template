package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {

	testConfig, err := New()
	assert.NoError(t, err)
	assert.Equal(t, ":8000", testConfig.ListenAddr)
	assert.IsType(t, ":8000", testConfig.ListenAddr)
	assert.Equal(t, "debug", testConfig.LogLevel)
	assert.True(t, testConfig.LogOutputConsole)
}

func TestNewConfigWithEnvValues(t *testing.T) {

	listenAddr := "9000"
	logLevel := "info"

	os.Setenv("LISTEN_ADDR", listenAddr)
	os.Setenv("LOG_LEVEL", logLevel)
	os.Setenv("LOG_OUTPUT_CONSOLE", "false")

	testConfig, err := New()
	assert.NoError(t, err)
	assert.Equal(t, listenAddr, testConfig.ListenAddr)
	assert.Equal(t, logLevel, testConfig.LogLevel)
	assert.False(t, testConfig.LogOutputConsole)

}
