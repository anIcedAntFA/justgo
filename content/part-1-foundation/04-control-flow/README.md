# Chapter 04: Control Flow

> **if/else, for (the only loop), switch (much more powerful than JS), and defer.**

## TL;DR

Go has fewer control flow constructs than JS — no `while`, no `do...while`, no `forEach`, no `for...of`, no ternary `? :`. Just `if`, `for`, `switch`, and `defer`. But each one is more capable than its JS equivalent. The simplicity is the point: fewer ways to do things means less to remember and more consistent code.

---

## if / else

### Basic Form

```go
age := 20

if age >= 18 {
    fmt.Println("adult")
} else if age >= 13 {
    fmt.Println("teenager")
} else {
    fmt.Println("child")
}
```

**No parentheses** around the condition. Braces `{}` are always required — even for one-liners.

```go
// ❌ Compile error — no braces
if age >= 18
    fmt.Println("adult")

// ⚠️ Compiles, but gofmt strips the parens — don't write them
if (age >= 18) {
    fmt.Println("adult")
}

// ✅ The Go way
if age >= 18 {
    fmt.Println("adult")
}
```

Coming from JS where you might write `if (x) return true` on one line — Go always requires braces. This eliminates ambiguity and makes `if` blocks consistent everywhere.

### if with Init Statement

This is Go's most distinctive `if` feature. You can declare a variable scoped to the `if/else` block:

```go
if err := doSomething(); err != nil {
    fmt.Println("error:", err)
    return
}
// err does NOT exist here — scoped to the if block
```

This pattern is **everywhere** in Go. It keeps the error variable scoped tightly so it doesn't leak into the rest of the function.

```go
// Without init statement — err leaks
err := doSomething()
if err != nil {
    return err
}
// err still exists here, potentially shadowing later errs

// With init statement — err is contained
if err := doSomething(); err != nil {
    return err
}
// clean scope here
```

JS has no equivalent. The closest is:

```javascript
// JavaScript — no init-statement in if
const err = doSomething()
if (err) {
    // handle
}
// err still in scope
```

### No Ternary Operator

```javascript
// JavaScript
const status = age >= 18 ? "adult" : "minor"
```

```go
// Go — no ternary. Use if/else.
var status string
if age >= 18 {
    status = "adult"
} else {
    status = "minor"
}
```

Yes, it's more lines. The Go team deliberately excluded the ternary operator because nested ternaries become unreadable fast. The extra lines buy you clarity.

---

## for — The Only Loop

Go has exactly one loop keyword: `for`. It covers every loop pattern.

### Classic for (like C/JS for)

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

Note: `i` is scoped to the loop — after the loop, `i` doesn't exist. Go never had JS's `var` scope-leak, where the counter escaped the loop body.

### while-style (condition only)

```go
// Go's "while loop" — just for with a condition
n := 1
for n < 100 {
    n *= 2
}
fmt.Println(n)   // 128
```

No `while` keyword. `for condition { }` is `while`.

### Infinite loop

```go
for {
    // runs forever until break or return
    input := readInput()
    if input == "quit" {
        break
    }
    process(input)
}
```

No `while (true)`. Just `for { }`. Clean.

### range — Iterating Over Collections

`range` is Go's answer to `for...of`, `forEach`, `.map()`, and `Object.entries()`:

```go
// Slice (array)
names := []string{"Alice", "Bob", "Charlie"}
for i, name := range names {
    fmt.Printf("%d: %s\n", i, name)
}

// Skip index with _
for _, name := range names {
    fmt.Println(name)
}

// Only index
for i := range names {
    fmt.Println(i)
}

// String — iterates over runes, not bytes
for i, ch := range "Hello, 世界" {
    fmt.Printf("byte %d: %c\n", i, ch)
}

// Map
ages := map[string]int{"Alice": 30, "Bob": 25}
for key, value := range ages {
    fmt.Printf("%s is %d\n", key, value)
}

// Channel (we'll cover this in Part 3)
for msg := range messageChannel {
    process(msg)
}
```

### range over integers (Go 1.22+)

A recent addition — iterate over a number:

```go
// Go 1.22+: range over integer
for i := range 5 {
    fmt.Println(i)   // 0, 1, 2, 3, 4
}

// Equivalent to:
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

Simple but convenient. No need for a classic 3-part `for` when you just want to run something N times. The official release notes summarise it as: _""For" loops may now range over integers."_

### Loop variables are per-iteration (Go 1.22+)

Each iteration of a `for` loop gets its **own** copy of the loop variable. This matters the moment you capture that variable in a closure (and, in Part 3, a goroutine):

```go
funcs := []func(){}
for i := range 3 {
    funcs = append(funcs, func() { fmt.Print(i, " ") })
}
for _, f := range funcs {
    f()
}
// Go 1.22+ : 0 1 2   — each closure captured its own i
// Go ≤ 1.21: 3 3 3   — every closure shared one i
```

Before Go 1.22 the loop variable was created **once** and reused across iterations, so every captured closure saw the final value — a notorious footgun. The release notes state it directly: _"In Go 1.22, each iteration of the loop creates new variables, to avoid accidental sharing bugs."_ This repo targets Go 1.26, so you get the safe behaviour by default — but you'll still meet the old bug in pre-1.22 code and blog posts, so recognise it.

### JS Loop Comparison

```javascript
// JavaScript has 5+ loop styles:
for (let i = 0; i < arr.length; i++) { }   // classic for
for (const item of arr) { }                 // for...of (iterables)
for (const key in obj) { }                  // for...in (object keys — avoid on arrays)
arr.forEach((item, i) => { })               // array method
while (condition) { }                       // while
do { } while (condition)                    // do...while
```

```go
// Go has 1 loop keyword covering all cases:
for i := 0; i < len(arr); i++ { }   // classic
for i, v := range arr { }           // range (most common)
for condition { }                    // while-style
for { }                              // infinite
for i := range 5 { }                // N times
```

One keyword, multiple forms. Every Go developer reads the same patterns.

### break, continue, labels

```go
// break — exit the loop
for i := 0; i < 100; i++ {
    if i == 5 {
        break   // stops the loop entirely
    }
}

// continue — skip to next iteration
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue   // skip even numbers
    }
    fmt.Println(i)  // 1, 3, 5, 7, 9
}
```

Labels let you `break` or `continue` an **outer** loop from inside a nested one:

```go
// break outer — leave BOTH loops
outer:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 1 && j == 1 {
                break outer   // breaks out of BOTH loops
            }
            fmt.Printf("(%d,%d) ", i, j)
        }
    }
// (0,0) (0,1) (0,2) (1,0)

// continue outer — skip to the next iteration of the OUTER loop
rows:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 1 {
                continue rows   // abandon this row, go to next i
            }
            fmt.Printf("(%d,%d) ", i, j)
        }
    }
// (0,0) (1,0) (2,0)
```

In JS, labeled `break`/`continue` exist too but are rarely used. In Go, they're the idiomatic way to control nested loops — no need for flag variables.

---

## switch

Go's `switch` is significantly more powerful than JavaScript's.

### Basic switch

```go
day := "Tuesday"

switch day {
case "Monday":
    fmt.Println("Start of week")
case "Tuesday", "Wednesday", "Thursday":    // multiple values in one case
    fmt.Println("Midweek")
case "Friday":
    fmt.Println("Almost weekend")
default:
    fmt.Println("Weekend")
}
```

**Key difference from JS:** Go's switch does NOT fall through by default. No `break` needed.

```javascript
// JavaScript: falls through without break (a constant source of bugs)
switch (day) {
    case "Monday":
        console.log("Start")
        // OOPS — forgot break, falls through to Tuesday!
    case "Tuesday":
        console.log("Tuesday")
        break
}
```

```go
// Go: each case breaks automatically
switch day {
case "Monday":
    fmt.Println("Start")
    // stops here — no fallthrough
case "Tuesday":
    fmt.Println("Tuesday")
}
```

If you actually want fallthrough (rare), you must explicitly say so:

```go
switch day {
case "Monday":
    fmt.Println("Start")
    fallthrough            // explicitly falls through to next case
case "Tuesday":
    fmt.Println("Tuesday")
}
// prints both "Start" and "Tuesday" if day is "Monday"
```

### switch with init statement

Like `if`, switch supports an init statement:

```go
switch os := runtime.GOOS; os {
case "linux":
    fmt.Println("Linux")
case "darwin":
    fmt.Println("macOS")
default:
    fmt.Printf("Other: %s\n", os)
}
// os is not accessible here
```

### switch without a condition (tagless switch)

This is Go's clean replacement for long `if/else if/else if` chains:

```go
age := 25

switch {
case age < 13:
    fmt.Println("child")
case age < 18:
    fmt.Println("teenager")
case age < 65:
    fmt.Println("adult")
default:
    fmt.Println("senior")
}
```

Each `case` is a boolean expression evaluated top-down. First match wins. This is much cleaner than:

```go
// Don't do this — messy
if age < 13 {
    fmt.Println("child")
} else if age < 18 {
    fmt.Println("teenager")
} else if age < 65 {
    fmt.Println("adult")
} else {
    fmt.Println("senior")
}
```

### Type switch (preview — more in Chapter 07)

Go can switch on the type of a value:

```go
func describe(i any) {
    switch v := i.(type) {
    case int:
        fmt.Printf("integer: %d\n", v)
    case string:
        fmt.Printf("string: %q\n", v)
    case bool:
        fmt.Printf("boolean: %t\n", v)
    default:
        fmt.Printf("unknown type: %T\n", v)
    }
}

describe(42)       // integer: 42
describe("hello")  // string: "hello"
describe(true)     // boolean: true
```

(`any` is the idiomatic alias for `interface{}` since Go 1.18 — same type, nicer name.) JS has `typeof`, but it's a runtime check that returns a string. Go's type switch gives you a **typed** variable `v` in each case, checked by the compiler.

---

## defer

`defer` is unique to Go — it schedules a function call to run when the enclosing function returns. Think of it as a guaranteed cleanup mechanism.

### Basic Usage

```go
func readFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()     // will run when readFile returns, no matter what

    // ... read file, process content ...
    // even if something panics, file.Close() still runs
    return nil
}
```

**Why this matters:** In JS, you'd use `try/finally` for cleanup:

```javascript
// JavaScript
let file
try {
    file = await fs.open(path)
    // ... process ...
} finally {
    if (file) await file.close()
}
```

```go
// Go — cleaner, the cleanup is right next to the open
file, err := os.Open(path)
if err != nil {
    return err
}
defer file.Close()
// ... process (no need for try/finally wrapper) ...
```

The `defer` sits right next to the resource acquisition. Open → defer Close. Lock → defer Unlock. Start → defer Stop. The reader immediately sees the lifecycle.

### defer is LIFO (Last In, First Out)

Multiple defers execute in reverse order:

```go
func main() {
    fmt.Println("start")
    defer fmt.Println("first defer")
    defer fmt.Println("second defer")
    defer fmt.Println("third defer")
    fmt.Println("end")
}
// Output:
// start
// end
// third defer
// second defer
// first defer
```

Like a stack: last deferred, first executed.

### defer Arguments Are Evaluated Immediately

A common gotcha:

```go
x := 10
defer fmt.Println(x)     // captures x=10 NOW
x = 20
fmt.Println(x)

// Output:
// 20
// 10      ← defer captured the value at defer time, not at execution time
```

If you need the final value, use a closure:

```go
x := 10
defer func() {
    fmt.Println(x)         // reads x when the deferred func runs
}()
x = 20

// Output:
// 20      ← the deferred closure sees x=20
```

### Common defer Patterns

```go
// File cleanup
f, err := os.Open("data.txt")
if err != nil { return err }
defer f.Close()

// Mutex unlock
mu.Lock()
defer mu.Unlock()

// Timing a function
start := time.Now()
defer func() {
    fmt.Printf("took %v\n", time.Since(start))
}()

// Recovering from panics (like try/catch for panics)
defer func() {
    if r := recover(); r != nil {
        fmt.Println("recovered from panic:", r)
    }
}()
```

### defer in Loops — Caution

```go
// ⚠️ Bad — defers pile up, files stay open until function returns
for _, path := range files {
    f, err := os.Open(path)
    if err != nil { continue }
    defer f.Close()            // won't close until the entire function returns!
    process(f)
}

// ✅ Good — extract to a function so defer runs per iteration
for _, path := range files {
    if err := processFile(path); err != nil {
        log.Println(err)
    }
}

func processFile(path string) error {
    f, err := os.Open(path)
    if err != nil { return err }
    defer f.Close()            // closes when processFile returns (each iteration)
    return process(f)
}
```

This is a real-world gotcha. `defer` runs when the **enclosing function** returns, not when the block ends. In a loop, that means all defers pile up until the loop's function returns. The fix: extract the loop body into its own function.

---

## goto (Exists, Don't Use)

Go has `goto`, but you'll almost never write it — it exists for generated code and a few low-level patterns. Mentioning it only so you recognise it:

```go
    goto end
    fmt.Println("skipped")
end:
    fmt.Println("done")
```

Reach for a loop, `break`, `continue`, or an extracted function instead.

---

## Runnable Examples

The [`examples/`](./examples/) folder has three programs that combine this chapter's constructs. Run each with `go run .` from its directory:

- **[`examples/grades`](./examples/grades/)** — the capstone: `defer` timing, `range` over an integer, `range` over a slice, a tagless `switch`, and an `if` with init statement, all in one small program.
- **[`examples/countdown`](./examples/countdown/)** — a `for` countdown with `time.Sleep` and a `defer`-ed launch message.
- **[`examples/repl`](./examples/repl/)** — an infinite `for` loop reading stdin, dispatching commands with a `switch`.

---

## Common Mistakes

### Mistake 1: Adding parentheses to if/for conditions

```go
if (x > 0) {                  // compiles, but gofmt removes the parens
for (i := 0; i < 10; i++) {   // ❌ syntax error — parens not allowed in for
```

`gofmt` strips unnecessary parentheses from `if`. For `for`, parentheses are a syntax error.

### Mistake 2: Forgetting that switch doesn't fall through

If you want multiple cases to share code:

```go
// ❌ Wrong — trying to fall through
switch day {
case "Saturday":
    fallthrough    // this just falls to Sunday, doesn't share logic well
case "Sunday":
    fmt.Println("Weekend")
}

// ✅ Right — multiple values in one case
switch day {
case "Saturday", "Sunday":
    fmt.Println("Weekend")
}
```

### Mistake 3: defer in loops

As covered above — extract the loop body into a function when deferring inside loops.

### Mistake 4: Modifying a slice while ranging over it

```go
// ⚠️ range evaluates the slice's length once, at the start
nums := []int{1, 2, 3}
for _, n := range nums {
    if n == 2 {
        nums = append(nums, 4)   // don't do this
    }
}
```

Don't modify a slice while iterating over it with `range`. Use a classic `for` loop with an index if you need to mutate during iteration.

---

## Exercises

Two graded exercises live in [`exercises/`](./exercises/) — pure functions with table-driven tests. Remove the `t.Skip` at the top of each test, implement the function, and run `go test ./...` until it passes.

### Exercise 1: FizzBuzz (`FizzBuzz(n int) string`)

Return the FizzBuzz string for a single number `n` — `"Fizz"` when divisible by 3, `"Buzz"` when divisible by 5, `"FizzBuzz"` when divisible by both, otherwise the number as a string. Implement it with a **tagless `switch`**, no `if/else`. (A driver loop over 1–100 that prints each result is a nice extra, but the graded part is the pure function.)

### Exercise 2: Number Classifier (`Classify(n int) string` + `IsPrime(n int) bool`)

Classify an integer with a **tagless `switch`**, evaluated top-down so the first match wins:

1. divisible by **both** 3 and 5 → `"fizzbuzz"`
2. prime → `"prime"`
3. even → `"even"`
4. otherwise → `"odd"`

Write the `IsPrime(n int) bool` helper with a `for` loop. Precedence examples: `15 → "fizzbuzz"`, `2 → "prime"` (prime beats even), `8 → "even"`, `9 → "odd"`.

---

## Key Takeaways

1. **One loop to rule them all.** `for` covers classic, while, infinite, and range iteration. No `while`, no `do-while`, no `forEach`.

2. **Loop variables are per-iteration (Go 1.22+).** Each iteration gets a fresh copy, so closures capture the value you expect. Older code doesn't — recognise the pre-1.22 sharing bug.

3. **switch doesn't fall through.** No forgotten `break` bugs. Use comma-separated values for multiple matches. Tagless switch replaces long if/else chains.

4. **`if` with init statement** is a Go signature pattern. Use it to scope error variables tightly: `if err := doX(); err != nil { }`.

5. **defer = guaranteed cleanup.** Put it right after resource acquisition. Remember: LIFO order, arguments evaluated immediately, runs when the function returns (not the block).

6. **No ternary operator.** Write the `if/else`. The extra two lines buy readability that scales.

---

## 🧭 Navigation

| Direction    | Link                                                                   |
| ------------ | ---------------------------------------------------------------------- |
| **Previous** | [← Chapter 03: Types & Variables](../03-types-and-variables/README.md) |
| **Next**     | [Chapter 05: Functions →](../05-functions/README.md)                   |
