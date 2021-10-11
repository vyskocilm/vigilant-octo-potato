//
// Type parameters in type sets
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#type-parameters-in-type-sets
package main

import "fmt"

// interface satisfied for any of string/[]byte or slice of T
type lener [T any] interface {
    ~string | []byte | []T
}

/*
func len2[T lener](value T) int {
    return len(value)
}
// cannot use generic type lener[T interface{}] without instantiation
*/

func len2[E any, T lener[E]](value T) int {
    return len(value)
}

// this works, but compiler can't infern type E as it exists only as a type parameter

func main() {

    fmt.Printf("len2 string=%d\n", len2[string]("hi"))
    fmt.Printf("len2 []byte=%d\n", len2[byte]([]byte("hi-five")))
    fmt.Printf("len2 []string=%d\n", len2[string]([]string{"1"}))
    fmt.Printf("len2 []int=%d\n", len2[int]([]int{1, 2, 3, 4}))

}

/*
I heard you like type constraints
So I put a type constraint inside your type constraint
type lener3[T any, E lener[T]] interface {
    Slice() []T
    Len([]T) int
}
*/


