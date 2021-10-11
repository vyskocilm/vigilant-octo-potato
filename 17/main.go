package main

import (
    "fmt"
    "strings"
    "reflect"
)

func typeAssert[T any](a T) {
    switch b := interface{}(a).(type) {
    case int:
        fmt.Printf("a=%+v INT, b=%+v INT\n", a, b)
    default:
        fmt.Printf("a=%+v %T, b=%+v %T\n", a, a, b, b)
    }
    v :=  reflect.ValueOf(interface{}(a))
    fmt.Printf("reflect.ValueOf = %s %T\n", v.String(), v)
}

func main() {
    typeAssert(11)
    typeAssert("hello")
    typeAssert([]string{"hel", "lo"})
    var sb strings.Builder
    typeAssert(&sb)
}
