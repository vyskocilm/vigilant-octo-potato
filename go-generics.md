<!-- bg=black fg=green -->
# Go generics

5 Oct 2021
Summary: Go generics

Michal Vyskocil
BE Developer, Platform Team
michal.vyskocil@90poe.io

---
<!-- bg=black fg=green -->

## Go road to generics

 * Go was released without generics in 2009
 * Since then it was the most requested feature
 * First attempt in 2010
 * Official proposal submitted 12th January 2021: https://github.com/golang/go/issues/43651
 * Accepted in 10th February 2021

More info at https://go.dev/blog/generics-proposal

---
<!-- bg=black fg=green -->

## Can I try generics?

 * Generics are still a moving target
 * In a past it was https://github.com/golang/go/blob/dev.go2go/README.go2go.md
 * Or https://go2goplay.golang.org/
 * Can be tried in go 1.17 as `go run -gcflags=-G=3 main.go`

---
<!-- bg=black fg=green -->

## Can I really try generics?

 * Syntax was changed so many times
 * Most of go generics blogs are obsoleted now
 * dev.go2go and go 1.17 incompatible with latest proposal
 * https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md
 * go 1.18 (not yet released) with -gcflags=-G=3

---
<!-- bg=black fg=green -->

## Why it took so long?

 * Retrofitting of generics into existing language is not easy
 * Go team has a special relationship with a C++
 * Wanted to avoid C++ templates gotchas
 * And there are many opinions in the community
 * Paper https://arxiv.org/abs/2005.11710 described Go generics in a formal way
 * And there are many more opinions in the community

---
<!-- bg=black fg=green -->

## Why? It's complicated

 1. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#type-parameters
 2. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#constraints
 3. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#defining-constraints
 4. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#generic-types
 5. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#methods-may-not-take-additional-type-arguments
 6. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#type-sets-of-constraints
 7. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#comparable-types-in-constraints
 8. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#mutually-referencing-type-parameters
 9. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#type-inference
 11. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#type-unification
 12. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#function-argument-type-inference
 13. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#constraint-type-inference
 14. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#using-types-that-refer-to-themselves-in-constraints
 15. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#more-on-type-sets
 16. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#pointer-method-example
 17. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#composite-types-in-constraints
 18. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#type-parameters-in-type-sets
 19. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#type-conversions
 20. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#untyped-constants
 21. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#interface-types-in-union-elements
 22. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#empty-type-sets
 23. type assert + reflect package

---
<!-- bg=black fg=green -->
# Does Go needs generics?
### Should Go 2.0 support generics?
https://dave.cheney.net/2017/07/22/should-go-2-0-support-generics

> So my first answer is: Go should have some form of generics because it is a
> mainstream, imperative, block scoped language and it is expected these days.

---
<!-- bg=black fg=green -->

## Will Go change a lot?
### Why Go Getting Generics Will Not Change Idiomatic Go
http://www.jerf.org/iri/post/2955

> Generics in Go will give us data structures, tools to manipulate channels,
> maps, and slices in various generic ways beyond what we have now,
> possibly type-safe iterators, and a handful of other useful things. I
> suspect there are going to be some changes to the web frameworks, some
> moves on the numeric computation front, and some other similar things.
> I've got a few little places where I can avoid a bit of casting. But
> overall, what is "idiomatic" isn't going to change much.


---
<!-- bg=black fg=green -->

## Go already has generic functions

 * `append`
 * `copy`
 * `delete`
 * `make`
 * `new`

 Yet the language itself can't express them, those are runtime hacks!

---
<!-- bg=black fg=green -->

## Syntax

 `[Name constraint]`

 ```go
func new2[T any]() *T {
    return new(T)
}
type klass[T any] struct { value T }
func (k *klass[T]) foo() { }
 ```

 constraint is mandatory or compiler will complain
 `syntax error: missing type constraint`

---
<!-- bg=black fg=green -->

# And that's
# all folks'

---
<!-- bg=black fg=green -->

## Lets try some real code

 * 01/ no generics
 * 02/ generic new2
 * 03/ type sets in constraints
 * 04/ interface constraints
 * 05/ generic types
 * 06/ befare of int
 * 07/ comparable types
 * 08/ more on type sets

---
<!-- bg=black fg=green -->

## Lets try some real code

 * 09/ composition types
 * 10/ mutually referencing type parameters
 * 11/ type conversions and untyped constants
 * 12/ interface types in union
 * 13/ empty type sets
 * 14/ (function argument) type inference
 * 15/ constraint type inference
 * 16/ pointer method example
 * 17/ type asserts and a reflec package

### more examples
 * https://github.com/golang/go/tree/master/test/typeparam
 * https://github.com/mattn/go-generics-example/tree/main/constraints-chan

---
<!-- effect=matrix -->

# func [E any]magic(a T) []T {
#     return []T{a}
# }
