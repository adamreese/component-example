package main

import (
	"os"

	"github.com/adamreese/component-example/adder/docs/adder/v0.1.0/add"
)

func init() {
	// never panics
	panic("oh, snap")
}

func init() {
	add.Exports.Add = func(a, b int32) int32 {
		os.Stdout.WriteString("hello") // this causes a panic

		return a + b
	}
}

func main() {}
