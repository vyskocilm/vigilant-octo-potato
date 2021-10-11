//
// Type inference
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#type-inference
// Function argument type inference
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#function-argument-type-inference
//
package main

import (
    "fmt"
    "strconv"
)

// 👏 can infer everything
func arr[T interface{}](a T) []T {
    return []T{a}
}

func map2[T, E any](arr []T, fn func(a T) E) []E {
    ret := make([]E, len(arr))
    for idx, a := range arr {
        ret[idx] = fn(a)
    }
    return ret
}

type intish interface {
    ~int
}

func strFun[T intish](a T) string {
    return string(strconv.Itoa(int(a)))
}

func main() {
    fmt.Printf("arr(1)=%+v %T\n", arr(1), arr(1))
    fmt.Printf("arr(1)=%+v %T\n", arr[int16](1), arr[int16](1))

    /*
Function argument type inference is used with a function call to infer type arguments from non-type arguments. Function argument type inference is not used when a type is instantiated, and it is not used when a function is instantiated but not called.
    */

    // can be infered
    ret := map2([]int{1}, func(a int) string { return strconv.Itoa(a) })
    fmt.Printf("ret=%+v %T\n", ret, ret)

    // but sometimes Go can't
    // strFun[int] must be stated
    ret = map2([]int{1}, strFun)
    fmt.Printf("ret=%+v %T\n", ret, ret)
}
