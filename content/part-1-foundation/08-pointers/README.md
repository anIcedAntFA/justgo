# Chapter 08: Pointers

> **Why pointers exist, when to use them, pointer receivers, nil pointers, and no pointer arithmetic.**

## TL;DR

Pointers let you reference and modify a value's memory location instead of copying it. JavaScript hides pointers entirely — objects are always references, primitives always copied. Go makes the choice explicit: `&` takes an address, `*` dereferences it. The good news — Go pointers are safe: no pointer arithmetic, garbage collected, and the compiler prevents most foot-guns.

---

## What Is a Pointer?

A pointer holds the **memory address** of a value, rather than the value itself.

```go
x := 42
p := &x          // p is a pointer to x; & means "address of"

fmt.Println(x)   // 42          (the value)
fmt.Println(p)   // 0xc00001... (the memory address)
fmt.Println(*p)  // 42          (* means "value at this address" — dereferencing)
```

Two operators to learn:

- `&` — "address of" — gives you a pointer to a value.
- `*` — "value at" (dereference) — gives you the value a pointer points to (and, in a type, denotes a pointer type).

```go
var x int = 42
var p *int = &x      // p is of type *int (pointer to int)

*p = 100             // change the value at p's address
fmt.Println(x)       // 100  ← x changed because p points to it
```

### The Mental Model

```
x := 42

    Variable x          Pointer p
    ┌────────┐         ┌──────────┐
    │   42   │ ◄────── │ 0xc00001 │
    └────────┘         └──────────┘
    address:           p holds the
    0xc00001           address of x

*p reads/writes the value AT that address (42 → 100)
&x gets the address of x (0xc00001)
```

> Runnable demo: [`examples/basics`](./examples/basics/main.go).

---

## Why Pointers Exist

Two main reasons:

### 1. Mutation — Modify the Original

In Go, function arguments are **passed by value** (copied). To modify the caller's variable, you need a pointer.

```go
// ❌ Without pointer — modifies a copy, original unchanged
func increment(n int) {
    n++                  // modifies the local copy
}

x := 10
increment(x)
fmt.Println(x)           // 10 — unchanged!

// ✅ With pointer — modifies the original
func increment(n *int) {
    *n++                 // dereference and modify the original
}

x := 10
increment(&x)            // pass the address
fmt.Println(x)           // 11 — changed!
```

### 2. Efficiency — Avoid Copying Large Data

Passing a large struct by value copies the entire struct. Passing a pointer copies just the address (8 bytes on a 64-bit system).

```go
type LargeStruct struct {
    Data [1000000]int    // ~8MB
}

// ❌ Copies ~8MB every call
func processValue(s LargeStruct) { /* ... */ }

// ✅ Copies 8 bytes (just the pointer)
func processPointer(s *LargeStruct) { /* ... */ }
```

> Runnable demo: [`examples/mutation`](./examples/mutation/main.go).

---

## How This Differs from JavaScript

JavaScript has pointers too — it just hides them. Understanding the JS model clarifies Go's explicit approach.

```javascript
// JavaScript: primitives copied, objects referenced (automatically)

// Primitives — copied (like Go values)
let a = 10
let b = a
b = 20
console.log(a)   // 10 — a unchanged

// Objects — referenced (like Go pointers, but implicit)
let obj1 = { count: 10 }
let obj2 = obj1
obj2.count = 20
console.log(obj1.count)   // 20 — obj1 changed! (same reference)
```

```go
// Go: YOU choose. Everything is a value unless you use a pointer.

// Value — copied
a := 10
b := a
b = 20
fmt.Println(a)   // 10

// Struct — also copied by default (unlike JS objects!)
s1 := Counter{count: 10}
s2 := s1
s2.count = 20
fmt.Println(s1.count)   // 10 — s1 unchanged! (struct was copied)

// Pointer — referenced (explicit)
p1 := &Counter{count: 10}
p2 := p1
p2.count = 20
fmt.Println(p1.count)   // 20 — p1 changed (same pointer)
```

**The critical difference:** in JavaScript, objects are ALWAYS references. In Go, structs are VALUES by default — they get copied. If you want JS-style reference behavior, you use a pointer explicitly. This is why Chapter 6's pointer receivers matter so much.

> Runnable demo: [`examples/struct-copy`](./examples/struct-copy/main.go).

---

## Pointers and Structs

This is where pointers show up most in real Go code.

```go
type User struct {
    Name string
    Age  int
}

// Create a pointer to a struct
u := &User{Name: "Alice", Age: 30}

// Go auto-dereferences for field access — write u.Name, not (*u).Name
fmt.Println(u.Name)    // Alice
u.Age = 31             // modifies the struct through the pointer
fmt.Println(u.Age)     // 31
```

Go's automatic dereferencing is a big convenience. You never write `(*u).Name` — just `u.Name`. This works for **both** field access and method calls, though the language reaches the result two different ways: field access is a selector shorthand (`u.f` is defined as `(*u).f`), while a method call resolves through the pointer's **method set** (the method set of `*User` includes methods with a `User` or `*User` receiver). Either way, you skip the `(*u)`.

### Value vs Pointer in Functions

```go
type Account struct {
    Balance float64
}

// Value receiver / parameter — gets a copy
func deposit(a Account, amount float64) {
    a.Balance += amount    // modifies the copy
}

// Pointer receiver / parameter — modifies the original
func depositPtr(a *Account, amount float64) {
    a.Balance += amount    // modifies the original
}

acc := Account{Balance: 100}

deposit(acc, 50)
fmt.Println(acc.Balance)   // 100 — unchanged (copy was modified)

depositPtr(&acc, 50)
fmt.Println(acc.Balance)   // 150 — changed (original modified)
```

This connects directly to Chapter 6's pointer receivers. A method with a pointer receiver `func (a *Account)` is the same idea — it gets a pointer to the original, so it can modify it.

---

## nil Pointers

A pointer's zero value is `nil` — it points to nothing.

```go
var p *int          // p is nil (points to nothing)
fmt.Println(p)      // <nil>

// ⚠️ Dereferencing a nil pointer PANICS:
// fmt.Println(*p)  // panic: runtime error: invalid memory address or nil pointer dereference
```

Dereferencing `nil` is one of the most common runtime panics in Go. Always check before dereferencing if a pointer might be nil:

```go
func printUser(u *User) {
    if u == nil {
        fmt.Println("no user")
        return
    }
    fmt.Println(u.Name)   // safe — u is not nil here
}
```

### nil Pointers vs JavaScript null/undefined

```javascript
// JavaScript: accessing property of null/undefined
let user = null
console.log(user.name)   // TypeError: Cannot read properties of null
// Or use optional chaining:
console.log(user?.name)  // undefined (safe)
```

```go
// Go: dereferencing nil pointer panics
var user *User = nil
// fmt.Println(user.Name)   // panic!

// Go has no optional chaining — you check explicitly:
if user != nil {
    fmt.Println(user.Name)
}
```

Go has no `?.` optional chaining. You guard with explicit nil checks. More verbose, but the nil-handling is always visible.

> Runnable demo: [`examples/nil`](./examples/nil/main.go).

---

## When to Use Pointers — Practical Guide

This is the question every Go beginner asks. Here's a practical decision guide.

### Use a pointer when:

**1. You need to modify the value**

```go
func (u *User) SetName(name string) {   // must be pointer to modify
    u.Name = name
}
```

**2. The struct is large** (avoid expensive copies)

```go
func process(data *LargeStruct) { }   // copy 8 bytes, not the whole struct
```

**3. The type's methods use pointer receivers** (consistency)

```go
// If some methods need pointers, use pointers everywhere on that type
```

**4. You need to represent "absence"** (nil as "not set")

```go
type Config struct {
    Timeout *int   // nil means "use default", non-nil means "explicitly set"
}
```

### Use a value when:

**1. The data is small** (a few fields, primitives)

```go
type Point struct{ X, Y int }   // small — pass by value is fine
func distance(a, b Point) float64 { }
```

**2. You want immutability** (the function can't modify the caller's data)

```go
func calculate(input Config) Result {   // input is a safe copy
}
```

**3. The type is naturally a value** (like `time.Time`, small structs).

### The Default Rule

When unsure, especially for structs with methods: **use pointers**. It's the more common choice in production Go, avoids copy surprises, and enables modification. The main exceptions are small immutable value types (coordinates, money amounts, time values).

---

## No Pointer Arithmetic

Unlike C, Go does NOT allow pointer arithmetic. This is a safety feature.

```go
// In C, you can do this (dangerous):
// p++           // move pointer to next memory location
// p = p + 5     // jump 5 positions

// In Go, this is a COMPILE ERROR:
p := &x
// p++           // ❌ invalid operation: p++ (non-numeric type *int)
// p = p + 1     // ❌ invalid operation
```

You cannot accidentally read or write arbitrary memory. Go pointers can only:

- Point to a valid value (via `&`).
- Be dereferenced (via `*`).
- Be compared (`==`, `!=`) — two pointers are equal when they point at the same variable, or both are `nil`. (This is also why pointers are usable as map keys.)
- Be `nil`.

This eliminates entire categories of C/C++ bugs (buffer overflows, use-after-free via arithmetic). Go pointers are "safe pointers" — references without the danger.

There's an `unsafe` package for low-level pointer manipulation (`unsafe.Pointer`, and `unsafe.Add` for offsetting), but it's for systems programming and interop — code that imports `unsafe` is non-portable and outside the Go 1 compatibility promise. You'll likely never need it.

---

## Pointers Are Garbage Collected

You never manually free memory in Go. The garbage collector handles it.

```go
func createUser() *User {
    u := &User{Name: "Alice"}   // allocated
    return u                     // returning a pointer to a local variable — totally fine!
}

user := createUser()
// Go's escape analysis detected u "escapes" the function, so it lives on the heap.
// The GC will free it when no references remain.
```

In C, returning a pointer to a local variable is a catastrophic bug (the variable's memory is reclaimed when the function returns). In Go, the compiler's **escape analysis** detects this and automatically allocates the variable on the heap so it survives. You don't think about stack vs heap — the compiler decides. You can watch it decide with `go build -gcflags='-m'` (add `-l` to disable inlining for cleaner output) — for a returned local it prints `moved to heap: u`.

This means:

- No `malloc`/`free`.
- No manual memory management.
- No use-after-free bugs.
- Returning pointers to locals is safe and common.

The tradeoff is GC overhead, but Go's collector is **concurrent** — it does most of its work alongside your program. Its stop-the-world pauses are brief and, by design, are **not proportional to heap size**; Go 1.5 already targeted pauses "almost always under 10 ms," and modern releases are typically far lower (Go 1.26's "Green Tea" collector cuts GC overhead further on GC-heavy programs). We touch on profiling memory in Chapter 34.

> Runnable demo: [`examples/escape`](./examples/escape/main.go) — build it with `-gcflags='-m'` to see the escape decision.

---

## `new()` — Allocating Pointers

Go has a built-in `new()` function that allocates zeroed memory and returns a pointer:

```go
p := new(int)        // p is *int, points to a zeroed int (0)
fmt.Println(*p)      // 0
*p = 42
fmt.Println(*p)      // 42

// new(T) is equivalent to &T{} for structs:
u := new(User)       // same as u := &User{}
fmt.Println(u.Name)  // "" (zero value)
```

In practice, you rarely use `new()`. The `&User{}` syntax is more common and lets you initialize fields:

```go
u := new(User)             // zeroed, can't set fields inline
u := &User{Name: "Alice"}  // preferred — can initialize fields
```

Reach for `&T{}` by default. `new(T)` earns its keep for non-composite types where there's no literal to take the address of — classically `new(int)`, since `&int` isn't valid syntax.

**Go 1.26 note:** `new` now also accepts a **value expression**, not just a type — `new(x)` allocates a variable initialized to `x` and returns its address. So `p := new(30)` gives you a `*int` pointing at `30`, a built-in replacement for the old `func ptr[T any](v T) *T` helper people wrote to get "a pointer to this value" (handy for the optional-field pattern below).

> Runnable demo: [`examples/new-vs-lit`](./examples/new-vs-lit/main.go).

---

## Pointers to Pointers (Rare)

Go allows `**T` (pointer to a pointer), but you'll rarely need it.

```go
x := 42
p := &x        // *int
pp := &p       // **int (pointer to pointer)

fmt.Println(**pp)   // 42 (dereference twice)
```

If you find yourself reaching for `**T`, reconsider your design — it's almost always unnecessary in idiomatic Go.

---

## Common Mistakes

### Mistake 1: Dereferencing nil

```go
// ❌ Panic
var u *User
fmt.Println(u.Name)   // panic: nil pointer dereference

// ✅ Check first
if u != nil {
    fmt.Println(u.Name)
}
```

### Mistake 2: Assuming the pre-1.22 loop-variable gotcha

```go
// Pre-1.22 gotcha: all pointers pointed to the same loop variable
var ptrs []*int
for i := 0; i < 3; i++ {
    ptrs = append(ptrs, &i)   // pre-1.22: all point to the same i!
}
// Go 1.22+: each iteration creates a new i, so this works correctly now.
```

Like the closure gotcha from Chapter 5, Go 1.22 fixed this by giving loop variables **per-iteration** scope. One nuance: it's a _language-version_ behavior, gated on the `go 1.22` (or later) directive in your `go.mod` — a module still declaring `go 1.21` keeps the old shared-variable behavior even on a newer toolchain. Know the history for old code.

### Mistake 3: Unnecessary pointers to small values

```go
// ❌ Pointless — int is tiny, no benefit
func double(n *int) int {
    return *n * 2
}

// ✅ Just pass the value
func double(n int) int {
    return n * 2
}
```

Don't use pointers for small values you're only reading. Pointers add indirection and nil-risk — only use them when you need mutation or are avoiding large copies.

### Mistake 4: Confusing the `*` in a declaration vs a dereference

```go
var p *int     // here, * is part of the TYPE (pointer to int)
x := *p        // here, * is the DEREFERENCE operator (value at p)
```

The `*` symbol does double duty: in a type (`*int`) it means "pointer to"; as an operator (`*p`) it means "dereference." Context tells them apart.

---

## Exercises

Stubs and tests live in [`exercises/`](./exercises/). Remove the `t.Skip` in each `_test.go`, implement the stub, and run `go test ./...` until it passes.

### Exercise 1: Swap

In [`exercises/swap.go`](./exercises/swap.go), implement `swap(a, b *int)` so it exchanges the two values its arguments point at. This is the smallest possible demonstration of "pass a pointer to modify the caller's variable" — swapping by value would only swap local copies.

### Exercise 2: Linked List Node

In [`exercises/list.go`](./exercises/list.go), model a singly linked list. `Node` has `Value int` and `Next *Node` — the self-reference is only possible through a pointer (a struct can't contain itself by value; that would be infinite size). Implement `Append(head *Node, value int) *Node` (returns the head, so the empty-list case works) and `Values(head *Node) []int` (walk `Next` until nil).

### Exercise 3: Optional Config

In [`exercises/settings.go`](./exercises/settings.go), implement `Resolve(s Settings) Resolved`. `Settings` uses pointer fields (`*int`, `*bool`) so `nil` means "not set — use the default" while a non-nil pointer means "explicitly set" — even when it points at the zero value. This is the difference between "absent" and "set to `0`/`false`" that a plain field can't express.

### Exercise 4: Mutate Through a Pointer

In [`exercises/counter.go`](./exercises/counter.go), implement `Counter` with a **pointer-receiver** `Increment()` and a **value-receiver** `Value() int`. The test increments a counter several times and checks the value persists — proving the pointer receiver mutates the original. Try switching `Increment` to a value receiver and watch the change vanish.

---

## Key Takeaways

1. **`&` takes an address, `*` dereferences (and denotes pointer types).** A pointer holds a memory address; dereferencing gives you the value there.

2. **Go structs are values (copied) by default — unlike JS objects (always references).** Use pointers to get reference semantics and enable mutation.

3. **Pointers exist for two reasons:** mutation (modify the original) and efficiency (avoid copying large data). A third use is representing absence (`nil` as "not set").

4. **nil pointer dereference panics.** Go has no optional chaining — check `if p != nil` explicitly.

5. **No pointer arithmetic.** Go pointers are safe — you can't accidentally access arbitrary memory. They can be taken, dereferenced, compared, and nil; `unsafe` is the rarely-needed escape hatch.

6. **The GC frees memory, and escape analysis makes returning pointers to locals safe.** No `malloc`/`free`; the compiler decides stack vs heap. Watch it with `go build -gcflags='-m'`.

7. **Prefer `&T{}` over `new(T)`;** in Go 1.26, `new(expr)` also gives you a pointer to an initialized value.

8. **When in doubt, use pointers for structs with methods.** Use values for small, immutable types. Don't use pointers for small values you only read.

---

## 🧭 Navigation

| Direction    | Link                                                           |
| ------------ | -------------------------------------------------------------- |
| **Previous** | [← Chapter 07: Interfaces](../07-interfaces/README.md)         |
| **Next**     | [Chapter 09: Error Handling →](../09-error-handling/README.md) |
