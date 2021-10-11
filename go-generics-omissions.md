<!-- bg=black fg=green -->
# Omissions

---
<!-- bg=black fg=green -->

## No specialization.
    There is no way to write multiple versions of a generic function that are
    designed to work with specific type arguments.

## No metaprogramming.
    There is no way to write code that is executed at compile time to generate
    code to be executed at run time.

---
<!-- bg=black fg=green -->

## No higher level abstraction.
    There is no way to use a function with type arguments other than to call it
    or instantiate it. There is no way to use a generic type other than to
    instantiate it.

## No general type description.
    In order to use operators in a generic function, constraints list specific
    types, rather than describing the characteristics that a type must have.
    This is easy to understand but may be limiting at times.

---
<!-- bg=black fg=green -->

## No covariance or contravariance of function parameters.
    Aka no subtyping support

## No operator methods.
    You can write a generic container that is compile-time type-safe, but you
    can only access it with ordinary methods, not with syntax like c[k].

---
<!-- bg=black fg=green -->

## No currying.
    There is no way to partially instantiate a generic function or type, other
    than by using a helper function or a wrapper type. All type arguments must
    be either explicitly passed or inferred at instantiation time.

## No variadic type parameters.

    There is no support for variadic type parameters, which would permit
    writing a single generic function that takes different numbers of both type
    parameters and regular parameters.

---
<!-- bg=black fg=green -->

## No adaptors.
    There is no way for a constraint to define adaptors that could be used to
    support type arguments that do not already implement the constraint, such
    as, for example, defining an == operator in terms of an Equal method, or
    vice-versa.

## No parameterization on non-type values such as constants.
    This arises most obviously for arrays, where it might sometimes be
    convenient to write type Matrix[n int] [n][n]float64. It might also
    sometimes be useful to specify significant values for a container type,
    such as a default value for elements.
