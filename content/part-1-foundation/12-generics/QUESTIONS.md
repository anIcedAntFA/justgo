# Chapter 12 — Questions

Recall questions. Try to answer out loud before unfolding. If you can't, re-read the
matching section of [`README.md`](./README.md).

<details>
<summary>1. What two bad options did Go developers have before generics, and what does a generic function give them instead?</summary>

Either **duplicate the code** — one `MaxInt`, `MaxFloat`, `MaxString`… per type — or
use **`any` (`interface{}`) and lose type safety**, needing type assertions and risking
runtime panics. A generic function (`func Max[T cmp.Ordered](a, b T) T`) is written
**once**, keeps **full compile-time type safety**, and — thanks to inference — reads
almost like ordinary code.

</details>

<details>
<summary>2. Why does Go use square brackets <code>[T]</code> for type parameters instead of angle brackets <code>&lt;T&gt;</code> like TS?</summary>

Angle brackets are **ambiguous to parse**: `f<T>(true)` could mean "call `f<T>` with
`(true)`" or "compare `f < T`, then `(true)`". Go sidesteps this with `[]`. Go also
**requires a constraint on every type parameter**, which resolves the remaining
array-vs-type-parameter ambiguity.

</details>

<details>
<summary>3. Name the three built-in constraints and the operation each unlocks. What does <code>~</code> do?</summary>

- `any` (alias for `interface{}`) — anything, but only type-agnostic operations (store,
  pass, print).
- `comparable` — `==` and `!=` (and use as a map key).
- `cmp.Ordered` (from `cmp`, Go 1.21) — `<`, `<=`, `>`, `>=`; covers ints, floats, strings.

`~T` in a constraint means "all types whose **underlying type** is `T`", so **defined
types** like `type Celsius float64` satisfy `~float64`. Without `~`, only the exact
predeclared types match.

</details>

<details>
<summary>4. What is a "type set", and what are the rules for a union like <code>int | float64</code>?</summary>

A constraint interface's **type set** is the collection of types that satisfy it. A
**union** `t1 | t2 | …` is the union of the terms' type sets, and inside the function you
may use only operations that **all** listed types support (e.g. `+`). Restrictions: the
non-interface terms must be **disjoint**, and a union with more than one term **cannot
contain methods** or `comparable`. In a `~T` term, `T`'s underlying type must be itself
and `T` can't be an interface (so `~error` is illegal).

</details>

<details>
<summary>5. In <code>func (s *Stack[T]) Push(item T)</code>, why is <code>[T]</code> repeated on the receiver? Can a method add its own type parameter?</summary>

The receiver must **name the type parameters** of its generic base type so the method
body can use them — hence `*Stack[T]`. On **Go 1.26** (this repo's target) a method may
**not** declare its own _extra_ type parameters; only the receiver's. (**Since Go 1.27**,
methods on _concrete_ types may declare their own type parameters — but interface methods
still may not.)

</details>

<details>
<summary>6. How do you produce the "zero value" of a type parameter <code>T</code>, and why can't you just write <code>return nil</code> or <code>return 0</code>?</summary>

Use **`var zero T`** — it yields the correct zero value for whatever `T` is. You can't
write `return nil` (only valid for pointers, slices, maps, …) or `return 0` (only numeric)
because `T` is _arbitrary_; the literal that fits one type is a compile error for another.
`var zero T` is the one form that works for every `T`.

</details>

<details>
<summary>7. When does type inference fail? Is a typed empty slice like <code>[]int{}</code> a problem?</summary>

Inference reads the type off the **arguments**, so a typed empty slice is **fine** —
`Sum([]int{})` infers `T = int` from the `[]int` type; length is irrelevant. Inference
fails only when the type parameter **doesn't appear in the ordinary parameters** — most
commonly when it appears only in the **return type** (`func New[T any]() *Stack[T]`), so
there's nothing to infer from. Then you instantiate explicitly: `New[int]()`.

</details>

<details>
<summary>8. The decision: interface, generic, or concrete? Give the rule of thumb.</summary>

- Call **methods** on the value (you need _behavior_) → **interface**.
- Store / move / transform values of **varying types** without calling type-specific
  methods (you need _type flexibility_) → **generics**.
- Only ever one type → **concrete**.

Rule of thumb: if the body calls methods, reach for an interface; if it just shuffles
values around, a generic fits. And **start concrete** — extract a generic only when real
duplication across types appears.

</details>

<details>
<summary>9. How does Go implement generics under the hood, and what's the practical performance takeaway?</summary>

**GC-shape stenciling plus dictionaries**: one instantiation per "GC shape" (all pointer
types share **one** shape), with a runtime **dictionary** carrying the exact type
arguments. It's a middle ground — not C++/Rust monomorphization (one copy per type) and
not fully boxed generics. Practically: almost always faster than `any` + assertions;
usually negligible overhead for value types; occasionally a touch slower for pointer types
(dictionary indirection) and slightly larger binaries. Don't worry about it for normal
code.

</details>

<details>
<summary>10. As of Go 1.26, where do the ready-made generics live, and is <code>golang.org/x/exp/constraints</code> still needed?</summary>

The stdlib: **`cmp`** (`Ordered`, `Compare`, `Less`, `Or`), **`slices`** (`Sort`,
`Contains`, `Max`, `Clone`, `Equal`, `Collect`, …), and **`maps`** (`Clone`, `Copy`,
`Equal`, `Keys`, `Values`). `x/exp/constraints` is now **mostly redundant** —
`constraints.Ordered` is literally `= cmp.Ordered`; import it only for `Signed`,
`Unsigned`, `Integer`, `Float`, or `Complex`. Note `slices` has **no** `Map`/`Filter`/
`Reduce` — those you still write by hand.

</details>

<details>
<summary>11. What changed about <code>comparable</code> in Go 1.20?</summary>

Before Go 1.20, `comparable` accepted only **strictly comparable** types. Since Go 1.20,
ordinary **interface types satisfy `comparable`** too (even though comparing them can
**panic at runtime**). That's why `Set[any]` and `map[any]V` are legal now — the
constraint is satisfied, with the runtime-panic risk moved onto you.

</details>
