// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

package add

// Exports represents the caller-defined exports from "docs:adder/add@0.1.0".
var Exports struct {
	// Add represents the caller-defined, exported function "add".
	//
	//	add: func(a: s32, b: s32) -> s32
	Add func(a int32, b int32) (result int32)
}
