//
// Type conversions
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#type-conversions
// Untyped constants
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#untyped-constants
//
package main

import "fmt"

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func convert[To, From integer](from From) To {
	to := To(from)
	if From(to) != from {
		panic("conversion out of range")
	}
	return to
}

func inc[T integer](a T) T {
    return a + 1
}

func main() {
    from  := int16(41)
    to := convert[int64](from)
    fmt.Printf("convert[%T](%d %T)=%d %T\n", to, from, from, to, to)
    fmt.Printf("add(%d %T)=%d %T\n", from, from, inc(from), inc(from))
}
