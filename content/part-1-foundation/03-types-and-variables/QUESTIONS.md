# Chapter 03 — Questions

Recall questions. Try to answer out loud before unfolding. If you can't, re-read the
matching section of [`README.md`](./README.md).

<details>
<summary>1. Go has no <code>undefined</code> and no <code>null</code> for most types. What does it have instead, and what is the zero value of <code>int</code>, <code>string</code>, <code>bool</code>, and a pointer?</summary>

Go gives every type a **zero value** — the value a variable holds the instant it's
declared without initialization. `int` → `0`, `string` → `""` (empty, _not_ nil),
`bool` → `false`, a pointer → `nil`. A declared variable is always usable; there is
no "undefined" state to guard against.

</details>

<details>
<summary>2. What's the difference between <code>var x = 5</code>, <code>var x int = 5</code>, and <code>x := 5</code>? When can you <em>not</em> use <code>:=</code>?</summary>

All three create an `int`. `var x int = 5` is explicit; `var x = 5` infers the type;
`x := 5` is the short form (infer + declare). `:=` only works **inside a function** —
at package level you must use `var`.

</details>

<details>
<summary>3. Why does <code>var f float64 = 3; var i int = f</code> not compile? What does that tell you about Go vs JS?</summary>

Go has **no implicit type conversion**. Assigning a `float64` to an `int` needs an
explicit `int(f)`. Unlike JS, Go never coerces numeric types silently — mixing types
is a compile error, which catches a whole class of bugs at build time.

</details>
