//
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#constraint-type-inference
//
package main

import "fmt"

type number interface {
    ~int | ~uint
}

// Double returns a new slice that contains all the elements of s, doubled.
func Double[E number](s []E) []E {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v + v
	}
	return r
}

func main0() {
    type MySlice []int
    ret := Double((MySlice{1}))
    // // The type of V1 will be []int, not MySlice.
    // Here we are using function argument type inference,
    // but not constraint type inference.
    fmt.Printf("func argument inference: %+v %T\n", ret, ret)
}

// slice of any type
type SC[E any] interface {
	[]E // non-interface type constraint element
}

// DoubleDefined returns a new slice that contains the elements of s,
// doubled, and also has the same type as s.
func DoubleDefined[S SC[E], E number](s S) S {
	// Note that here we pass S to make, where above we passed []E.
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v + v
	}
	return r
}


func main1() {
    type MySlice []int
    // // The type of V2 will be MySlice.
    ret := DoubleDefined[MySlice, int](MySlice{1})
    ret2 := DoubleDefined(MySlice{1})
    // The type of V1 will be []int, not MySlice.
    // Here we are using function argument type inference,
    // but not constraint type inference.
    fmt.Printf("constraint type inference: %+v %T\n", ret, ret)
    fmt.Printf("constraint type inference (auto): %+v %T\n", ret2, ret2)
}

func main() {
    main0()
    main1()
}
