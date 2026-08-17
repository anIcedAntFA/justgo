# Chapter 05 — Questions

Recall questions. Try to answer out loud before unfolding. If you can't, re-read the
matching section of [`README.md`](./README.md).

<details>
<summary>1. Write the signature of a function <code>divide</code> that returns a <code>float64</code> and an <code>error</code>. Why is this two-value shape called the "backbone" of Go?</summary>

`func divide(a, b float64) (float64, error)`. It's the backbone because Go has no
exceptions — a function that can fail returns its result **and** an `error` value, and
the caller checks the error immediately (`v, err := ...; if err != nil { ... }`).
Almost every fallible function in the standard library uses it, so the `value, err :=`
line is the most common line in real Go code.

</details>

<details>
<summary>2. What is the <code>(value, ok)</code> pattern, and why is it unambiguous where JS's <code>obj[key] === undefined</code> is not?</summary>

A second boolean return that reports existence rather than an error — e.g.
`age, ok := ages["Bob"]`. If the key is missing you get the zero value **and**
`ok == false`, so you can tell "missing" apart from "present but zero". JS's
`obj[key] === undefined` can't: a key that legitimately holds `undefined` looks
identical to a missing one. The same shape appears in type assertions
(`v, ok := x.(int)`) and channel receives (`v, ok := <-ch`).

</details>

<details>
<summary>3. What is a <em>naked return</em>, and what is the one situation where named returns are genuinely necessary?</summary>

A `return` with no arguments in a function with **named** return values — it returns
whatever those names currently hold. Named returns are genuinely necessary when a
`defer`-ed closure must **read or rewrite the result**, e.g. recovering from a panic
and assigning to the named `err`. Otherwise prefer explicit returns; avoid naked
returns in long functions where the returned values aren't obvious.

</details>

<details>
<summary>4. Write a variadic <code>sum</code>. How do you pass an existing <code>[]int</code> into it, and what's the one placement rule?</summary>

```go
func sum(nums ...int) int { /* nums is a []int */ }
```

Spread a slice with `...`: `sum(nums...)`. The rule: a variadic parameter must be the
**last** parameter — `func logf(level, format string, args ...any)` is fine,
`func bad(args ...int, name string)` is a compile error. (`...any` is the modern
spelling of `...interface{}`.)

</details>

<details>
<summary>5. Closures capture by reference. What did Go 1.22 change about <code>for</code> loop variables, and what did it <em>not</em> change?</summary>

Before Go 1.22 the loop variable was created once and shared, so closures capturing it
all saw the final value (the classic `3 3 3`). Go 1.22 gives **each iteration its own
copy**, so you get `0 1 2`. What it did **not** change: closures still capture by
reference in general — any time several closures close over the same mutable variable
(not just a loop variable), they all observe its latest value.

</details>

<details>
<summary>6. Go has no default parameters and no function overloading. What is the idiomatic replacement for "configurable struct with defaults", and which chapter concepts does it combine?</summary>

The **functional options pattern**: `NewServer(opts ...Option)` where
`type Option func(*Config)`, each `WithX(...)` returns a closure that mutates the
config, and the constructor starts from defaults then applies each option. It ties the
chapter together — variadic parameters, first-class function values, and closures — and
lets callers override only what they need without breaking existing call sites.

</details>

<details>
<summary>7. Since Go 1.23 you can <code>range</code> over a function. In one sentence, what <em>is</em> an iterator in Go, and what is the type <code>iter.Seq[V]</code>?</summary>

An iterator is just an ordinary first-class function that takes a `yield` callback and
calls it once per element — `iter.Seq[V]` is the alias
`func(yield func(V) bool)`. `yield` returns `false` when the ranging loop breaks early,
telling the iterator to stop. It's not new machinery — it's first-class functions plus
closures. (Full treatment in Chapter 10.)

</details>

<details>
<summary>8. Does Go do tail-call optimisation? What's the practical consequence for recursion?</summary>

No — Go does not perform tail-call optimisation, so deeply/unbounded recursion can
overflow the stack; prefer iteration for large depths. Goroutine stacks do start small
and grow dynamically, giving more headroom than most languages, but there's no TCO
guarantee to rely on.

</details>
