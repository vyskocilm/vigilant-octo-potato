//
// # Type set constraints
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#type-sets-of-constraints
// 
package main

import "fmt"

/*
func min[T any](a, b T) T {
    if a > b {
        return b
    }
    return a
}

func printMin[T any](a, b T) {
    res := min(a, b)
    fmt.Printf("min(%+v %T, %+v %T)=%+v %T\n", a, a, b, b, res, res)
}

func main() {
    printMin(1, 11)

type checking failed for main
main.go2:7:8: cannot compare a > b (operator > not defined for T)
*/

// constraints
// https://github.com/golang/go/blob/dev.go2go/src/cmd/go2go/testdata/go2path/src/constraints/constraints.go2
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#type-sets
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#union-constraint-element
// Ordered is a type constraint that matches int or float64
// An ordered type is one that supports the <, <=, >, and >= operators.
type Ordered interface {
    ~int | float64
}

// Approximation constraint element
// ~ makes it work even for defined types

func min[T Ordered](a, b T) T {
    if a > b {
        return b
    }
    return a
}

func printMin[T Ordered](a, b T) {
    res := min(a, b)
    fmt.Printf("min(%+v %T, %+v %T)=%+v %T\n", a, a, b, b, res, res)
}

type myInt int
type myFloat64 float64
type myFloat64Alias = float64

func main() {
    printMin(1, 11)
    printMin(1.0, float64(11))
    printMin(myInt(1), 11)
    printMin(myFloat64Alias(1.1), 11.1)
    /*
    printMin(uint8(1), uint8(11))   // uint8 not in type set
    printMin(1.0, 11)   // default type int of 11 does not match inferred type float64 for T

    type myInt int
    type myFloat64 float64

    printMin(myFloat64(1.1), 11.1)
    */
}
