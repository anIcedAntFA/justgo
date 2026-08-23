# Chapter 06: Structs & Methods

> **Structs as data containers, methods with receivers, pointer vs value receivers, and composition over inheritance.**

## TL;DR

Go has no classes. Instead, it has structs (data) and methods (behavior attached to types). This separation is deliberate — data and behavior are decoupled. Composition replaces inheritance entirely. The trickiest new concept for JS devs: **pointer vs value receivers**, which determines whether a method can modify the struct.

---

## Structs: Data Containers

A struct is a typed collection of fields. Think of it as a TypeScript interface/type, but it's a concrete value, not just a shape.

```go
type User struct {
    Name  string
    Email string
    Age   int
}
```

### Creating Struct Instances

```go
// 1. Struct literal with field names (preferred — clear and order-independent)
u1 := User{
    Name:  "Alice",
    Email: "alice@example.com",
    Age:   30,
}

// 2. Positional (fragile — breaks if you reorder fields, avoid)
u2 := User{"Bob", "bob@example.com", 25}

// 3. Zero value struct (all fields at their zero values)
var u3 User                    // {Name:"" Email:"" Age:0}

// 4. Partial initialization (unspecified fields get zero values)
u4 := User{Name: "Charlie"}    // {Name:"Charlie" Email:"" Age:0}

// 5. Pointer to a new struct
u5 := &User{Name: "Dave"}      // *User
```

### Accessing Fields

```go
u := User{Name: "Alice", Age: 30}

fmt.Println(u.Name)     // Alice
u.Age = 31              // modify a field
fmt.Println(u.Age)      // 31

// Works the same through a pointer (Go auto-dereferences)
p := &u
fmt.Println(p.Name)     // Alice (no need for (*p).Name)
p.Age = 32              // modifies the underlying struct
```

Go automatically dereferences pointers when accessing struct fields. You write `p.Name`, not `(*p).Name`. This is a convenience Go provides — more on pointers in [Chapter 08](../08-pointers/README.md).

### JS/TS Comparison

```typescript
// TypeScript: interface defines shape, object literal creates value
interface User {
    name: string
    email: string
    age: number
}

const u: User = {
    name: "Alice",
    email: "alice@example.com",
    age: 30,
}
```

```go
// Go: struct defines type AND is the concrete value
type User struct {
    Name  string
    Email string
    Age   int
}

u := User{
    Name:  "Alice",
    Email: "alice@example.com",
    Age:   30,
}
```

Key difference: in TS, `interface` is just a compile-time shape; the actual value is a plain object. In Go, the `struct` type and the value are tightly coupled — `User` is a real type with a defined memory layout.

---

## Struct Tags

Struct fields can have "tags" — metadata strings used by libraries (especially for JSON, databases, validation).

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age,omitempty"`     // omit if zero value
    Pass  string `json:"-"`                 // never include in JSON
}
```

When you serialize this to JSON:

```go
u := User{Name: "Alice", Email: "alice@example.com", Pass: "secret"}
data, _ := json.Marshal(u)
fmt.Println(string(data))
// {"name":"Alice","email":"alice@example.com"}
// Note: Age omitted (omitempty + zero value), Pass excluded (json:"-")
```

Struct tags are how Go handles what decorators do in TypeScript. You'll use them constantly in Part 2 for JSON APIs and database mapping. We cover them deeply in [Chapter 15](../../part-2-backend/15-json-and-serialization/README.md).

---

## Methods: Behavior on Types

A method is a function with a special "receiver" — the type it's attached to.

```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Method with receiver (r Rectangle)
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Usage
rect := Rectangle{Width: 10, Height: 5}
fmt.Println(rect.Area())        // 50
fmt.Println(rect.Perimeter())   // 30
```

The `(r Rectangle)` part is the **receiver**. It's like `this` in JavaScript, but explicit and named. You choose the name (`r` here, conventionally a short abbreviation of the type).

### JS/TS Comparison

```typescript
// TypeScript: methods live inside the class, use implicit `this`
class Rectangle {
    constructor(public width: number, public height: number) {}

    area(): number {
        return this.width * this.height    // implicit this
    }
}

const rect = new Rectangle(10, 5)
rect.area()    // 50
```

```go
// Go: methods are defined separately, use explicit receiver
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height          // explicit receiver r
}

rect := Rectangle{Width: 10, Height: 5}
rect.Area()    // 50
```

Two big differences:

1. **No `this`** — you name the receiver explicitly (`r`). This makes it crystal clear what you're operating on. No `this` binding confusion, no arrow-function-vs-regular-function `this` issues that plague JavaScript.

2. **Methods are defined outside the type** — the struct definition and its methods are separate. You can even add methods to a type defined elsewhere in your package. This separation of data and behavior is fundamental to Go's design.

---

## Pointer vs Value Receivers

**This is the most important concept in this chapter.** It trips up every JS developer.

When you define a method, the receiver can be a **value** (`r Rectangle`) or a **pointer** (`r *Rectangle`). The choice matters.

### Value Receiver — Operates on a Copy

```go
func (r Rectangle) Scale(factor float64) {
    r.Width *= factor      // modifies the COPY, not the original
    r.Height *= factor
}

rect := Rectangle{Width: 10, Height: 5}
rect.Scale(2)
fmt.Println(rect.Width)    // 10  ← UNCHANGED! Scale modified a copy
```

A value receiver gets a **copy** of the struct. Changes don't affect the original.

### Pointer Receiver — Operates on the Original

```go
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor      // modifies the ORIGINAL through the pointer
    r.Height *= factor
}

rect := Rectangle{Width: 10, Height: 5}
rect.Scale(2)
fmt.Println(rect.Width)    // 20  ← CHANGED! Scale modified the original
```

A pointer receiver gets a **pointer** to the struct. Changes affect the original.

### The Mental Model

```
Value receiver  (r Type)   → method gets a COPY  → can't modify original → read-only operations
Pointer receiver (r *Type) → method gets a POINTER → can modify original → mutations
```

This concept doesn't exist in JavaScript because objects are always passed by reference. In Go, structs are **values** by default (copied when passed), and you opt into reference semantics with pointers.

```javascript
// JavaScript — objects always by reference, mutations always stick
class Rectangle {
    scale(factor) {
        this.width *= factor    // always modifies the original
    }
}
```

```go
// Go — you CHOOSE: copy (value) or reference (pointer)
func (r Rectangle) ScaleCopy(f float64)  { r.Width *= f }  // copy
func (r *Rectangle) ScaleReal(f float64) { r.Width *= f }  // original
```

### When to Use Which

**Use a pointer receiver when:**

- The method needs to modify the struct
- The struct is large (copying is expensive)
- For consistency — if any method needs a pointer, use pointers for all methods on that type

**Use a value receiver when:**

- The method only reads (doesn't modify)
- The struct is small (a few fields)
- The type is naturally a value (like `time.Time`)

**The practical rule:** When in doubt, use a pointer receiver. It's the more common choice in real Go code, and it avoids surprises with mutations. Most Go style guides recommend: pick one (usually pointer) and use it consistently for all methods on a type.

```go
// Consistent — all pointer receivers
func (u *User) SetName(name string)  { u.Name = name }
func (u *User) GetName() string      { return u.Name }
func (u *User) Validate() error      { /* ... */ }
```

### Go Auto-Handles the & and *

A convenience: Go automatically takes the address or dereferences as needed:

```go
type Counter struct{ count int }

func (c *Counter) Increment() { c.count++ }   // pointer receiver

c := Counter{}        // c is a value, not a pointer
c.Increment()         // Go automatically does (&c).Increment()
fmt.Println(c.count)  // 1

// Works because c is addressable. You don't write (&c).Increment().
```

This auto-conversion only works when the value is "addressable" (stored in a variable). You can't call a pointer method on a literal: `Counter{}.Increment()` won't compile.

---

## Constructors (There's No `new` Keyword for This)

Go has no constructors like `new MyClass()`. The convention is a function named `NewTypeName`:

```go
type Server struct {
    host    string
    port    int
    timeout time.Duration
}

// Convention: NewXxx constructor function
func NewServer(host string, port int) *Server {
    return &Server{
        host:    host,
        port:    port,
        timeout: 30 * time.Second,   // sensible default
    }
}

// Usage
server := NewServer("localhost", 8080)
```

Why a function instead of a keyword?

- You can validate inputs and return an error: `func NewServer(...) (*Server, error)`
- You can set defaults
- You control exactly how the struct is initialized
- It's just a regular function — no magic

```go
// Constructor that can fail
func NewServer(host string, port int) (*Server, error) {
    if port < 1 || port > 65535 {
        return nil, fmt.Errorf("invalid port: %d", port)
    }
    return &Server{host: host, port: port, timeout: 30 * time.Second}, nil
}

server, err := NewServer("localhost", 8080)
if err != nil {
    log.Fatal(err)
}
```

There IS a built-in `new()` function, but it's rarely used directly — it just allocates zeroed memory and returns a pointer: `new(Server)` is equivalent to `&Server{}`.

---

## Composition Over Inheritance

Go has **no inheritance**. No `extends`, no `super`, no class hierarchy. Instead, Go uses **composition** through struct embedding.

### Struct Embedding

```go
type Animal struct {
    Name string
}

func (a Animal) Eat() {
    fmt.Printf("%s is eating\n", a.Name)
}

// Dog embeds Animal (composition, not inheritance)
type Dog struct {
    Animal          // embedded field (no field name — just the type)
    Breed string
}

func (d Dog) Bark() {
    fmt.Printf("%s barks!\n", d.Name)   // Name is promoted from Animal
}

// Usage
d := Dog{
    Animal: Animal{Name: "Rex"},
    Breed:  "Labrador",
}

d.Eat()              // "Rex is eating" — Eat is promoted from Animal
d.Bark()             // "Rex barks!"
fmt.Println(d.Name)  // "Rex" — Name is promoted from Animal
fmt.Println(d.Animal.Name)  // "Rex" — can also access explicitly
```

When you embed `Animal` in `Dog`, all of `Animal`'s fields and methods are **promoted** to `Dog`. You can call `d.Eat()` even though `Eat` is defined on `Animal`. This looks like inheritance, but it's fundamentally different.

### Why Composition ≠ Inheritance

```typescript
// TypeScript inheritance: Dog IS-A Animal
class Animal {
    constructor(public name: string) {}
    eat() { console.log(`${this.name} eats`) }
}

class Dog extends Animal {
    bark() { console.log(`${this.name} barks`) }
}

// Dog can be used anywhere Animal is expected (polymorphism via inheritance)
const a: Animal = new Dog("Rex")   // valid — Dog IS-A Animal
```

```go
// Go composition: Dog HAS-A Animal
type Dog struct {
    Animal       // Dog contains an Animal
    Breed string
}

// Dog is NOT an Animal. You can't use Dog where Animal is expected
// (unless via interfaces — see Chapter 07).
// var a Animal = Dog{...}   // ❌ won't compile
```

The distinction:

- **Inheritance (IS-A):** Dog IS an Animal. Tight coupling. Changes to Animal ripple to all subclasses.
- **Composition (HAS-A):** Dog HAS an Animal. Loose coupling. Dog uses Animal's capabilities but isn't bound to its identity.

Go's designers believe deep inheritance hierarchies are a source of fragility (the "fragile base class problem," diamond inheritance, etc.). Composition is flatter, more explicit, and easier to reason about.

### Embedding Multiple Types

```go
type Logger struct{ prefix string }
func (l Logger) Log(msg string) { fmt.Printf("%s: %s\n", l.prefix, msg) }

type Validator struct{}
func (v Validator) Validate() bool { return true }

// Service composes both
type Service struct {
    Logger
    Validator
    name string
}

s := Service{
    Logger: Logger{prefix: "SVC"},
    name:   "payment",
}
s.Log("starting")    // "SVC: starting" — from Logger
s.Validate()         // true — from Validator
```

Multiple embedding gives you the capabilities of multiple types. No diamond problem because Go resolves conflicts at compile time (ambiguous promotions are errors you must resolve explicitly).

### Overriding Promoted Methods

```go
type Base struct{}
func (b Base) Describe() string { return "I am Base" }

type Derived struct {
    Base
}
func (d Derived) Describe() string { return "I am Derived" }   // "overrides" Base.Describe

d := Derived{}
fmt.Println(d.Describe())        // "I am Derived" — Derived's method wins
fmt.Println(d.Base.Describe())   // "I am Base" — explicit access to Base's method
```

The outer type's method "shadows" the embedded type's method. This looks like method overriding, but remember — there's no virtual dispatch through the embedded type. It's just name resolution.

---

## Anonymous Structs

Go lets you create structs without naming the type — useful for one-off data:

```go
// Anonymous struct — defined and instantiated at once
config := struct {
    Host string
    Port int
}{
    Host: "localhost",
    Port: 8080,
}

fmt.Println(config.Host)   // localhost
```

Common use cases: table-driven tests ([Chapter 13](../13-testing/README.md)), one-off JSON responses, temporary groupings. Don't overuse — named structs are clearer for anything reused.

---

## Comparing Structs

Structs are comparable with `==` if all their fields are comparable:

```go
type Point struct{ X, Y int }

p1 := Point{1, 2}
p2 := Point{1, 2}
p3 := Point{3, 4}

fmt.Println(p1 == p2)   // true  — same field values
fmt.Println(p1 == p3)   // false

// Structs work as map keys (if comparable)
seen := map[Point]bool{}
seen[Point{1, 2}] = true
fmt.Println(seen[Point{1, 2}])   // true
```

This is impossible in JavaScript — `{x:1} === {x:1}` is always `false` (reference comparison). In Go, struct equality compares field values. Note: structs containing slices, maps, or functions are NOT comparable (those types aren't comparable), and comparing them with `==` is a compile error.

---

## Runnable Examples

The [`examples/`](./examples/) folder has three programs. Run each with `go run .` from its directory:

- **[`examples/receivers`](./examples/receivers/)** — the same `Scale` method written with a value receiver and a pointer receiver, side by side, so you can watch one mutate the original and the other quietly modify a copy.
- **[`examples/embedding`](./examples/embedding/)** — composition through struct embedding: promoted fields and methods, multiple embedding, and an outer method shadowing an embedded one.
- **[`examples/builder`](./examples/builder/)** — a `NewXxx` constructor with defaults plus method chaining (the builder pattern), which only works because the mutating methods use pointer receivers.

---

## Common Mistakes

### Mistake 1: Value receiver when you meant to modify

```go
// ❌ Bug — modification doesn't persist
func (u User) SetAge(age int) {
    u.Age = age      // modifies a copy
}
user.SetAge(30)
fmt.Println(user.Age)   // still old value!

// ✅ Use pointer receiver to modify
func (u *User) SetAge(age int) {
    u.Age = age      // modifies the original
}
```

This is the #1 struct bug for newcomers. If a method changes the struct, it needs a pointer receiver.

### Mistake 2: Mixing value and pointer receivers on the same type

```go
// ⚠️ Inconsistent — confusing and can cause interface issues
func (u User) GetName() string  { return u.Name }   // value
func (u *User) SetName(s string) { u.Name = s }      // pointer

// ✅ Pick one (usually pointer) and be consistent
func (u *User) GetName() string  { return u.Name }
func (u *User) SetName(s string) { u.Name = s }
```

### Mistake 3: Positional struct literals

```go
// ❌ Fragile — breaks silently if you reorder/add fields
u := User{"Alice", "alice@x.com", 30}

// ✅ Named fields — explicit and refactor-safe
u := User{Name: "Alice", Email: "alice@x.com", Age: 30}
```

### Mistake 4: Expecting inheritance behavior

```go
// Go embedding is NOT inheritance.
// var a Animal = Dog{}   // ❌ won't compile — Dog is not an Animal
// Use interfaces for polymorphism (next chapter).
```

---

## Exercises

Four graded exercises live in [`exercises/`](./exercises/) — structs and methods with table-driven tests. Remove the `t.Skip` at the top of each test, implement the code, and run `go test ./...` until it passes.

### Exercise 1: Bank Account (`account.go`)

Model a `BankAccount` struct with `Owner` (string) and `Balance` (float64). Implement `Deposit(amount float64) error` (errors on a negative amount), `Withdraw(amount float64) error` (errors on a negative amount or insufficient funds), and `String() string` (a formatted summary). Decide — and justify in your own head — which receiver type each method needs.

### Exercise 2: Shape Composition (`shapes.go`)

Create a `Shape` struct with a `Name` field and a `Describe()` method. Create `Circle` (embeds `Shape`, adds `Radius`) and `Square` (embeds `Shape`, adds `Side`), each with its own `Area()`. Verify that `Describe()` is promoted from the embedded `Shape`.

### Exercise 3: Builder with Constructor (`httpclient.go`)

Build an `HTTPClient` with fields `baseURL`, `timeout`, `retries`. Write `NewHTTPClient(baseURL string) *HTTPClient` with defaults (`timeout=30s`, `retries=3`), plus `SetTimeout` and `SetRetries` that return `*HTTPClient` for chaining:

```go
client := NewHTTPClient("https://api.example.com").
    SetTimeout(60 * time.Second).
    SetRetries(5)
```

The builder pattern via method chaining — pointer receivers are essential here.

### Exercise 4: Struct Equality Cache (`coordinate.go`)

Model a `Coordinate{Lat, Lng float64}` and a `PlaceCache` (a `map[Coordinate]string`) with `Set` and `Get`. Demonstrate that two separately-created coordinates with the same field values hit the **same** cache entry — struct value equality as a map key.

---

## Key Takeaways

1. **Structs are data, methods are behavior — and they're separate.** No classes bundling everything together. This decoupling is intentional.

2. **Pointer vs value receiver is the key decision.** Pointer receiver (`*Type`) to modify the original or for large structs. Value receiver for small read-only types. When in doubt, use pointer, and be consistent across a type's methods.

3. **No `this` — you name the receiver.** Explicit, clear, no binding surprises.

4. **Composition replaces inheritance.** Embed structs to promote their fields and methods. But embedding is HAS-A, not IS-A — for polymorphism, use interfaces (next chapter).

5. **Constructors are `NewXxx` functions**, not keywords. They can validate, set defaults, and return errors.

6. **Structs are comparable by value** with `==` (if fields are comparable) and can be map keys — impossible in JavaScript.

---

## 🧭 Navigation

| Direction    | Link                                                   |
| ------------ | ------------------------------------------------------ |
| **Previous** | [← Chapter 05: Functions](../05-functions/README.md)   |
| **Next**     | [Chapter 07: Interfaces →](../07-interfaces/README.md) |
