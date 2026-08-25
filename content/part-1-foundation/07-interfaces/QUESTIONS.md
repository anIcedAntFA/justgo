# Chapter 07 — Questions

Recall questions. Try to answer out loud before unfolding. If you can't, re-read the
matching section of [`README.md`](./README.md).

<details>
<summary>1. How does a type "satisfy" an interface in Go, and how is that different from TypeScript?</summary>

**Implicitly.** A type satisfies an interface simply by having all of the interface's
methods — there is no `implements` keyword. TypeScript uses explicit, nominal
declaration (`class Rectangle implements Shape`); Go uses **structural** satisfaction
checked by the compiler. The consequence: you can satisfy an interface you don't own
(e.g. `io.Reader`), and an interface can be defined where it's _used_, not where the
type is _defined_.

</details>

<details>
<summary>2. What is <code>any</code>, and how does it relate to <code>interface{}</code>?</summary>

`any` is a **predeclared alias** for the empty interface: `type any = interface{}`,
added in Go 1.18. It is equivalent to `interface{}` in every way — same method set,
same behavior in assertions and type switches. Because an interface with no methods is
satisfied by _every_ type, an `any` value can hold anything. Use it sparingly; in
modern Go, generics often replace what `any` used to do.

</details>

<details>
<summary>3. What's the difference between <code>s := i.(string)</code> and <code>s, ok := i.(string)</code>?</summary>

Both are type assertions. The single-value form `s := i.(string)` **panics** if `i`'s
dynamic type isn't `string`. The two-value comma-ok form `s, ok := i.(string)` never
panics: `ok` reports whether the assertion succeeded. On failure, `s` is the **zero
value** of the asserted type (`""` for string, `0` for int) — not undefined. Prefer the
comma-ok form unless the type is guaranteed.

</details>

<details>
<summary>4. When would you use a type switch instead of chained type assertions?</summary>

When you need to handle **many** possible dynamic types. `switch v := i.(type)` gives
you, in each `case`, a `v` already narrowed to that case's type. It's cleaner than a
chain of `if s, ok := i.(string); ok { ... }` blocks, supports a `case nil`, and a
`default`. Note `.(type)` is legal _only_ inside a type switch.

</details>

<details>
<summary>5. Write the definitions of <code>io.Reader</code> and <code>io.Writer</code>. Why are one-method interfaces such a big deal?</summary>

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

One method each, yet hundreds of stdlib types satisfy them (`os.File`, `bytes.Buffer`,
`strings.Reader`, `net.Conn`, `os.Stdout`, `http.ResponseWriter`, …). Because they
share the interface, functions like `io.Copy(dst Writer, src Reader)` connect any
source to any destination. "The bigger the interface, the weaker the abstraction" —
small interfaces are easy to satisfy, mock, and compose.

</details>

<details>
<summary>6. What is interface embedding? Give a standard-library example.</summary>

An interface can list other interfaces as elements; its method set becomes the union.
The stdlib builds compound interfaces this way:

```go
type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```

A type satisfies `io.ReadWriteCloser` by having `Read`, `Write`, and `Close`. The rule:
when embedding, methods with the same name must have **identical signatures**. This
keeps interfaces small at the point of definition while letting callers ask for exactly
the combined behavior they need.

</details>

<details>
<summary>7. What is the <code>error</code> interface, and how do you make your own type an error?</summary>

`error` is just an interface with one method: `type error interface { Error() string }`.
Any type with an `Error() string` method _is_ an error. You make a custom error by
implementing that method (idiomatically on a pointer receiver):

```go
func (e *ValidationError) Error() string { ... }
```

Functions should keep returning the `error` interface in their signature, never the
concrete `*ValidationError` — that both keeps callers decoupled and avoids the typed-nil
trap (see Q9).

</details>

<details>
<summary>8. What is <code>fmt.Stringer</code>, and when does <code>fmt</code> call it?</summary>

`type Stringer interface { String() string }`. When you print a value with `fmt`
(e.g. `fmt.Println`), `fmt` checks whether it implements `Stringer` and, if so, uses
`String()` to render it. It's the interface-based equivalent of overriding `toString()`
in JS. (If a type implements _both_ `error` and `Stringer`, `fmt` uses `Error()` first.)

</details>

<details>
<summary>9. Why can an <code>error</code> value be non-nil even when the underlying pointer is nil?</summary>

An interface value is a pair: a **type** and a **value**. It equals `nil` only when
_both_ are nil. If you assign a nil `*MyError` pointer to an `error`, the interface's
type part is `*MyError` (non-nil), so the interface itself is **non-nil** even though
the pointer inside is nil — and `if err != nil` is unexpectedly true. The fix: return
the literal `nil`, never a typed nil pointer, and keep function signatures returning
`error` rather than a concrete pointer type.

</details>

<details>
<summary>10. "Accept interfaces, return structs" — is it a law? And when should you define an interface at all?</summary>

It's a **heuristic**, not a law. Accepting interfaces makes inputs flexible; returning
concrete types keeps outputs predictable. But the stdlib deliberately _returns_
interfaces when the point is to hide/swap implementations — `net.Dial` → `net.Conn`,
`crc32.NewIEEE` → `hash.Hash32`. On defining interfaces: keep them small (1–3 methods),
define them in the **consuming** package, and don't create them speculatively — write
concrete code first and extract an interface when a second implementation or a test mock
actually requires it.

</details>
