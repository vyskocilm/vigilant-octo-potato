//
// Type constraints based on interfaces
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#constraints
//
package main

import (
    "fmt"
    "strings"
)

type Stringer interface {
    String() string
}

func sprintf[T Stringer](value T) string {
    return value.String()
}
// 👆 CAN be done via interface with a better ergonomy (Go compiler
//    automatically reference/dereference)
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#values-of-type-parameters-are-not-boxed

func sprintx[T Stringer](slice []T) []string {
    ret := make([]string, len(slice))
    for idx, s := range slice {
        ret[idx] = s.String()
    }
    return ret
}
// 👆 CAN'T be done better via interfaces, []Stringer != []*strings.Builder
// cannot use foo (type []strings.Builder) as type []Stringer in argument to fun
// cannot use foo (type []*strings.Builder) as type []Stringer in argument to fun

func main() {
    var b strings.Builder
    b.WriteString("hello, generics!")
    fmt.Printf("sprintf=%s\n", sprintf(&b))
    fmt.Printf("sprintx=%s\n", sprintx([]*strings.Builder{&b}))
}
