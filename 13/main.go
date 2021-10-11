//
// Emty type sets
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#empty-type-sets
//
package main

import "fmt"

type emptySet interface {
    int
    fmt.Stringer
}

type stringish interface {
    ~string
    fmt.Stringer
}

func str1[T stringish](a T) string {
    return a.String()
}

type myString string
func (m myString) String() string { return string(m) }

func main() {
    //fmt.Printf("str1=%s\n", str1("hello"))
    fmt.Printf("str1=%s\n", str1(myString("myString")))

    fmt.Printf("int1=%d\n", 11)
}
