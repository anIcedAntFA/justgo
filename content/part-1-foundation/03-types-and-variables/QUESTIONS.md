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

<details>
<summary>4. For <code>s := "café"</code>, what are <code>len(s)</code> and <code>len([]rune(s))</code>, and why do they differ? What does indexing <code>s[0]</code> return?</summary>

`len(s)` is **5** — it counts UTF-8 **bytes**, and `é` is 2 bytes. `len([]rune(s))`
is **4** — the number of Unicode **code points** (c, a, f, é). Indexing `s[0]` returns
a **`byte`** (`uint8`), not a one-character string — and indexing into a multi-byte
character would hand you just one of its bytes.

</details>

<details>
<summary>5. What are the three "views" of text in Go, and when do you use each?</summary>

- **`string`** — immutable UTF-8 bytes; the default for passing around and comparing.
- **`[]byte`** — a mutable copy of the bytes; for I/O, buffers, editing in place.
- **`[]rune`** — one `int32` per code point; only when you need the _N-th character_
  (reverse, count, index by character).

Converting `string ↔ []byte` / `[]rune` **copies**, because strings are immutable.

</details>

<details>
<summary>6. Why is <code>s += part</code> inside a loop a problem, and what do you use instead?</summary>

Strings are immutable, so each `+=` allocates a **whole new string** — O(n²) work and
lots of garbage. Use **`strings.Builder`** (`WriteString`/`WriteByte`, then
`.String()`), which writes into one growing buffer. Its zero value is ready to use —
no constructor needed.

</details>

<details>
<summary>7. Both <code>strconv.Itoa(n)</code> and <code>fmt.Sprintf("%d", n)</code> produce <code>"42"</code>. When would you prefer <code>strconv</code>?</summary>

Prefer `strconv` for simple conversions and hot paths — it's faster and allocation-light
because it neither parses a format string nor uses reflection. Reach for `fmt.Sprintf`
when interpolating several values into one message. And unlike JS's `parseInt`/`Number`
(which yield `NaN`), every `strconv` parse returns an explicit `error` to handle.

</details>
