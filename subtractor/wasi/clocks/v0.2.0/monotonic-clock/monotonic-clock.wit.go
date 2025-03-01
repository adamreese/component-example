// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package monotonicclock represents the imported interface "wasi:clocks/monotonic-clock@0.2.0".
//
// WASI Monotonic Clock is a clock API intended to let users measure elapsed
// time.
//
// It is intended to be portable at least between Unix-family platforms and
// Windows.
//
// A monotonic clock is a clock which has an unspecified initial value, and
// successive reads of the clock will produce non-decreasing values.
//
// It is intended for measuring elapsed time.
package monotonicclock

import (
	"github.com/adamreese/component-example/subtractor/wasi/io/v0.2.0/poll"
	"github.com/ydnar/wasm-tools-go/cm"
)

// Instant represents the u64 "wasi:clocks/monotonic-clock@0.2.0#instant".
//
// An instant in time, in nanoseconds. An instant is relative to an
// unspecified initial value, and can only be compared to instances from
// the same monotonic-clock.
//
//	type instant = u64
type Instant uint64

// Duration represents the u64 "wasi:clocks/monotonic-clock@0.2.0#duration".
//
// A duration of time, in nanoseconds.
//
//	type duration = u64
type Duration uint64

// Now represents the imported function "now".
//
// Read the current value of the clock.
//
// The clock is monotonic, therefore calling this function repeatedly will
// produce a sequence of non-decreasing values.
//
//	now: func() -> instant
//
//go:nosplit
func Now() (result Instant) {
	result0 := wasmimport_Now()
	result = (Instant)((uint64)(result0))
	return
}

//go:wasmimport wasi:clocks/monotonic-clock@0.2.0 now
//go:noescape
func wasmimport_Now() (result0 uint64)

// Resolution represents the imported function "resolution".
//
// Query the resolution of the clock. Returns the duration of time
// corresponding to a clock tick.
//
//	resolution: func() -> duration
//
//go:nosplit
func Resolution() (result Duration) {
	result0 := wasmimport_Resolution()
	result = (Duration)((uint64)(result0))
	return
}

//go:wasmimport wasi:clocks/monotonic-clock@0.2.0 resolution
//go:noescape
func wasmimport_Resolution() (result0 uint64)

// SubscribeInstant represents the imported function "subscribe-instant".
//
// Create a `pollable` which will resolve once the specified instant
// occured.
//
//	subscribe-instant: func(when: instant) -> pollable
//
//go:nosplit
func SubscribeInstant(when Instant) (result poll.Pollable) {
	when0 := (uint64)(when)
	result0 := wasmimport_SubscribeInstant((uint64)(when0))
	result = cm.Reinterpret[poll.Pollable]((uint32)(result0))
	return
}

//go:wasmimport wasi:clocks/monotonic-clock@0.2.0 subscribe-instant
//go:noescape
func wasmimport_SubscribeInstant(when0 uint64) (result0 uint32)

// SubscribeDuration represents the imported function "subscribe-duration".
//
// Create a `pollable` which will resolve once the given duration has
// elapsed, starting at the time at which this function was called.
// occured.
//
//	subscribe-duration: func(when: duration) -> pollable
//
//go:nosplit
func SubscribeDuration(when Duration) (result poll.Pollable) {
	when0 := (uint64)(when)
	result0 := wasmimport_SubscribeDuration((uint64)(when0))
	result = cm.Reinterpret[poll.Pollable]((uint32)(result0))
	return
}

//go:wasmimport wasi:clocks/monotonic-clock@0.2.0 subscribe-duration
//go:noescape
func wasmimport_SubscribeDuration(when0 uint64) (result0 uint32)
