//
// # Type parameters
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#type-parameters 
// operations on type any
// see https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#operations-permitted-for-any-type
//
package main

import "fmt"

/*
    Type parameter [T $constraint] - any means any type
*/
func new2[K any]() *K {
    return new(K)
}

type Bar[T any] struct{ v T }
func (b Bar[MyTypeParameterName]) bar() MyTypeParameterName { return b.v }

type klass struct{}

func main() {
    intp := new2[int]()         // compiler instantiate new2_int() function
    stringp := new2[string]()
    klassp := new2[klass]()

    fmt.Printf("inpt=%T %+v\n", intp, intp)
    fmt.Printf("stringp=%T %+v\n", stringp, stringp)
    fmt.Printf("klassp=%T %+v\n", klassp, klassp)
}
