package main

import (
	"fmt"

	"github.com/penDerGraft/js.go/emitters"
)

func main() {
	e := emitters.New()

	e.On("evt", func(args ...interface{}) {
		fmt.Printf("first event callback called with %s, %s \n", args...)
	})

	e.On("evt", func(args ...interface{}) {
		fmt.Printf("second event callback called with %s, %s \n", args...)
	})

	e.Emit("evt", "2nd arg", "3rd arg")
}
