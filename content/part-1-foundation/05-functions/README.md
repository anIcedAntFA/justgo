# Chapter 05: Functions

> **Multiple returns, named returns, variadic functions, first-class functions, and closures.**

## TL;DR

Functions in Go are first-class values like in JavaScript, but with three differences that shape everything else: they can return **multiple values** (the foundation of Go's error handling), every parameter is **typed**, and there are **no default parameters and no overloading**. Multiple returns are what make `value, err :=` possible everywhere — the single most common line in real Go code.

---

## Basic Function Syntax

```go
func add(a int, b int) int {
    return a + b
}

// When consecutive params share a type, omit the repeats:
func add(a, b int) int {
    return a + b
}

// No return value
func greet(name string) {
    fmt.Printf("Hello, %s\n", name)
}

result := add(3, 5) // 8
greet("Gopher")     // Hello, Gopher
```

The type comes **after** the name, and the return type comes after the parameter list — the opposite order from TypeScript:

```typescript
// TypeScript: colons, return type after the params
function add(a: number, b: number): number {
    return a + b
}
```

```go
// Go: type after name, no colons, return type after the parens
func add(a, b int) int {
    return a + b
}
```

Once you adjust to "name first, type second," it reads naturally. Go chose this order because it scales better for complex declarations — especially function types and pointers, which you'll meet below.

---

## Multiple Return Values

This is Go's defining feature for functions: one function, several return values.

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 2)
if err != nil {
    fmt.Println("error:", err)
    return
}
fmt.Println("result:", result) // 5
```

**This `(value, error)` shape is the backbone of Go.** Almost every function that can fail returns it, and you handle the error immediately — no exceptions, no `try/catch`, just values you check. JavaScript can only _fake_ this by returning an array or object and destructuring; Go has it in the language, type-checked on every return value.

```javascript
// JavaScript: return a tuple-shaped array and destructure
function divide(a, b) {
    if (b === 0) return [null, new Error("division by zero")]
    return [a / b, null]
}
const [result, err] = divide(10, 2)
```

### The (value, ok) Pattern

Beyond errors, Go uses a second return value to answer "does this exist?":

```go
ages := map[string]int{"Alice": 30}

age, ok := ages["Alice"]
fmt.Println(age, ok) // 30 true

age, ok = ages["Bob"]
fmt.Println(age, ok) // 0 false — 0 is the zero value; ok tells you it was missing

// Idiomatic: check existence inline with an if-init statement
if age, ok := ages["Alice"]; ok {
    fmt.Println("Alice is", age)
}
```

In JS you'd write `if ("Alice" in ages)` or test `=== undefined` — which breaks if the value legitimately _is_ `undefined`. Go's `(value, ok)` is unambiguous: `ok` definitively tells you whether the key existed. You'll see the same shape with type assertions (`v, ok := x.(int)`) and channel receives (`v, ok := <-ch`) in later chapters.

---

## Named Return Values

Go lets you name the return values. They become variables, pre-initialised to their zero values.

```go
func divide(a, b float64) (result float64, err error) {
    if b == 0 {
        err = errors.New("division by zero")
        return // "naked return" — returns the current result and err
    }
    result = a / b
    return // returns result, nil
}
```

A `return` with no arguments (a **naked return**) returns whatever the named values currently hold.

**Good use — documentation.** The names tell the caller what each value means:

```go
func parseConfig(path string) (config *Config, err error) {
    // ...
}
```

**Good use — modifying the return inside a `defer`.** This is the one place named returns are genuinely necessary — a deferred closure can inspect and rewrite the result (e.g. turning a recovered panic into an error):

```go
func process() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered: %v", r) // rewrites the named return
        }
    }()
    // ... code that might panic ...
    return nil
}
```

**Avoid — naked returns in long functions.** Fifty lines down, `return` on its own forces the reader to scroll up and reconstruct what `a`, `b`, `c` hold. Google's Go Style Guide recommends naming returns for clarity but writing **explicit** returns anyway; reserve naked returns for very short functions.

```go
// ⚠️ Hard to read — what is returned 50 lines down?
func complexThing() (a int, b string, c error) {
    // ... 50 lines ...
    return // have to scroll up to know
}
```

---

## Variadic Functions

Functions that take a variable number of arguments — Go's `...rest`.

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

sum(1, 2, 3)       // 6
sum(1, 2, 3, 4, 5) // 15
sum()              // 0 — nums is an empty (nil) slice

// Spread a slice into the variadic parameter with ...
numbers := []int{1, 2, 3, 4}
sum(numbers...) // 10
```

Inside the function `nums` is a `[]int`. The concept mirrors JS almost exactly — `...int` is the variadic parameter (like `...nums`), and `slice...` is the spread (like `...arr`). The most familiar variadic function you already use is `fmt.Println(a, b, c, ...)`.

**Rule: the variadic parameter must be last.**

```go
func logf(level, format string, args ...any) {
    // level and format are fixed; args is variadic — must come last
}

// ❌ Compile error — variadic parameter is not last
// func bad(args ...int, name string) {}
```

> **Modern Go note.** That `args ...any` used to be written `args ...interface{}`. Since Go 1.18, `any` is the built-in alias for `interface{}` — same type, nicer name. Since **Go 1.26**, `go fix` bundles a modernizer that rewrites `interface{}` to `any` (among other updates) automatically, so you rarely type the old spelling anymore.

---

## Functions as First-Class Values

Functions are values. Assign them to variables, pass them as arguments, return them from other functions — just like JavaScript, but with explicit type signatures.

```go
// Assign a function to a variable
add := func(a, b int) int { return a + b }
fmt.Println(add(3, 4)) // 7

// A function-typed variable
var operation func(int, int) int
operation = add
fmt.Println(operation(5, 6)) // 11
```

### Higher-Order Functions

```go
// Take a function as a parameter
func applyTwice(f func(int) int, x int) int {
    return f(f(x))
}

double := func(n int) int { return n * 2 }
fmt.Println(applyTwice(double, 3)) // 12  (3 → 6 → 12)

// Return a function
func multiplier(factor int) func(int) int {
    return func(n int) int { return n * factor }
}

triple := multiplier(3)
fmt.Println(triple(5)) // 15
```

This is the same shape as JavaScript, with the type spelled out:

```javascript
const multiplier = (factor) => (n) => n * factor
multiplier(3)(5) // 15
```

```go
func multiplier(factor int) func(int) int {
    return func(n int) int { return n * factor }
}
multiplier(3)(5) // 15
```

The type `func(int) int` reads as "a function taking an `int` and returning an `int`." When a signature gets noisy, give it a name with `type`:

```go
type IntOp func(int) int // now `func applyTwice(f IntOp, x int) int` reads cleanly
```

### Functions as Iterators (Go 1.23+)

There's a modern payoff to functions-are-values worth meeting now, even though the deep dive is [Chapter 10](../10-collections/README.md): since **Go 1.23**, a `for range` loop can range over a _function_ — one whose single parameter is a `yield` callback. That iterator function _is_ a first-class value of type `iter.Seq[V]`:

```go
// iter.Seq[V] is just: func(yield func(V) bool)
func countTo(n int) iter.Seq[int] {
    return func(yield func(int) bool) {
        for i := 1; i <= n; i++ {
            if !yield(i) { // yield returns false when the loop breaks early
                return
            }
        }
    }
}

for v := range countTo(3) {
    fmt.Println(v) // 1, 2, 3
}
```

The standard library returns these iterators from `maps` and `slices`, which is why you'll see idioms like ranging a map's keys in sorted order:

```go
for _, k := range slices.Sorted(maps.Keys(m)) {
    fmt.Println(k)
}
```

The takeaway for _this_ chapter: an iterator is not new machinery — it's an ordinary function with a specific signature, made possible by first-class functions and closures. Chapter 10 covers writing, composing, and pulling from them (`iter.Pull`) in full.

---

## Closures

A closure is a function that references variables from outside its own body. Go closures capture variables **by reference**, exactly like JS closures.

```go
func counter() func() int {
    count := 0
    return func() int {
        count++ // captures and mutates count from the enclosing scope
        return count
    }
}

c := counter()
fmt.Println(c(), c(), c()) // 1 2 3

c2 := counter()
fmt.Println(c2()) // 1 — independent; each counter() call gets its own count
```

Each call to `counter()` creates a fresh `count`, closed over by the returned function and kept alive for as long as that function is reachable.

### Practical Closure: Memoization

```go
func memoizedFib() func(int) int {
    cache := map[int]int{}

    var fib func(int) int
    fib = func(n int) int {
        if n < 2 {
            return n
        }
        if v, ok := cache[n]; ok {
            return v // cached
        }
        result := fib(n-1) + fib(n-2)
        cache[n] = result
        return result
    }
    return fib
}

fib := memoizedFib()
fmt.Println(fib(40)) // fast — the closed-over cache persists across calls
```

### The Loop-Variable Gotcha (fixed in Go 1.22)

The most famous closure trap in both languages. Before Go 1.22 the `for` loop variable was created **once** and shared across iterations, so every closure captured the _same_ variable and saw its final value:

```go
funcs := []func(){}
for i := 0; i < 3; i++ {
    funcs = append(funcs, func() { fmt.Print(i, " ") })
}
for _, f := range funcs {
    f()
}
// Go ≤ 1.21: 3 3 3   (all closures shared one i)
// Go 1.22+ : 0 1 2   (each iteration gets its own i)
```

This was the exact bug JavaScript had with `var` in loops, fixed by `let`. Go 1.22 fixed it the same way — a fresh loop variable per iteration — so on this repo's Go 1.26 you get `0 1 2` by default. You'll still meet the old behaviour in pre-1.22 code and old Stack Overflow answers, so recognise it.

**But the underlying rule still bites.** Closures capture by reference, so sharing _any_ mutable variable across closures — not just a loop variable — means they all see its latest value. The 1.22 change only fixed the loop-variable case; be deliberate whenever several closures close over the same variable.

---

## No Default Parameters, No Overloading

Two conveniences from other languages that Go deliberately omits.

### No Default Parameters

```go
// ❌ Not valid Go:
// func greet(name string = "World") {}
```

The idiomatic replacements:

**Option 1 — separate functions** for the simple case:

```go
func Greet(name string) { fmt.Printf("Hello, %s\n", name) }
func GreetWorld()       { Greet("World") }
```

**Option 2 — the functional options pattern** for configurable structs with defaults. This is _the_ idiomatic Go answer, and you'll see it across the ecosystem (HTTP servers, database drivers, gRPC):

```go
type ServerConfig struct {
    Port    int
    Timeout time.Duration
    MaxConn int
}

type Option func(*ServerConfig)

func WithPort(port int) Option {
    return func(c *ServerConfig) { c.Port = port }
}

func WithTimeout(t time.Duration) Option {
    return func(c *ServerConfig) { c.Timeout = t }
}

func NewServer(opts ...Option) *ServerConfig {
    config := &ServerConfig{ // defaults
        Port:    8080,
        Timeout: 30 * time.Second,
        MaxConn: 100,
    }
    for _, opt := range opts { // apply overrides
        opt(config)
    }
    return config
}

// Override only what you care about:
server := NewServer(WithPort(9000), WithTimeout(60*time.Second))
// Port=9000, Timeout=60s, MaxConn=100 (default)
```

Notice it ties the whole chapter together: variadic `opts ...Option`, first-class function values, and closures (each `With…` returns a closure over its argument). More verbose than a default parameter, but self-documenting and extensible without breaking callers.

### No Function Overloading

```go
// ❌ Can't declare the same name twice:
// func process(s string) {}
// func process(n int)    {} // compile error: process redeclared
```

JS fakes overloading with `typeof` checks; TypeScript has overload signatures. Go has neither — one name, one function. Reach for distinct names, an interface ([Chapter 07](../07-interfaces/README.md)), or generics ([Chapter 12](../12-generics/README.md)):

```go
func ProcessString(s string) {}
func ProcessInt(n int)       {}

func Process[T any](v T) {} // generics, when the logic is truly type-independent
```

---

## Recursion

Standard recursion, no surprises:

```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

fmt.Println(factorial(5)) // 120
```

Go does **not** perform tail-call optimisation, so deeply recursive code can overflow the stack — prefer iteration for unbounded depth. In practice goroutine stacks start small and grow dynamically, so you have more headroom than most languages, but the guarantee isn't there.

---

## init Functions

A special function that runs automatically when a package loads, before `main`:

```go
package main

import "fmt"

var config string

func init() {
    config = "loaded" // runs automatically, before main()
    fmt.Println("init: config initialized")
}

func main() {
    fmt.Println("main:", config)
}

// init: config initialized
// main: loaded
```

You can have multiple `init()` functions (even in one file); they run in declaration order, after all package-level variables are initialised, and you can't call them yourself. Use `init` sparingly — for setup that genuinely must happen first (registering a database driver, validating embedded config). Overusing it makes control flow "magical" and hard to trace.

---

## Runnable Examples

The [`examples/`](./examples/) folder has three programs. Run each with `go run .` from its directory:

- **[`examples/pipeline`](./examples/pipeline/)** — the capstone: first-class functions passed around, a variadic transformer, a closure, and a range-over-func `iter.Seq` iterator (Go 1.23+) in one small program.
- **[`examples/closures`](./examples/closures/)** — a counter and a memoized Fibonacci, showing state captured and kept alive by closures.
- **[`examples/options`](./examples/options/)** — the functional options pattern building a configured value from defaults + overrides.

---

## Common Mistakes

### Mistake 1: Ignoring a returned error

```go
// ❌ Silently dropping the error — you won't know it failed
result, _ := divide(10, 0)

// ✅ Handle it
result, err := divide(10, 0)
if err != nil {
    // ...
}
```

`golangci-lint`'s `errcheck` flags dropped errors. Don't `_` an error unless you truly don't care.

### Mistake 2: Expecting default parameters

```go
// ❌ Not valid Go
// func connect(host string, port int = 8080) {}
```

Use the functional options pattern or separate functions.

### Mistake 3: Naked returns in long functions

```go
// ⚠️ Avoid — what is returned here?
func longFunc() (result int, err error) {
    // 40 lines...
    return
}

// ✅ Explicit returns read clearly regardless of length
func longFunc() (int, error) {
    // 40 lines...
    return result, nil
}
```

### Mistake 4: A variadic parameter that isn't last

```go
// ❌ Compile error
// func bad(args ...int, name string) {}

// ✅ Variadic must be last
func good(name string, args ...int) {}
```

---

## Exercises

Four graded exercises live in [`exercises/`](./exercises/) — pure functions with table-driven tests. Remove the `t.Skip` at the top of each test, implement the function, and run `go test ./...` until it passes.

### Exercise 1: Safe Divider (`Divide(a, b float64) (float64, error)`)

Return `(quotient, err)`. On division by zero, return a non-nil error and a zero quotient; otherwise return the result and `nil`. This is the `(value, error)` backbone in miniature.

### Exercise 2: Function Composition (`Compose(f, g func(int) int) func(int) int`)

Return a new function that applies `g` then `f`, i.e. `f(g(x))`:

```go
addOne := func(n int) int { return n + 1 }
double := func(n int) int { return n * 2 }

doubleThenAddOne := Compose(addOne, double)
doubleThenAddOne(5) // (5*2)+1 = 11
```

### Exercise 3: Rate Limiter Closure (`Limiter(max int) func() error`)

Return a closure that allows up to `max` calls (each returning `nil`), then returns a non-nil error on every call after that:

```go
allow := Limiter(3)
allow() // nil  (call 1)
allow() // nil  (call 2)
allow() // nil  (call 3)
allow() // error: limit exceeded
```

### Exercise 4: Functional Options (`NewLogger(opts ...Option) *Logger`)

Build a `Logger` with the functional options pattern. Provide sensible defaults and the options `WithPrefix`, `WithLevel`, and `WithColor`:

```go
logger := NewLogger(
    WithPrefix("[APP]"),
    WithLevel("debug"),
)
// Prefix="[APP]", Level="debug", UseColor default
```

---

## Key Takeaways

1. **Multiple return values are the foundation of Go.** The `(value, error)` and `(value, ok)` shapes appear everywhere — embrace handling them immediately.

2. **Type comes after the name.** `func add(a, b int) int` — different from TS, but reads naturally once adjusted.

3. **Functions are first-class values.** Assign, pass, and return them, just like JS. Since Go 1.23, an _iterator_ is just a function of type `iter.Seq[V]` you can `range` over — more in Chapter 10.

4. **Closures capture by reference.** The famous loop-variable bug is fixed in Go 1.22+, but the by-reference rule still bites any time several closures share one mutable variable.

5. **No default params, no overloading.** The functional options pattern is the idiomatic way to get configurable defaults — and it combines variadics, first-class functions, and closures.

6. **Use named returns sparingly** — great for documentation and for `defer`-modifying the result; avoid naked returns in long functions.

---

## 🧭 Navigation

| Direction    | Link                                                                   |
| ------------ | ---------------------------------------------------------------------- |
| **Previous** | [← Chapter 04: Control Flow](../04-control-flow/README.md)             |
| **Next**     | [Chapter 06: Structs & Methods →](../06-structs-and-methods/README.md) |
