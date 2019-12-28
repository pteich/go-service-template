package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	Config
	StringValue string `env:"STRING_TEST"`
	BoolValue   bool   `env:"BOOL_TEST"`
	IntValue    int    `env:"INT_TEST"`
}

func TestParseEnv(t *testing.T) {

	confExpected := TestConfig{
		StringValue: "foo",
		BoolValue:   true,
		IntValue:    222,
	}

	defer os.Clearenv()

	os.Setenv("STRING_TEST", "foo")
	os.Setenv("BOOL_TEST", "true")
	os.Setenv("INT_TEST", "222")

	conf := TestConfig{}

	err := ParseEnv(&conf)
	assert.NoError(t, err)
	assert.Equal(t, confExpected, conf)
}

func TestParseEnvDefaults(t *testing.T) {

	conf := TestConfig{
		StringValue: "foo",
		BoolValue:   true,
		IntValue:    111,
	}

	confExpected := TestConfig{
		Config:      Config{},
		StringValue: "bar",
		BoolValue:   false,
		IntValue:    111,
	}

	defer os.Clearenv()

	os.Setenv("STRING_TEST", "bar")
	os.Setenv("BOOL_TEST", "false")

	ParseEnv(&conf)

	assert.Equal(t, confExpected, conf)

}
