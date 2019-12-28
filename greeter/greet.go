package greeter

import "fmt"

type Greet struct {
}

func NewGreet() *Greet {
	return &Greet{}
}

func (greet *Greet) SayHello(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
