//
// Generic containers
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#generic-typespackage main
//
package main

import (
    "fmt"
)

// bye bye slice tricks!
type vector[T any] []T

// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md
// bye bye slice tricks
func (v *vector[T]) Push(value T) {
    *v = append(*v, value)
}

// functional
func (v *vector[T]) Map(fn func(T) T) []T {
    ret := make([]T, len(*v))
    for idx, val := range *v {
        ret[idx] = fn(val)
    }
    return ret
}

// methods CAN'T have more type parameters than the struct
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#methods-may-not-take-additional-type-arguments
/*
func (v *vector[T, V]) Map2(fn func(T) V) []V {
    return V{}
}
*/
//./main.go:26:10: got 2 arguments but 1 type parameters
//./main.go:27:12: invalid composite literal type V

func main() {
    v := vector[interface{}]{}
    v.Push(1)
    v.Push("one")
    fmt.Printf("v=%+v %T\n", v, v)

    vi := vector[int]{}
    vi.Push(11)
    fmt.Printf("vi=%+v %T\n", vi, vi)

    vs := vector[string]{}
    vs.Push("11-hello")
    fmt.Printf("vs=%+v %T\n", vs, vs)

    /*
    vi2 := vi.Map(func(x int64)int64 {return x * 2})
    fmt.Printf("vi2=%+v\n", vi2)
    */
}
