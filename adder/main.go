package main

import "github.com/adamreese/component-example/adder/docs/adder/v0.1.0/add"

func init() {
	add.Exports.Add = func(a, b int32) int32 {
		return a + b
	}
}

func main() {}
