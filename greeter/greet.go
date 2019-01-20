package greeter

import "fmt"

type Greeter interface {
	SayHello(name string) string
}

type Greet struct {
}

func NewGreet() *Greet {
	return &Greet{}
}

func (greet *Greet) SayHello(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
