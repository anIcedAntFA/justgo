# Chapter 13 — Questions

Recall questions. Try to answer out loud before unfolding. If you can't, re-read the
matching section of [`README.md`](./README.md).

<details>
<summary>1. What are the four conventions <code>go test</code> enforces for a test function to be discovered and run?</summary>

The file must end in `_test.go`; the function name must start with `Test` followed by a
capitalized letter; it must take exactly one parameter, `t *testing.T`; and it lives in
the same package (`package foo`) or the black-box test package (`package foo_test`) in
the same folder. Miss any of these and `go test` silently ignores the function.

</details>

<details>
<summary>2. Go ships no assertion library. Why is that a deliberate choice, and how do you report a failure without one?</summary>

Tests are just Go code — no DSL, no matchers to learn, nothing to import. You write a
plain `if` check and call `t.Errorf(...)` (or `t.Error`) on failure. The philosophy
matches the language: simple and explicit over a rich framework. If you miss Jest's
`expect().toBe()` ergonomics, `testify` adds them, but the built-in style is the
starting point.

</details>

<details>
<summary>3. <code>t.Error</code> vs <code>t.Fatal</code> — what's the difference and when do you use each?</summary>

`t.Fatal`/`t.Fatalf` logs the failure and **stops the test immediately**; `t.Error`/
`t.Errorf` logs it but **keeps running**. Use `Fatal` for preconditions (a `nil` result
you'd otherwise dereference), and `Error` for independent assertions where you want to
see every failure in one run. Caveat: `Fatal` only stops the goroutine it runs on —
never call it from a goroutine you spawned inside the test.

</details>

<details>
<summary>4. Sketch a table-driven test and say why it's the idiomatic Go pattern.</summary>

Define cases as a slice of structs (`name` plus inputs and `want`), then loop and call
`t.Run(tt.name, func(t *testing.T){ ... })` for each. It wins because adding a case is
one line, each case is a named, isolated, individually-runnable subtest, failures name
the exact case, and the test logic is written once. It was born in the standard library
and is the ecosystem's shared dialect.

</details>

<details>
<summary>5. Since which Go version is <code>tt := tt</code> inside the loop unnecessary, and why did it used to be needed?</summary>

Since **Go 1.22**. Before that, the loop variable was shared across iterations, so a
subtest that ran later (e.g. with `t.Parallel()`) could capture the final value of `tt`
instead of its own. Go 1.22 made the loop variable per-iteration, removing the footgun —
so the `tt := tt` copy is now redundant, though you'll still see it in older code.

</details>

<details>
<summary>6. What does <code>t.Helper()</code> do, and what goes wrong if you forget it in an assertion helper?</summary>

`t.Helper()` marks the function as a test helper, so a failure reported inside it points
at the **caller's** line rather than the line inside the helper. Forget it and every
failure blames the same line inside the helper, making it impossible to tell which call
site actually failed.

</details>

<details>
<summary>7. When would you use <code>t.Cleanup</code>, <code>t.TempDir</code>, and <code>TestMain</code>?</summary>

`t.Cleanup(fn)` registers teardown that runs (LIFO) when the test finishes, pass or
fail — cleaner than `defer` across subtests/helpers. `t.TempDir()` returns a temp
directory that's auto-removed, no cleanup needed. `TestMain(m *testing.M)` wraps the
whole package: do one-time setup, call `m.Run()`, tear down, then `os.Exit(code)` — for
things like spinning up a shared test database.

</details>

<details>
<summary>8. White-box (<code>package foo</code>) vs black-box (<code>package foo_test</code>) testing — what's the trade-off, and which is the better default?</summary>

White-box tests share the package and can touch unexported identifiers — handy for
testing internal helpers directly, but they couple to internals. Black-box tests live in
`package foo_test` (same folder) and can only use the **exported** API, exactly as a
real consumer would, so they test behaviour through the public contract. Black-box is
the better default; drop to white-box only for the few internals worth testing directly.

</details>

<details>
<summary>9. Name the three main test tiers of the pyramid and say which a backend/CLI dev should prioritise.</summary>

**Unit** (one function/type in isolation — fast, deterministic, the majority),
**integration** (several real components together — code + real DB/HTTP, slower), and
**end-to-end** (drive the whole system like a user — highest confidence, most brittle).
Prioritise unit tests of your core logic, add a handful of integration tests on the
seams that actually break, and keep a thin layer of E2E on critical paths. Don't invert
the pyramid.

</details>

<details>
<summary>10. Coverage is reported as a percentage. Why is chasing 100% a bad goal?</summary>

Coverage only measures which lines _executed_ during tests, not whether the tests
_assert_ anything meaningful. A line run by a test with no real assertions counts as
"covered" and proves nothing. Coverage is a smoke detector for what's untested, not a
quality score — chase coverage of your branches and edge cases, not a round number.

</details>

<details>
<summary>11. What replaced <code>for i := 0; i &lt; b.N; i++</code> in benchmarks, and what problem does it solve?</summary>

`for b.Loop() { ... }` (Go 1.24+). Besides being cleaner, it prevents the compiler from
optimising away the benchmarked call — a classic footgun where a benchmark measured
nothing because the result was unused and eliminated. Run with `go test -bench=.`, add
`-benchmem` for allocation stats.

</details>

<details>
<summary>12. What is fuzzing good for, and what's the key idea that lets a fuzzer find bugs you never wrote a case for?</summary>

Fuzzing (`go test -fuzz`, Go 1.18+) feeds a function generated random inputs to find
crashes and edge cases. The key idea is asserting a **property** that must hold for
_any_ input — e.g. `Reverse(Reverse(s)) == s` — instead of hard-coding expected outputs.
The fuzzer then searches for a counterexample. It shines on code that parses untrusted
input: decoders, validators, parsers.

</details>

<details>
<summary>13. Why is <code>go test -race</code> considered the single highest-value flag, and what are its limits?</summary>

It instruments the program to detect data races — unsynchronized concurrent access to
shared memory, which are otherwise nondeterministic and brutal to debug. CI should run
it always (this repo's `just test` does). Its limit: it only reports races it actually
_observes at runtime_, so it needs tests that exercise the concurrent paths — it can't
find a race in code the tests never run.

</details>

<details>
<summary>14. Built-in checks vs <code>testify</code> — what does testify add, and what's the standard advice?</summary>

testify's `assert`/`require` give concise, Jest-like assertions with nice failure diffs;
`require` stops the test (like `t.Fatal`), `assert` continues (like `t.Error`). It's the
de-facto third-party standard, but the Go project and Google's style stay built-in.
Advice: master the built-in `if`/`t.Errorf` style first, add testify only when the
boilerplate wears on you, and keep a single repo consistent either way.

</details>

<details>
<summary>15. What's the difference between a fake and a mock, and which should you prefer?</summary>

A **fake** is a working lightweight implementation (e.g. an in-memory store); a **mock**
additionally records and asserts on _how_ it was called ("`GetUser` was called once with
id 1"). Prefer fakes and assert on outcomes — they're less brittle. Reach for mocks
(hand-written or generated with `mockery`/`mockgen`) only when the interaction itself is
the behaviour under test.

</details>

<details>
<summary>16. Give three properties that make a test worth writing (rather than noise).</summary>

Any three of: it tests **behaviour** (observable output for an input) not implementation
details; it would actually **fail if the code were wrong** (real assertions — verify by
breaking the code and watching it go red); it covers the **edges** (empty, zero,
negative, nil, max length, malformed) not just the happy path; it's **deterministic and
independent** (no wall-clock, random seed, map order, or ordering dependence); and it
**reads clearly on failure** (a message naming got vs want).

</details>
