# Chapter 04 — Questions

Recall questions. Try to answer out loud before unfolding. If you can't, re-read the
matching section of [`README.md`](./README.md).

<details>
<summary>1. Go has one loop keyword. Name the four <em>forms</em> of <code>for</code> and the JS construct each one replaces.</summary>

`for` is the only loop keyword. Its forms: **classic** `for i := 0; i < n; i++` (JS
classic `for`); **while-style** `for cond {}` (JS `while`); **infinite** `for {}`
(JS `while (true)`); and **range** `for i, v := range coll` (JS `for...of` /
`forEach` / `Object.entries`). Since Go 1.22, `for i := range n` also loops N times.

</details>

<details>
<summary>2. What changed about <code>for</code> loop variables in Go 1.22, and why does it matter for closures?</summary>

Before Go 1.22 the loop variable was created **once** and reused every iteration, so
closures (and goroutines) capturing it all saw the **final** value — the classic
`3 3 3` bug. Go 1.22 makes **each iteration create a new variable**, so each closure
captures its own copy (`0 1 2`). This repo is on Go 1.26, so you get the safe
behaviour by default.

</details>

<details>
<summary>3. How does Go's <code>switch</code> differ from JavaScript's on fallthrough, and how do you get each behaviour?</summary>

Go's `switch` does **not** fall through — each case breaks automatically, no `break`
keyword needed. To match several values, list them comma-separated in one case
(`case "Sat", "Sun":`). To deliberately fall into the next case, use the explicit
`fallthrough` keyword (rare). JS is the opposite: it falls through unless you write
`break`, a frequent source of bugs.

</details>

<details>
<summary>4. What is a <em>tagless</em> switch, and what does it replace?</summary>

A `switch` with no value after the keyword — `switch { case cond1: ...; case cond2:
... }`. Each case is a boolean expression evaluated top-down; the first true case
wins. It's the idiomatic replacement for a long `if / else if / else if` chain.

</details>

<details>
<summary>5. When does a <code>defer</code>-ed call run, in what order, and when are its arguments evaluated?</summary>

A deferred call runs when the **enclosing function returns** (not when the block
ends), including on `panic`. Multiple defers run **LIFO** — last deferred, first
executed. Its **arguments are evaluated immediately** at the `defer` statement, so
`defer fmt.Println(x)` captures `x`'s value then; wrap it in a closure
(`defer func(){ ... }()`) if you want the value at execution time.

</details>

<details>
<summary>6. Why is <code>defer f.Close()</code> inside a <code>for</code> loop a trap, and what's the fix?</summary>

`defer` fires when the **function** returns, not each iteration — so the deferred
closes pile up and resources stay open until the whole loop-containing function
exits. Fix: extract the loop body into its own function so each call's `defer` runs
per iteration.

</details>

<details>
<summary>7. Go has no ternary operator. What do you write instead, and why did the Go team leave it out?</summary>

Write a plain `if/else` assigning to a pre-declared variable. The Go team omitted
`?:` deliberately: nested ternaries become unreadable fast, and the extra couple of
lines buy clarity that scales. (There is also no `while`, `do...while`, or `goto` in
day-to-day use.)

</details>
