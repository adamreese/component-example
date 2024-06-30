package main

import "github.com/adamreese/component-example/adder/docs/adder/add"

func init() {
	add.Exports.Add = func(a, b uint32) uint32 {
		return a + b
	}
}

func main() {}
