//
// Interface types in union elements
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#interface-types-in-union-elements
//
package main

import (
    "fmt"
    "strings"
)

// should work according design document, but does not
// cannot use fmt.Stringer in union (interface contains methods)
type stringish interface {
	string | fmt.Stringer
}

func str2[T stringish](a T) string {
    switch foo := a.(type) {
    default:
        return fmt.Sprintf("%+v %T", a, foo)
    }
}


func main() {
    fmt.Printf("str2=%s\n", "string")
    var sb strings.Builder
    sb.WriteString("strings.Builder")
    fmt.Printf("str2=%s\n", sb)
}

