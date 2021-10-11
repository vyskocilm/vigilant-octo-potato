//
// Generic containers: a surprise!
//
package main

import (
    "fmt"
    "strconv"
)

type vector[T any] []T

// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md
// 🤷 not sure why pointer is needed here, copy from Go Generics design document
func (v *vector[T]) Push(value T) {
    *v = append(*v, value)
}

func map2[T, V any](slice []T, fn func(T) V) []V {
    ret := make([]V, len(slice))
    for idx, val := range slice {
        ret[idx] = fn(val)
    }
    return ret
}

// you can type alias parametrized type
type ivector = vector[int]
// you can't provide partial implementation for a given type
/*
func (v *vector[int]) Push(value int) {
    panic("specialization")
}
/main.go:19:23: vector.Push redeclared in this block
        /home/mvyskocil/projects/90poe/generics/06/main.go:15:21: other declaration of Pus
*/

// and there are more thing forbidden with not that clear errors
/*
func (iv *ivector) Push(value int) {
    panic("specialization for ivector")
}
# command-line-arguments
./main.go:45:15: InitTextSym double init for "".(*vector[int]).Push
*/

func (iv *vector[int]) PushInt(value int) {
    panic("specialization for vector[int]")
}

func main() {
    v := ivector{}
    v.Push(11)
    v2 := map2(v, func(x int)string {return "hello: " + strconv.Itoa(x)})
    fmt.Printf("vf=%+v\n", v2)

    v.PushInt(42)

    vs := vector[string]{}
    vs.PushInt("42")            // 💣 this compiles
}










// explanation
// https://github.com/golang/go/issues/47419#issuecomment-887859803
/*
Definition 

func (b Bar[int]) bar() int { return b.v }

just means that for the body of the method, the type parameter is named int.
It's no different from writing `T`, `int` or MyAwesomeType there.
Ian Lance Taylor, Go core team

See list of ommissions
https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#omissions
 * No specialization. There is no way to write multiple versions of a generic function that are designed to work with specific type arguments.
*/
