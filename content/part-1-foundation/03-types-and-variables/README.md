# Chapter 03: Types & Variables

> **Go's type system, variable declarations, zero values, and the "no undefined" guarantee.**

## TL;DR

Go is statically typed like TypeScript, but simpler — no union types, no mapped types, no conditional types. Every variable has exactly one type, determined at compile time. The key mind-shift: Go has **zero values** instead of `null`/`undefined`, and **no implicit type coercion** ever. What you declare is what you get.

---

## Variable Declarations

Go has three ways to declare variables. Each has a specific use case.

### 1. `var` with explicit type

```go
var name string
var age int
var isActive bool
```

Use when you want a variable at its zero value, or at package level.

### 2. `var` with initialization

```go
var name string = "Go Developer"
var age int = 19

// Type can be inferred — compiler figures it out
var name = "Go Developer"    // inferred as string
var age = 19                 // inferred as int
```

### 3. Short declaration `:=` (most common)

```go
name := "Go Developer"      // string
age := 19                   // int
isActive := true             // bool
ratio := 3.14               // float64
```

`:=` is the one you'll use 90% of the time inside functions. It declares AND assigns in one step with type inference.

**Critical rule:** `:=` only works inside functions. Package-level variables must use `var`.

```go
package main

var globalConfig = "production"     // ✅ var at package level

func main() {
    localName := "Go Developer"    // ✅ := inside function
}

// name := "broken"                // ❌ := at package level — compile error
```

### JS Comparison

```javascript
// JavaScript: three keywords, different scoping rules
var name = "old school"       // function-scoped, hoisted
let age = 25                  // block-scoped
const MAX = 100               // block-scoped, can't reassign

// Plus the chaos:
let x;                        // x is undefined
console.log(y)                // ReferenceError (or undefined if var)
```

```go
// Go: two forms, one set of rules
var name string               // zero value: "" (not undefined, not null)
age := 25                     // short declaration
const MAX = 100               // constant

// No chaos:
var x int                     // x is 0 (always initialized)
// fmt.Println(y)             // compile error — y doesn't exist
```

No hoisting. No temporal dead zone. No `undefined`. If a variable exists, it has a value.

---

## Basic Types

Go has a small, fixed set of basic types. No `any` by default, no union types.

### Numbers

```go
// Integers — signed
var i int       // platform-dependent: 64-bit on modern systems
var i8 int8     // -128 to 127
var i16 int16   // -32768 to 32767
var i32 int32   // -2,147,483,648 to 2,147,483,647
var i64 int64   // -9.2 quintillion to 9.2 quintillion

// Integers — unsigned
var u uint      // platform-dependent
var u8 uint8    // 0 to 255
var u16 uint16  // 0 to 65535
var u32 uint32  // 0 to 4,294,967,295
var u64 uint64  // 0 to 18.4 quintillion

// Floating point
var f32 float32 // ~7 decimal digits precision
var f64 float64 // ~15 decimal digits precision (default for float literals)
```

**Which integer type to use?** Just use `int` unless you have a specific reason (binary protocols, memory optimization, interop with C). `int` is 64-bit on any modern system.

In JS, `number` is always a 64-bit float — there's no distinction between `42` and `42.0`. In Go, `42` is an `int` and `42.0` is a `float64`. They're different types and you cannot mix them without explicit conversion.

### Strings

```go
name := "Hello, Go"             // double quotes — interpreted string (escapes work)
raw := `Hello\nGo`              // backticks — raw string (escapes are literal)

fmt.Println(name)   // Hello, Go     (with newline)
fmt.Println(raw)    // Hello\nGo     (literal backslash n)
```

Raw strings with backticks are especially useful for regex, SQL queries, JSON templates, and multi-line strings — same as template literals in JS but without interpolation.

**Strings are immutable** in Go, just like in JavaScript. You can't modify individual characters.

**Strings are UTF-8 encoded bytes**, not arrays of characters:

```go
s := "Hello, 世界"
fmt.Println(len(s))         // 13 (bytes, not characters!)

// To count actual characters (runes):
import "unicode/utf8"
fmt.Println(utf8.RuneCountInString(s))  // 9 (characters)

// Iterate over characters (runes), not bytes:
for i, ch := range s {
    fmt.Printf("index %d: %c (U+%04X)\n", i, ch, ch)
}
```

This is a gotcha coming from JS where `"Hello, 世界".length` gives 9 (characters). In Go, `len()` on a string gives **bytes**. The `range` loop over a string, however, iterates over **runes** (Unicode code points), not bytes.

### bool

```go
isReady := true
isComplete := false

// No truthy/falsy. This is a compile error:
// if name { ... }        // ❌ string is not bool
// if count { ... }       // ❌ int is not bool

// Must be explicit:
if name != "" { ... }     // ✅
if count > 0 { ... }      // ✅
```

Coming from JS where `if (name)` works on any type — Go forces you to be explicit. `0`, `""`, `nil` are NOT falsy. Only `bool` values work in conditions. This eliminates an entire class of subtle bugs.

### byte and rune

These are aliases, not new types:

```go
// byte = uint8 — raw binary data, ASCII characters
var b byte = 'A'          // 65

// rune = int32 — Unicode code point, any character
var r rune = '世'          // 19990
var r2 rune = 'A'         // 65

// Single quotes = rune literal (NOT string)
// 'A' is a rune (int32), "A" is a string
```

In JS, there's no character type — `'A'` and `"A"` are both strings. In Go, single quotes create a `rune` (character), double quotes create a `string`. Different types, different purposes.

---

## Zero Values

This is one of Go's most important design decisions. Every type has a default "zero value" when declared without initialization.

```go
var i int           // 0
var f float64       // 0.0
var b bool          // false
var s string        // "" (empty string, not null/undefined)
var p *int          // nil (pointers, we'll cover in Chapter 08)
var sl []int        // nil (slices, Chapter 10)
var m map[string]int // nil (maps, Chapter 10)
```

**Why this matters:** In JavaScript, you constantly check for `null`, `undefined`, `NaN`, and wonder which one you'll get. In Go, every declared variable has a predictable, usable starting value. No `null` for basic types. No `undefined` ever.

```javascript
// JavaScript: the "is it really there?" dance
function greet(name) {
    if (name === undefined || name === null || name === '') {
        name = 'World'
    }
    // Or: name = name || 'World'  (but breaks on name = 0 or false)
    // Or: name = name ?? 'World'  (nullish coalescing)
    // Or: name ??= 'World'
    console.log(`Hello, ${name}`)
}
```

```go
// Go: zero value is always predictable
func greet(name string) {
    if name == "" {
        name = "World"
    }
    fmt.Printf("Hello, %s\n", name)
}
```

One check. One condition. No `undefined` vs `null` vs `NaN` confusion.

### Make the Zero Value Useful

This is a Go proverb in action. Well-designed Go types work at their zero value:

```go
// bytes.Buffer — zero value is an empty, ready-to-use buffer
var buf bytes.Buffer          // no constructor needed
buf.WriteString("hello")     // just works
buf.WriteString(" world")
fmt.Println(buf.String())    // "hello world"

// sync.Mutex — zero value is an unlocked mutex
var mu sync.Mutex             // no constructor needed
mu.Lock()                     // just works
mu.Unlock()

// Compare to JS: new Map(), new Set(), new Buffer() — always need constructors
```

When you design your own types later, aim for useful zero values too.

---

## Type Conversions (Not Coercion)

Go has **no implicit type conversion**. Ever. You must be explicit.

```go
var i int = 42
var f float64 = float64(i)     // explicit conversion: int → float64
var u uint = uint(f)           // explicit conversion: float64 → uint

// This does NOT work:
// var f float64 = i            // ❌ compile error: cannot use i (int) as float64
// var sum = i + f              // ❌ compile error: mismatched types int and float64
```

Compare to JavaScript's chaos:

```javascript
// JavaScript implicit coercion — a hall of mirrors
"5" + 3        // "53"  (string)
"5" - 3        // 2     (number)
"5" * "3"      // 15    (number)
true + true    // 2     (number)
[] + []        // ""    (string)
[] + {}        // "[object Object]" (string)
{} + []        // 0     (number)  ...what?
```

```go
// Go: you get exactly what you ask for
// "5" + 3      // ❌ compile error: mismatched types
// true + true  // ❌ compile error: operator + not defined on bool
```

No surprises. No WAT talks. The compiler catches type mismatches before your code runs.

### String Conversions

```go
import "strconv"

// int → string
s := strconv.Itoa(42)                  // "42"

// string → int
n, err := strconv.Atoi("42")          // n=42, err=nil
n, err = strconv.Atoi("not a number") // n=0, err=error

// float → string
s = strconv.FormatFloat(3.14, 'f', 2, 64)  // "3.14"

// string → float
f, err := strconv.ParseFloat("3.14", 64)   // f=3.14, err=nil
```

Notice every parse function returns `(value, error)`. There's no `parseInt("abc")` returning `NaN` silently. If parsing fails, you get an error. You handle it.

**Important:** `string(65)` does NOT give you `"65"`. It gives you `"A"` (the UTF-8 character at code point 65). This catches many beginners.

```go
fmt.Println(string(65))        // "A"  — converts int to rune to string
fmt.Println(strconv.Itoa(65))  // "65" — converts int to its string representation
```

---

## Constants

```go
const Pi = 3.14159265358979
const MaxRetries = 3
const AppName = "myapp"

// Grouped constants
const (
    StatusPending  = 0
    StatusActive   = 1
    StatusInactive = 2
)
```

Constants are compile-time values. They cannot be changed, cannot be computed at runtime, and cannot be assigned from function calls.

```go
const x = computeSomething()  // ❌ compile error — must be compile-time value
```

### iota — Go's Enum-Like Tool

Go doesn't have enums. Instead, it has `iota` — a constant generator that auto-increments.

```go
type Weekday int

const (
    Sunday Weekday = iota    // 0
    Monday                   // 1
    Tuesday                  // 2
    Wednesday                // 3
    Thursday                 // 4
    Friday                   // 5
    Saturday                 // 6
)

fmt.Println(Wednesday)       // 3
```

`iota` resets to 0 in each `const` block and increments by 1 for each line.

More advanced usage — bitmask flags:

```go
type Permission uint8

const (
    Read    Permission = 1 << iota  // 1  (binary: 001)
    Write                           // 2  (binary: 010)
    Execute                         // 4  (binary: 100)
)

// Combine with bitwise OR
userPerm := Read | Write            // 3  (binary: 011)

// Check with bitwise AND
canRead := userPerm&Read != 0       // true
canExec := userPerm&Execute != 0    // false
```

Compare to TypeScript enums:

```typescript
// TypeScript
enum Weekday {
    Sunday,    // 0
    Monday,    // 1
    Tuesday,   // 2
}

// Go equivalent is const + iota — simpler, no extra language feature
```

### Untyped Constants

A unique Go feature: constants can be "untyped" — they have a kind but no fixed type yet.

```go
const x = 42          // untyped integer constant — not int, not int32, just "42"
const y = 3.14        // untyped float constant

var i int = x         // ✅ x fits in int
var f float64 = x     // ✅ x fits in float64
var b byte = x        // ✅ x fits in byte (42 < 255)
var small int8 = 200  // ❌ 200 overflows int8 (-128 to 127)
```

Untyped constants adapt to whatever type context they're used in, as long as the value fits. This is why you can write `time.Sleep(5 * time.Second)` — the `5` adapts to `time.Duration` automatically.

---

## Multiple Variable Declaration

Go lets you declare multiple variables in one statement:

```go
// Multiple assignment
x, y := 10, 20

// Swap without temp variable
x, y = y, x

// Multiple return values (extremely common in Go)
value, err := strconv.Atoi("42")
if err != nil {
    // handle error
}

// Ignore a return value with blank identifier _
value, _ := strconv.Atoi("42")  // deliberately ignore error (use sparingly!)
```

The blank identifier `_` is Go's way of saying "I know this returns two values, I intentionally don't need the second one." This is important because **unused variables are compile errors** in Go:

```go
x := 42
// If you never use x, the program won't compile.
// This is not a warning. It's an error.
```

Coming from JS where unused variables are lint warnings you often ignore — Go enforces this at the compiler level.

---

## Type Aliases and Defined Types

```go
// Type alias — same type, just a different name
type Text = string       // Text IS string, completely interchangeable

// Defined type — new type based on existing one
type UserID int64        // UserID is a NEW type, not interchangeable with int64
type Celsius float64
type Fahrenheit float64
```

The difference matters:

```go
type Celsius float64
type Fahrenheit float64

var temp Celsius = 100
// var f Fahrenheit = temp  // ❌ compile error: cannot use Celsius as Fahrenheit

// Must explicitly convert:
var f Fahrenheit = Fahrenheit(temp * 9/5 + 32)  // ✅
```

This is powerful for domain modeling. The compiler prevents you from accidentally mixing up Celsius and Fahrenheit, user IDs and order IDs, etc. TypeScript can do similar things with branded types, but it's a workaround. In Go, it's a first-class feature.

---

## fmt Package: Printing and Formatting

You'll use `fmt` constantly. Here are the essentials:

```go
name := "Go"
age := 19
pi := 3.14159

// Basic printing
fmt.Println("Hello", name)              // Hello Go (adds newline, space-separated)
fmt.Print("no newline")                 // no newline (no newline added)

// Formatted printing
fmt.Printf("Name: %s, Age: %d\n", name, age)   // Name: Go, Age: 19
fmt.Printf("Pi: %.2f\n", pi)                    // Pi: 3.14
fmt.Printf("Type: %T\n", age)                   // Type: int
fmt.Printf("Value: %v\n", age)                  // Value: 19 (default format)
fmt.Printf("Binary: %b\n", age)                 // Binary: 10011
fmt.Printf("Quoted: %q\n", name)                // Quoted: "Go"

// Format to string (like template literals but type-safe)
msg := fmt.Sprintf("Hello %s, you are %d", name, age)
```

Common format verbs:

| Verb  | Meaning                 | Example                                        |
| ----- | ----------------------- | ---------------------------------------------- |
| `%v`  | Default format          | `fmt.Printf("%v", 42)` → `42`                  |
| `%T`  | Type of value           | `fmt.Printf("%T", 42)` → `int`                 |
| `%d`  | Integer (decimal)       | `fmt.Printf("%d", 42)` → `42`                  |
| `%s`  | String                  | `fmt.Printf("%s", "hi")` → `hi`                |
| `%f`  | Float                   | `fmt.Printf("%.2f", 3.14)` → `3.14`            |
| `%t`  | Boolean                 | `fmt.Printf("%t", true)` → `true`              |
| `%q`  | Quoted string           | `fmt.Printf("%q", "hi")` → `"hi"`              |
| `%p`  | Pointer address         | `fmt.Printf("%p", &x)` → `0xc000...`           |
| `%+v` | Struct with field names | `fmt.Printf("%+v", user)` → `{Name:Go Age:19}` |

---

## Common Mistakes

### Mistake 1: Using `:=` when you mean `=`

```go
x := 10
// Later...
x := 20      // ❌ compile error: "no new variables on left side of :="
x = 20       // ✅ assignment to existing variable
```

`:=` declares a **new** variable. `=` assigns to an **existing** one. In JS, you'd just use `=` after `let`/`const`. In Go, declaration and assignment are distinct operations.

Exception: `:=` works if at least one variable on the left is new:

```go
x := 10
x, y := 20, 30   // ✅ — y is new, x is reassigned
```

### Mistake 2: Expecting JS-style string concatenation with non-strings

```go
// age := 25
// msg := "Age: " + age      // ❌ compile error: mismatched types string and int
msg := "Age: " + strconv.Itoa(age)    // ✅
msg := fmt.Sprintf("Age: %d", age)    // ✅ (preferred)
```

### Mistake 3: Assuming `len()` counts characters

```go
s := "café"
fmt.Println(len(s))                        // 5 (bytes — é is 2 bytes in UTF-8)
fmt.Println(utf8.RuneCountInString(s))     // 4 (characters)
```

### Mistake 4: Thinking `string(number)` works like `String(number)` in JS

```go
fmt.Println(string(72))           // "H" — interprets 72 as Unicode code point
fmt.Println(strconv.Itoa(72))     // "72" — what you probably wanted
```

---

## Exercises

### Exercise 1: Zero Value Explorer

Write a program that declares variables of every basic type without initialization, then prints their zero values and types:

```go
package main

import "fmt"

func main() {
    var i int
    var f float64
    var b bool
    var s string
    var r rune
    var by byte
    
    // Print each variable's zero value and type using %v and %T
    // Expected output:
    // int:     0       (type: int)
    // float64: 0       (type: float64)
    // bool:    false   (type: bool)
    // string:  ""      (type: string)
    // rune:    0       (type: int32)
    // byte:    0       (type: uint8)
}
```

### Exercise 2: Temperature Converter

Create defined types `Celsius` and `Fahrenheit`. Write functions to convert between them. The compiler should prevent mixing the two types accidentally.

```go
type Celsius float64
type Fahrenheit float64

func CtoF(c Celsius) Fahrenheit {
    // implement: F = C × 9/5 + 32
}

func FtoC(f Fahrenheit) Celsius {
    // implement: C = (F - 32) × 5/9
}

// Verify the compiler catches this mistake:
// var temp Celsius = Fahrenheit(100)  // should not compile
```

### Exercise 3: String Deep Dive

Write a program that takes the string `"Hello, 世界! 🌍"` and prints:

1. Length in bytes (`len`)
2. Length in characters/runes (`utf8.RuneCountInString`)
3. Each rune with its byte index, character, and Unicode code point

Expected to see that the emoji `🌍` takes 4 bytes in UTF-8.

---

## Key Takeaways

1. **`:=` inside functions, `var` at package level.** Use `:=` for 90% of local variables. Use `var` when you want the zero value or need package-level scope.

2. **Zero values eliminate null/undefined.** Every variable starts with a predictable, usable default. Design your own types to have useful zero values.

3. **No implicit coercion, ever.** `"5" + 3` is a compile error. Type conversions are always explicit. This catches bugs at compile time, not in production at 3am.

4. **`int` for numbers, `string` for text, `bool` for logic.** Keep it simple. Use sized types (`int32`, `uint8`) only when the domain requires it (binary protocols, interop).

5. **`iota` replaces enums.** Combined with defined types, it's a clean pattern for constants with type safety.

6. **Unused variables are compile errors.** Go won't let you leave dead code around. Use `_` to deliberately ignore values.

---

## 🧭 Navigation

| Direction    | Link                                                       |
| ------------ | ---------------------------------------------------------- |
| **Previous** | [← Chapter 02: Setup & Tooling](./02-setup-and-tooling.md) |
| **Next**     | [Chapter 04: Control Flow →](./04-control-flow.md)         |
