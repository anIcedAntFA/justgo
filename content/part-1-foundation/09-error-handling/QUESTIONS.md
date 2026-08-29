# Chapter 09 — Questions

Recall questions. Try to answer out loud before unfolding. If you can't, re-read the
matching section of [`README.md`](./README.md).

<details>
<summary>1. What <em>is</em> an error in Go, and how do you know a call succeeded?</summary>

An error is any value implementing the one-method `error` interface:
`interface { Error() string }`. By convention a fallible function returns an `error`
as its **last** return value. `err == nil` means success; a non-nil `err` means
failure. There's no separate exception channel — the error is an ordinary value you
check with `if err != nil`.

</details>

<details>
<summary>2. Why does Go use returned error values instead of try/catch exceptions?</summary>

To make error paths **visible and explicit**. With exceptions, any line might throw,
handling is disconnected from where the error occurred, and it's easy to forget a
`catch`. Returning errors as values means every fallible operation is marked at the
call site, the compiler/linter won't let you silently drop them, and you decide what
to do at each point. The deliberate tradeoff: **explicit and a bit tedious beats
implicit and surprising.**

</details>

<details>
<summary>3. When do you use <code>errors.New</code> vs <code>fmt.Errorf</code>?</summary>

`errors.New("...")` for a **static** message with no dynamic content. `fmt.Errorf`
when you need to **interpolate values** (`fmt.Errorf("invalid age: %d", age)`) or to
**wrap** another error with `%w`. If there's nothing to format and nothing to wrap,
`errors.New` is the simpler choice.

</details>

<details>
<summary>4. What are the conventions for error message text, and why?</summary>

**Lowercase first letter, no trailing punctuation** — e.g. `"connection failed"`, not
`"Connection failed."`. The reason is wrapping: messages get concatenated into a chain
like `"initApp: loadConfig: connection failed"`, and capital letters or periods in the
middle of that chain read wrong. (The linter flags violations via `ST1005`/`stylecheck`.)

</details>

<details>
<summary>5. What does <code>%w</code> do in <code>fmt.Errorf</code>, and how does it differ from <code>%v</code>?</summary>

`%w` **wraps** the error — it adds context text while keeping the original error in the
chain, so `errors.Is`/`errors.As` can still find it. `%v` only formats the error as a
**string**, discarding the ability to inspect it. Use `%w` when callers may need to
detect the underlying error; use `%v` when you only want the text (or deliberately want
to hide the underlying type). `%w` was added in Go 1.13.

</details>

<details>
<summary>6. What's the difference between <code>errors.Is</code> and <code>errors.As</code>?</summary>

Both walk the wrapped chain, but ask different questions. `errors.Is(err, target)` —
"does the chain contain an error **equal to** this specific value?" → for **sentinel
errors** (`errors.Is(err, sql.ErrNoRows)`). `errors.As(err, &target)` — "does the chain
contain an error of this **type**? If so, assign it to `target`" → for **custom error
types** whose fields you want to read. `Is` compares by value/identity; `As` extracts by
type into a pointer you provide.

</details>

<details>
<summary>7. Why prefer <code>errors.Is(err, ErrX)</code> over <code>err == ErrX</code>?</summary>

Because `==` only matches when `err` **is** the sentinel directly — it breaks the moment
the error is wrapped (`fmt.Errorf("...: %w", ErrX)`). `errors.Is` unwraps the whole chain
and checks each level, so it still finds the sentinel however deeply it was wrapped. Since
wrapping is the norm, always reach for `errors.Is`.

</details>

<details>
<summary>8. What is a sentinel error, and what's the naming convention?</summary>

A **predefined, exported error value** callers can check for — e.g.
`var ErrNotFound = errors.New("item not found")`. Convention names them `ErrXxx`. The
function returns the sentinel; callers detect it with `errors.Is`. Standard-library
examples: `io.EOF`, `sql.ErrNoRows`, `os.ErrNotExist`, `context.Canceled`. Use them for
simple, known conditions callers only need to **identify**.

</details>

<details>
<summary>9. When do you reach for a custom error type instead of a sentinel, and what does <code>Unwrap()</code> do?</summary>

Use a **custom error type** when the error must carry **structured data** (fields like
`StatusCode`, `Field`, `Line`) that callers extract — not just identify. Implement
`Error() string` to satisfy the interface, and implement `Unwrap() error` if your type
**wraps** another error, so `errors.Is`/`errors.As` can traverse into the wrapped one.
Callers pull it out with `errors.As`.

</details>

<details>
<summary>10. What is <code>errors.AsType</code>, and why prefer it over <code>errors.As</code> on Go 1.26+?</summary>

`errors.AsType[E error](err error) (E, bool)` (Go 1.26) is the **generic** version of
`errors.As`. Instead of declaring a target and passing its address
(`errors.As(err, &t)`), you pass the type as a parameter and get the value back,
comma-ok style: `t, ok := errors.AsType[*ParseError](err)`. It's **type-safe** — the
target type is checked at compile time, so there's no `any` argument you can get wrong
(a mismatched `errors.As` target panics at runtime) — **faster**, and easier to read.
Both walk the same wrapped tree and match the first error of that type. There's no
`IsType`, because `errors.Is` already takes a typed `error` target. You'll still meet
`errors.As` throughout existing code and modules below Go 1.26, so know both.

</details>

<details>
<summary>11. How can one error wrap several others, and how do <code>errors.Is</code>/<code>As</code> still find them?</summary>

Via the **multi-error unwrap**: alongside `Unwrap() error` (single), Go 1.20 added
`Unwrap() []error` (multiple). An error implementing the slice form exposes several
wrapped errors, and `errors.Is`/`As`/`AsType` try both unwrap shapes, so they traverse
the whole **tree**, not just a line. `errors.Join` returns such a value, and your own
custom type can too by implementing `Unwrap() []error`. Note the singular
`errors.Unwrap` function returns `nil` for a multi-error (there's no single "next"),
so inspect multi-errors with `errors.Is`/`As`/`AsType`, never a manual unwrap loop.

</details>

<details>
<summary>12. What does <code>errors.Join</code> give you, and when would you use it?</summary>

`errors.Join(errs...)` (Go 1.20+) bundles **multiple** errors into a single `error`
whose chain `errors.Is`/`errors.As` can still walk against every joined error. It skips
`nil` arguments and returns `nil` if all are nil — so you can accumulate errors in a
slice (validating every field, closing many resources) and return them at once without
special-casing "no errors." The same release also lets `fmt.Errorf` take multiple `%w`
verbs.

</details>

<details>
<summary>13. When is <code>panic</code>/<code>recover</code> appropriate — and when is it a mistake?</summary>

Appropriate only for **truly unrecoverable** situations (programmer errors, impossible
states, startup config that's malformed) and for **server resilience** — `recover` in a
deferred function so one request's panic doesn't crash a long-running server (recovery
middleware). It's a **mistake** to use it as a try/catch substitute for expected
conditions (bad input, missing files, failed network calls) — that's the #1 error JS
devs make in Go. Return errors for anything expected. `recover` only works inside a
`defer`.

</details>

<details>
<summary>14. Why is "log <em>and</em> return the error" an anti-pattern?</summary>

Because the error gets logged again at every level as it propagates, producing duplicate
log spam for a single failure. The rule: **log once, at the top**, where you actually
handle the error and decide the response. Lower levels should just **return** the error
(wrapped with `%w` context) and let the caller decide. Also don't wrap when it would leak
internal details (e.g. a DB error into an API response) — return a generic error there
instead.

</details>
