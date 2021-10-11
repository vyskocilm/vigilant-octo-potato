// 
// comparable biltin constraint
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#comparable-types-in-constraints
//
package main

import "fmt"

type foo struct {
}

func (f foo)Hash() uintptr {
    return uintptr(0x42)
}

// comparable is a builtin and valid for all Go types with == and != defined
// satisfied by types can be compared and with Hash method
type ComparableHasher interface {
    comparable
    Hash() uintptr
}

/*
// ImpossibleConstraint is a type constraint that no type can satisfy,
// because slice types are not comparable.
type ImpossibleConstraint interface {
	comparable
	[]int
}
*/

func printt[T ComparableHasher](value T) {
    fmt.Printf("value(hash=%d)=%+v %T\n", value.Hash(), value, value)
}

func main() {
    printt(foo{})
}
