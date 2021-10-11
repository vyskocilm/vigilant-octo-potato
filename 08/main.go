//
// Type set and interface in constraint
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#both-elements-and-methods-in-constraints
//
package main

import "strconv"

type IntString interface {
    ~int | ~int32
    String() string
}

func sprint[T IntString](value T) string {
    return value.String()
}

type myInt int
func (i myInt) String() string { return strconv.Itoa(int(i)) }

type myInt16 int16
func (i myInt16) String() string { return strconv.Itoa(int(i)) }

func main() {
    sprint(myInt(42))
    /*
    sprint[int](42)     // does not compile, int has no String method
    sprint(myInt16(42)) // does not compile, int16 not in typeset
    */
}
