package main

// generic functions in existing Go
// 
//$ go doc builtin.make
// package builtin // import "builtin"
//
// func append(slice []Type, elems ...Type) []Type
// func copy(dst, src []Type) int
// func delete(m map[Type]Type1, key Type)
// func make(t Type, size ...IntegerType) Type
// func new(Type) *Type
// 

// not avialable outside of builtin
func newInt() *int { return new(int) }
func newString() *string { return new(string) }
type klass struct {}
func newKlass() *klass { return &klass{} }
