package main

import "github.com/adamreese/component-example/subtractor/docs/subtractor/v0.1.0/subtract"

func init() {
	subtract.Exports.Subtract = func(a, b int32) int32 {
		return a - b
	}
}

func main() {}
