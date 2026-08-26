# Chapter 07: Interfaces

> **Implicit satisfaction, the empty interface, type assertions, type switches, and the io.Reader/Writer pattern.**

## TL;DR

Interfaces are the heart of Go's polymorphism — and they work fundamentally differently from TypeScript. In Go, a type satisfies an interface **automatically** just by having the right methods. No `implements` keyword. This "implicit satisfaction" is the single most powerful and most alien concept for a TS developer. Master interfaces and you understand idiomatic Go.

---

## What Is an Interface?

An interface is a set of method signatures. Any type that has those methods "satisfies" the interface — automatically.

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

This says: "A Shape is anything that has an `Area() float64` method and a `Perimeter() float64` method."

```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }
```

Both `Rectangle` and `Circle` have `Area()` and `Perimeter()` methods. Therefore, both **automatically** satisfy the `Shape` interface. No declaration needed.

```go
// Both can be used as a Shape
shapes := []Shape{
    Rectangle{Width: 10, Height: 5},
    Circle{Radius: 3},
}

for _, s := range shapes {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}
```

This is polymorphism in Go. Different types, same interface, uniform handling.

> Runnable demo: [`examples/satisfaction`](./examples/satisfaction/main.go).
>
> Before/after demo — a coffee shop with and without an interface:
> [`examples/coffee`](./examples/coffee/main.go).

> **TODO (owner):** write the "before an interface vs. after" prose here — brewing
> coffee as a string-keyed `switch` (`examples/coffee/before.go`) vs. a `Brewer`
> interface (`examples/coffee/after.go`): open/closed, implicit satisfaction for a
> brewer you don't own, and compile-time safety over stringly-typed methods.

---

## Implicit Satisfaction — The Big Difference

This is THE concept to internalize. Coming from TypeScript, you explicitly declare `implements`:

```typescript
// TypeScript: explicit implementation
interface Shape {
    area(): number
    perimeter(): number
}

class Rectangle implements Shape {   // ← explicit "implements Shape"
    constructor(public width: number, public height: number) {}
    area(): number { return this.width * this.height }
    perimeter(): number { return 2 * (this.width + this.height) }
}
```

```go
// Go: implicit satisfaction — NO "implements" keyword
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width, Height float64
}
// Just by having these methods, Rectangle IS a Shape. Automatically.
func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }
```

### Why This Matters

The power of implicit satisfaction:

**1. You can satisfy interfaces you don't own.** You can make your type satisfy a standard library interface (like `io.Reader`, `fmt.Stringer`, `error`) without importing anything special — just implement the methods.

**2. Decoupling.** The type doesn't need to know the interface exists. You can define an interface in package A that's satisfied by a type in package B, even though B never heard of A. The interface lives where it's **used**, not where the type is **defined**.

**3. Retroactive interfaces.** You can define a new interface today that existing types already satisfy. No need to modify those types.

```go
// The standard library defines this:
// type Stringer interface { String() string }

// Your type satisfies it just by having a String() method:
type Temperature float64

func (t Temperature) String() string {
    return fmt.Sprintf("%.1f°C", float64(t))
}

// Now Temperature works anywhere a Stringer is expected,
// including fmt.Println (which checks for Stringer):
temp := Temperature(25.5)
fmt.Println(temp)   // "25.5°C" — fmt automatically used your String() method
```

You never wrote `implements Stringer`. You just implemented `String()`, and Go connected the dots.

### "Accept Interfaces, Return Structs"

A famous Go guideline. Functions should accept interfaces (flexible inputs) but return concrete types (specific outputs):

```go
// Accept an interface — caller can pass anything that satisfies it
func PrintArea(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
}

// Return a concrete type — caller knows exactly what they get
func NewRectangle(w, h float64) *Rectangle {
    return &Rectangle{Width: w, Height: h}
}
```

This makes your functions maximally flexible (accept any matching type) while keeping return values concrete and predictable.

> This is a **heuristic, not a law.** The standard library deliberately returns interfaces when the whole point is to hide or swap the implementation — `net.Dial` returns `net.Conn`, and `crc32.NewIEEE` / `adler32.New` both return `hash.Hash32` so you can change the algorithm by changing only the constructor call. Prefer concrete returns; return an interface when you mean to abstract the implementation away.

---

## The Power of Small Interfaces

Go's most famous interfaces have a single method. This is intentional — "the bigger the interface, the weaker the abstraction."

### io.Reader and io.Writer

The two most important interfaces in Go:

```go
// From the standard library:
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

One method each. Yet these power an enormous portion of Go's standard library. Anything that can be read from is a `Reader`. Anything that can be written to is a `Writer`:

```
io.Reader is satisfied by:
  - os.File (reading files)
  - http.Request.Body (reading HTTP request bodies)
  - bytes.Buffer (reading from memory)
  - strings.Reader (reading from a string)
  - net.Conn (reading from network connections)
  - gzip.Reader (reading compressed data)
  ... and hundreds more

io.Writer is satisfied by:
  - os.File (writing files)
  - os.Stdout (writing to terminal)
  - http.ResponseWriter (writing HTTP responses)
  - bytes.Buffer (writing to memory)
  - gzip.Writer (writing compressed data)
  ... and hundreds more
```

Because they all share the same interface, you can connect them like pipes:

```go
// io.Copy copies from ANY reader to ANY writer — files, network, memory, etc.
// Its real signature (you call it as io.Copy(...), it is declared as Copy):
//
//	func Copy(dst Writer, src Reader) (written int64, err error)

// Examples — all using the same io.Copy:
io.Copy(os.Stdout, file)               // file → terminal
io.Copy(httpResponse, file)            // file → HTTP response
io.Copy(gzipWriter, file)              // file → compressed output
io.Copy(file, httpRequest.Body)        // HTTP upload → file
```

This is the magic of small interfaces. `io.Copy` doesn't know or care what the source and destination are — files, network sockets, memory buffers, compressed streams. As long as they satisfy `Reader` and `Writer`, it works.

### A Practical Example

```go
import (
    "bytes"
    "fmt"
    "io"
    "log"
    "os"
    "strings"
)

func main() {
    // Read from a string, write to stdout
    src := strings.NewReader("Hello from a string reader!\n")
    if _, err := io.Copy(os.Stdout, src); err != nil {  // prints to terminal
        log.Fatal(err)
    }

    // Read from a string, write to a buffer (memory)
    src2 := strings.NewReader("captured in memory")
    var buf bytes.Buffer
    if _, err := io.Copy(&buf, src2); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Buffer now contains:", buf.String())
}
```

Same `io.Copy`, completely different source/destination types. This composability is why Go's standard library feels so cohesive.

> Runnable demo: [`examples/io-writer`](./examples/io-writer/main.go).

---

## Interface Embedding

Interfaces compose by **embedding** other interfaces. An interface can list another interface as an element; its method set becomes the union (the type set becomes the intersection). This is how the standard library builds `io.ReadWriter`, `io.ReadCloser`, and `io.ReadWriteCloser` out of the one-method `Reader`, `Writer`, and `Closer`:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

// ReadWriteCloser is just the three small interfaces embedded together.
// A type satisfies it by having Read, Write, AND Close.
type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```

Rule: when embedding interfaces, methods with the same name must have **identical signatures**. Embedding keeps interfaces small at the point of definition while letting callers ask for exactly the combined behavior they need.

> **TODO (owner):** expand the prose here — when to embed vs. re-declare, and the `os.File` → `io.ReadWriteCloser` payoff.
>
> Runnable demo: [`examples/embedding`](./examples/embedding/main.go).

---

## The Empty Interface and `any`

An interface with no methods is satisfied by **every** type (since every type has at least zero methods):

```go
// Pre-Go 1.18: interface{}
// Go 1.18+: any (a predeclared alias for interface{})

var x any
x = 42
x = "hello"
x = []int{1, 2, 3}
x = User{Name: "Alice"}
// x can hold anything
```

`any` is Go's equivalent to TypeScript's `any` / `unknown`. Use it sparingly — it throws away type safety.

```go
// A function accepting anything
func describe(i any) {
    fmt.Printf("value: %v, type: %T\n", i, i)
}

describe(42)            // value: 42, type: int
describe("hello")       // value: hello, type: string
describe([]int{1, 2})   // value: [1 2], type: []int
```

**When to use `any`:**

- Functions that genuinely handle arbitrary types (like `fmt.Println`)
- Decoding unknown JSON structures
- Before generics existed, for generic containers

**When NOT to use `any`:**

- When you know the types — use specific types or generics (Chapter 12)
- As a lazy escape hatch — you lose all compile-time safety

Coming from TypeScript, resist the urge to reach for `any`. In modern Go (1.18+), generics often replace what `any` used to do.

---

## Type Assertions

When you have an `any` (or any interface), you sometimes need to get the concrete type back. That's a type assertion:

```go
var i any = "hello"

// Type assertion — extract the concrete type
s := i.(string)        // s is now a string "hello"
fmt.Println(s)

// ⚠️ If the type is wrong, this PANICS:
// n := i.(int)        // panic: interface conversion: interface {} is string, not int
```

### The Safe "Comma, ok" Form

To avoid panics, use the two-value form:

```go
var i any = "hello"

s, ok := i.(string)
if ok {
    fmt.Println("it's a string:", s)
}

n, ok := i.(int)
if !ok {
    fmt.Println("not an int")    // this runs — i is a string
}
```

The `ok` boolean tells you whether the assertion succeeded. Always use this form unless you're certain of the type. It's the same safety pattern as map lookups: `value, ok`.

> When a comma-ok assertion **fails**, the value is not left undefined — it is the **zero value** of the asserted type. A failed `n, ok := i.(int)` leaves `n == 0`, not garbage.

### JS Comparison

```typescript
// TypeScript: type guards / assertions
let i: unknown = "hello"

if (typeof i === "string") {
    console.log(i.toUpperCase())   // TS narrows the type
}

// Or assertion (unsafe):
const s = i as string
```

```go
// Go: type assertion with comma-ok
var i any = "hello"

if s, ok := i.(string); ok {
    fmt.Println(strings.ToUpper(s))   // Go narrows the type
}
```

Similar concept. TypeScript uses `typeof` guards and `as`; Go uses `.(Type)` assertions with the `ok` pattern.

---

## Type Switches

When you need to handle many possible types, a type switch is cleaner than chained assertions:

```go
func describe(i any) string {
    switch v := i.(type) {
    case nil:
        return "nil value"
    case int:
        return fmt.Sprintf("int: %d", v)
    case string:
        return fmt.Sprintf("string of length %d: %q", len(v), v)
    case bool:
        return fmt.Sprintf("bool: %t", v)
    case []int:
        return fmt.Sprintf("int slice with %d elements", len(v))
    case User:
        return fmt.Sprintf("User named %s", v.Name)
    default:
        return fmt.Sprintf("unknown type: %T", v)
    }
}

fmt.Println(describe(42))           // int: 42
fmt.Println(describe("hello"))      // string of length 5: "hello"
fmt.Println(describe(true))         // bool: true
fmt.Println(describe([]int{1, 2}))  // int slice with 2 elements
```

The key part is `switch v := i.(type)`. Inside each `case`, `v` has the specific type, so you can use it directly. This is the special "type switch" syntax — note `.(type)` only works inside a switch.

> Runnable demo: [`examples/type-switch`](./examples/type-switch/main.go).

---

## The error Interface

You've been using errors since Chapter 5. Here's the secret: `error` is just an interface.

```go
// The entire definition of error in Go:
type error interface {
    Error() string
}
```

Any type with an `Error() string` method is an error. That's it. This is why you can create custom error types:

```go
type ValidationError struct {
    Field   string
    Message string
}

// Implement the error interface
func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Message)
}

// Now *ValidationError IS an error, usable anywhere an error is expected
func validateAge(age int) error {
    if age < 0 {
        return &ValidationError{Field: "age", Message: "must be non-negative"}
    }
    return nil
}

err := validateAge(-5)
if err != nil {
    fmt.Println(err)   // "validation failed on age: must be non-negative"
}
```

We cover error handling in depth in Chapter 9. For now, just internalize: **errors are interface values**, and you can make any type an error by giving it an `Error()` method.

---

## The Stringer Interface

Another tiny but ubiquitous interface — controls how a type prints:

```go
type Stringer interface {
    String() string
}
```

When you `fmt.Println` a value, `fmt` checks if it implements `Stringer`. If so, it calls `String()`:

```go
type Color struct {
    R, G, B uint8
}

func (c Color) String() string {
    return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}

c := Color{R: 255, G: 128, B: 0}
fmt.Println(c)   // "#FF8000" — fmt called your String() method
```

This is like overriding `toString()` in JavaScript, but interface-based. Implement `String()` and your type prints nicely everywhere.

---

## Interface Values and nil — A Subtle Gotcha

This is an advanced gotcha, but it bites people. An interface value has two parts: a **type** and a **value**. An interface is `nil` only if BOTH are nil.

```go
type MyError struct{}
func (e *MyError) Error() string { return "my error" }

func doSomething() error {
    var err *MyError = nil     // a nil pointer
    return err                  // ⚠️ returns a non-nil interface!
}

err := doSomething()
if err != nil {
    fmt.Println("error is not nil!")   // THIS RUNS — surprising!
}
```

Why? The returned interface has a type (`*MyError`) and a value (`nil`). Since the **type** part is non-nil, the interface itself is non-nil — even though the underlying pointer is nil.

**The fix:** return `nil` literally, not a typed nil pointer:

```go
func doSomething() error {
    // ... if no error:
    return nil      // ✅ truly nil interface
}
```

This is one of Go's genuinely tricky corners. The rule: don't return typed nil pointers as errors. Return the literal `nil`. We'll revisit this in Chapter 9.

---

## Designing Good Interfaces

A few principles for when you start defining your own:

**1. Keep them small.** One to three methods. The standard library's best interfaces (`io.Reader`, `io.Writer`, `fmt.Stringer`, `error`) have exactly one method.

**2. Define interfaces where they're used, not where types are defined.** In Go, the consumer defines the interface. If your function needs "something that can be read," define a small `Reader` interface in your package — don't force the producer to declare it.

**3. Don't create interfaces speculatively.** Unlike Java/C# where you might define interfaces upfront "just in case," Go culture says: write concrete types first, extract interfaces only when you actually need polymorphism (e.g., for testing or multiple implementations).

```go
// ❌ Premature — interface with one implementation and no test need
type UserGetter interface {
    GetUser(id int) (*User, error)
}

// ✅ Start concrete, extract interface when a second implementation
//    or a test mock actually requires it
type UserService struct { db *sql.DB }
func (s *UserService) GetUser(id int) (*User, error) { /* ... */ }
```

This is the opposite of how many enterprise Java/TS codebases work. In Go, interfaces emerge from need, not from upfront design.

---

## Common Mistakes

### Mistake 1: Reaching for `any` instead of proper types

```go
// ❌ Throwing away type safety
func process(data any) any { /* ... */ }

// ✅ Use concrete types or generics
func process[T any](data T) T { /* ... */ }
```

### Mistake 2: Unchecked type assertions

```go
// ❌ Panics if wrong type
s := value.(string)

// ✅ Safe form
s, ok := value.(string)
if !ok { /* handle */ }
```

### Mistake 3: Big interfaces

```go
// ❌ Too many methods — weak abstraction, hard to satisfy/mock
type Repository interface {
    Create(...) error
    Read(...) error
    Update(...) error
    Delete(...) error
    List(...) error
    Search(...) error
    Count(...) error
    // ... 10 more methods
}

// ✅ Split into focused interfaces (Interface Segregation)
type Reader interface { Read(id int) (*Item, error) }
type Writer interface { Write(item *Item) error }
```

### Mistake 4: Defining interfaces before you need them

Write concrete code first. Extract an interface when you have a real second implementation or a testing need. Speculative interfaces add indirection without benefit.

---

## Exercises

Stubs and tests live in [`exercises/`](./exercises/). Remove the `t.Skip` in each
`_test.go` as you solve it.

### Exercise 1: Shape Interface

Define a `Shape` interface with `Area()` and `Perimeter()`. Implement `Rectangle`, `Circle`, and `Triangle`. Write a function `TotalArea(shapes []Shape) float64` that sums the areas of a mixed slice of shapes.

### Exercise 2: Custom Stringer

Create a `Duration` type (based on int, representing seconds). Implement `String()` so it prints human-readable: `3661` → `"1h 1m 1s"`. Verify `fmt.Println` uses your method.

### Exercise 3: Writer Composition

Write a function that takes an `io.Writer` and writes a formatted report to it. Then call it with different writers: `os.Stdout` (terminal), a `bytes.Buffer` (memory), and an `os.File` (disk). Same function, three destinations.

```go
func WriteReport(w io.Writer, data []string) error {
    // write a formatted report to w
}
```

### Exercise 4: Type Switch Value Printer

Write a function `Describe(v any) string` that uses a type switch to render values differently based on type:

- `string` → wrapped in quotes
- `int`, `float64` → as numbers
- `bool` → "yes"/"no"
- `[]any` → recursively render each element
- `map[string]any` → render key: value pairs
- everything else → `%v` with type name

(This is essentially a mini value pretty-printer — great practice for type switches.)

### Exercise 5: Brewer Interface & a Fake for Testing

In [`exercises/brewer.go`](./exercises/brewer.go), implement `Serve(b Brewer, g Grounds) Cup` — it runs a brewer over the grounds and returns the cup. The point isn't `Serve`'s one-line body; it's _why_ it takes a `Brewer` interface: the provided test passes a **`fakeBrewer`** instead of a real machine and asserts `Serve` calls `Brew` and returns its cup. Accepting an interface is what makes the code testable. Remove the `t.Skip` in `brewer_test.go` once `Serve` is implemented.

---

## Key Takeaways

1. **Implicit satisfaction is the defining feature.** Types satisfy interfaces just by having the right methods — no `implements`. This enables decoupling and retroactive interfaces.

2. **Small interfaces are powerful.** `io.Reader` and `io.Writer` have one method each yet power the whole standard library. Aim for 1-3 methods.

3. **Compose with embedding.** Bigger interfaces like `io.ReadWriteCloser` are built by embedding small ones — same-name methods must share identical signatures.

4. **"Accept interfaces, return structs."** Flexible inputs, concrete outputs — a strong default, with deliberate stdlib exceptions when the goal is to hide the implementation.

5. **`any` (empty interface) holds anything — use sparingly.** Prefer concrete types or generics. Always use the comma-ok form for type assertions.

6. **Type switches** handle multiple possible types cleanly with `switch v := i.(type)`.

7. **error and Stringer are just interfaces** with one method each. Implement them on your own types for custom errors and custom printing.

8. **Define interfaces where they're used, when you need them.** Don't design them speculatively. Concrete first, abstract later.

---

## 🧭 Navigation

| Direction    | Link                                                                   |
| ------------ | ---------------------------------------------------------------------- |
| **Previous** | [← Chapter 06: Structs & Methods](../06-structs-and-methods/README.md) |
| **Next**     | [Chapter 08: Pointers →](../08-pointers/README.md)                     |
