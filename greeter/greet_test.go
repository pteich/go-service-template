package greeter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGreet_SayHello(t *testing.T) {

	greetingService := NewGreet()

	greeting := greetingService.SayHello("World")
	assert.Equal(t, "Hello World!", greeting)

}
