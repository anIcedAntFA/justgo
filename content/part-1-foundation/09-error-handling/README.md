# Chapter 09: Error Handling

> **Errors as values, error wrapping, errors.Is/As, custom error types, and sentinel errors.**

## TL;DR

Go has no exceptions, no try/catch. Errors are ordinary values returned from functions, and you handle them explicitly with `if err != nil`. This feels verbose coming from JavaScript, but it's Go's defining philosophy: error paths are always visible, never hidden. Master error wrapping (`%w`), `errors.Is`, and `errors.As` — these are the modern tools that make Go error handling powerful, not just verbose.

---

## Errors Are Values

In Go, an error is just a value that implements the `error` interface:

```go
type error interface {
    Error() string
}
```

Functions that can fail return an `error` as their last return value. `nil` means success; non-nil means failure.

```go
func readConfig(path string) ([]byte, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err      // propagate the error
    }
    return data, nil          // success — error is nil
}

// Caller handles the error
data, err := readConfig("config.yaml")
if err != nil {
    log.Fatal(err)            // handle the failure
}
// use data — we only reach here on success
```

This `value, err := ...; if err != nil { }` pattern is the rhythm of Go. You'll write it thousands of times. It becomes muscle memory.

> Runnable demo: [`examples/values`](./examples/values/main.go).

---

## Why No Exceptions?

This is the biggest philosophical difference from JavaScript. Let's understand the reasoning.

### The JavaScript Way

```javascript
// JavaScript: errors are exceptional, thrown and caught
async function processOrder(id) {
    const order = await fetchOrder(id)      // might throw
    const payment = await chargeCard(order) // might throw
    const receipt = await sendEmail(payment) // might throw
    return receipt
}

try {
    await processOrder(123)
} catch (err) {
    // Which line threw? What kind of error? Hard to tell.
    // The happy path looks clean, but error handling is disconnected.
    console.error(err)
}
```

The problems:

- **Invisible control flow.** Any line could throw. You can't tell by reading which operations might fail.
- **Easy to forget.** Forget a try/catch and an exception crashes your app or silently propagates.
- **Disconnected handling.** The error handling is far from where the error occurred.

### The Go Way

```go
// Go: errors are values, handled where they occur
func processOrder(id int) (*Receipt, error) {
    order, err := fetchOrder(id)
    if err != nil {
        return nil, fmt.Errorf("fetching order %d: %w", id, err)
    }

    payment, err := chargeCard(order)
    if err != nil {
        return nil, fmt.Errorf("charging card for order %d: %w", id, err)
    }

    receipt, err := sendEmail(payment)
    if err != nil {
        return nil, fmt.Errorf("sending receipt for order %d: %w", id, err)
    }

    return receipt, nil
}
```

The benefits:

- **Visible control flow.** Every operation that can fail is marked with explicit error handling. You see exactly where things can go wrong.
- **Can't forget.** The compiler won't let you ignore a return value silently (and the linter catches unused errors).
- **Local handling.** You decide what to do at each failure point — wrap with context, retry, return, log.

The cost is verbosity. The benefit is reliability. Go's designers made the deliberate tradeoff: **explicit and a bit tedious beats implicit and surprising.**

---

## Creating Errors

### errors.New — Simple Static Errors

```go
import "errors"

func validate(age int) error {
    if age < 0 {
        return errors.New("age cannot be negative")
    }
    return nil
}
```

Use `errors.New` for simple, static error messages with no dynamic content.

### fmt.Errorf — Errors with Context

```go
import "fmt"

func validate(age int) error {
    if age < 0 {
        return fmt.Errorf("invalid age: %d (must be non-negative)", age)
    }
    return nil
}
```

Use `fmt.Errorf` when you need to include dynamic values in the message.

### Error Message Conventions

Go has strict conventions for error messages (the linter enforces some):

```go
// ✅ Lowercase, no trailing punctuation
errors.New("connection failed")
fmt.Errorf("cannot open file %s", path)

// ❌ Capitalized
errors.New("Connection failed")

// ❌ Trailing punctuation
errors.New("connection failed.")
```

Why lowercase and no period? Because errors are often **wrapped** — combined into a chain like `"initApp failed: loadConfig: connection failed"`. Capitalized words and periods in the middle of that chain look wrong.

---

## Error Wrapping with %w

This is the modern (Go 1.13+) way to add context while preserving the original error. The `%w` verb in `fmt.Errorf` "wraps" an error.

```go
func loadConfig(path string) ([]byte, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        // %w wraps the original error, adding context
        return nil, fmt.Errorf("loadConfig(%s): %w", path, err)
    }
    return data, nil
}

func initApp() error {
    _, err := loadConfig("/etc/app/config.yaml")
    if err != nil {
        return fmt.Errorf("initApp: %w", err)   // wrap again
    }
    return nil
}
```

When the error propagates up, you get a full chain of context:

```
initApp: loadConfig(/etc/app/config.yaml): open /etc/app/config.yaml: no such file or directory
```

Each layer added its context, but the original error (`no such file or directory`) is preserved and still inspectable.

### %w vs %v

```go
// %w — WRAPS the error (preserves it for errors.Is/As inspection)
fmt.Errorf("context: %w", err)

// %v — formats the error as a STRING (loses the ability to inspect)
fmt.Errorf("context: %v", err)
```

Use `%w` when you want callers to be able to inspect the underlying error. Use `%v` when you just want the message text and don't need inspection (or when you deliberately want to "hide" the underlying error type).

### JS Comparison

```javascript
// JavaScript: error chaining with `cause` (ES2022)
try {
    await loadConfig()
} catch (err) {
    throw new Error("initApp failed", { cause: err })  // similar to %w
}
```

```go
// Go: %w wrapping
if err != nil {
    return fmt.Errorf("initApp failed: %w", err)
}
```

JS's `{ cause: err }` (added in ES2022) is conceptually similar to Go's `%w` — both preserve the original error in a chain. Go has had this pattern as a core idiom for longer.

> Runnable demo: [`examples/wrapping`](./examples/wrapping/main.go).

---

## Inspecting Errors: errors.Is and errors.As

Once errors are wrapped, you need ways to inspect them through the wrapping. Two functions:

### errors.Is — Check for a Specific Error Value

`errors.Is` checks whether an error (or anything it wraps) matches a specific **sentinel error**.

```go
import (
    "errors"
    "os"
)

_, err := os.ReadFile("missing.txt")

// Direct comparison FAILS through wrapping:
// if err == os.ErrNotExist { }   // ❌ won't match if err is wrapped

// errors.Is checks the whole chain:
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("file does not exist")    // ✅ works even through wrapping
}
```

`errors.Is` unwraps the error chain and checks each level. Even if the error was wrapped five layers deep, `errors.Is` finds the sentinel.

### errors.As — Extract a Specific Error Type

`errors.As` checks whether an error (or anything it wraps) is of a specific **type**, and if so, extracts it so you can access its fields.

```go
type ValidationError struct {
    Field string
    Value any
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on %s", e.Field)
}

// Somewhere an error is created and wrapped:
err := fmt.Errorf("processing failed: %w", &ValidationError{Field: "email", Value: "bad"})

// Extract the ValidationError from the chain:
var valErr *ValidationError
if errors.As(err, &valErr) {
    // valErr is now the extracted *ValidationError — access its fields
    fmt.Printf("field %q had bad value: %v\n", valErr.Field, valErr.Value)
}
```

The difference:

- **`errors.Is`** — "Is this error (or something it wraps) equal to THIS specific value?" → for sentinel errors
- **`errors.As`** — "Is this error (or something it wraps) of THIS type? If so, give it to me." → for custom error types with fields

### errors.AsType — the Go 1.26 Generic Form of As

`errors.As` is showing its age: you must declare a target variable, pass its address, and it takes an `any` — so a wrong target type is a **runtime panic**, not a compile error. Go 1.26 adds a generic version, `errors.AsType[E error]`, that fixes all three:

```go
// Signature: func AsType[E error](err error) (E, bool)

// errors.As — the pre-1.26 form: declare, pass a pointer, check the bool
var valErr *ValidationError
if errors.As(err, &valErr) {
    fmt.Println(valErr.Field)
}

// errors.AsType — Go 1.26: the type is a parameter, the value is returned
if valErr, ok := errors.AsType[*ValidationError](err); ok {
    fmt.Println(valErr.Field)
}
```

Both walk the same wrapped tree and find the first match for the type. `AsType` is **type-safe** (the type is a compile-time parameter, so there's no `any` target to get wrong), **faster**, and reads like the comma-ok idiom you already know from map lookups and type assertions. Prefer `errors.AsType` on Go 1.26+; you'll still see `errors.As` everywhere in existing code and older modules, so know both. (There's no `IsType` — `errors.Is` already takes a typed `error` target, so it was never `any`-unsafe the way `As` was.)

---

## Sentinel Errors

A sentinel error is a predefined, exported error value that callers can check for. The convention is to name them `ErrXxx`.

```go
package store

import "errors"

// Exported sentinel errors — callers can check for these
var (
    ErrNotFound     = errors.New("item not found")
    ErrUnauthorized = errors.New("unauthorized")
    ErrConflict     = errors.New("conflict: item already exists")
)

func (s *Store) Get(id string) (*Item, error) {
    item, ok := s.items[id]
    if !ok {
        return nil, ErrNotFound       // return the sentinel
    }
    return item, nil
}
```

Callers check with `errors.Is`:

```go
item, err := store.Get("123")
if errors.Is(err, store.ErrNotFound) {
    // handle "not found" specifically — maybe return 404
    http.Error(w, "not found", http.StatusNotFound)
    return
}
if err != nil {
    // some other error
    http.Error(w, "internal error", http.StatusInternalServerError)
    return
}
```

You've seen sentinel errors already: `os.ErrNotExist`, `io.EOF`, `sql.ErrNoRows`. These are all sentinel errors the standard library exports for you to check against.

```go
// Common standard library sentinels:
io.EOF                    // end of input stream
sql.ErrNoRows             // database query returned no rows
os.ErrNotExist            // file doesn't exist
context.Canceled          // context was canceled
context.DeadlineExceeded  // context timed out
errors.ErrUnsupported     // operation not supported (Go 1.21+) — the stdlib
                          // returns it (e.g. os.Link on some filesystems) so you
                          // can errors.Is against a single well-known sentinel
```

> Runnable demo: [`examples/sentinel`](./examples/sentinel/main.go).

---

## Custom Error Types

When you need to carry structured data with an error (not just a message), define a custom error type.

```go
type HTTPError struct {
    StatusCode int
    URL        string
    Err        error    // the underlying error (for wrapping)
}

// Implement the error interface
func (e *HTTPError) Error() string {
    return fmt.Sprintf("HTTP %d from %s: %v", e.StatusCode, e.URL, e.Err)
}

// Implement Unwrap so errors.Is/As can see the wrapped Err
func (e *HTTPError) Unwrap() error {
    return e.Err
}
```

The `Unwrap()` method is important — it lets `errors.Is` and `errors.As` traverse into your custom error's wrapped error. Implement it whenever your custom error wraps another.

Usage:

```go
func fetch(url string) error {
    resp, err := http.Get(url)
    if err != nil {
        return &HTTPError{StatusCode: 0, URL: url, Err: err}
    }
    if resp.StatusCode >= 400 {
        return &HTTPError{
            StatusCode: resp.StatusCode,
            URL:        url,
            Err:        fmt.Errorf("bad status"),
        }
    }
    return nil
}

// Caller extracts the structured error
err := fetch("https://api.example.com/data")
var httpErr *HTTPError
if errors.As(err, &httpErr) {
    if httpErr.StatusCode == 429 {
        // rate limited — back off and retry
    }
}
```

### Sentinel vs Custom Type — When to Use Which

| Use a **sentinel error** when          | Use a **custom error type** when           |
| -------------------------------------- | ------------------------------------------ |
| The error is a simple, known condition | The error carries structured data (fields) |
| Callers just need to identify it       | Callers need to extract details            |
| `ErrNotFound`, `ErrTimeout`            | `ValidationError{Field, Rule}`             |
| Check with `errors.Is`                 | Check with `errors.As`                     |

> Runnable demo: [`examples/custom-type`](./examples/custom-type/main.go).

---

## Combining Errors — errors.Join

Sometimes an operation produces **several** errors at once — validating every field of a form, closing multiple resources, running a batch. Since Go 1.20, `errors.Join` bundles them into a single `error` whose chain `errors.Is`/`errors.As` can still walk.

```go
func validate(u User) error {
    var errs []error
    if u.Name == "" {
        errs = append(errs, ErrEmptyName)
    }
    if u.Age < 0 {
        errs = append(errs, ErrNegativeAge)
    }
    return errors.Join(errs...)   // nil if errs is empty — no error
}

err := validate(User{Name: "", Age: -1})
fmt.Println(err)                    // each joined error on its own line
errors.Is(err, ErrEmptyName)        // true
errors.Is(err, ErrNegativeAge)      // true
```

`errors.Join` skips `nil` arguments and returns `nil` if every argument is `nil`, so you can collect as you go without special-casing "no errors." The same Go 1.20 release also lets `fmt.Errorf` take **multiple** `%w` verbs (`fmt.Errorf("%w and %w", e1, e2)`) — both build a multi-error chain.

### The multi-error Unwrap: `Unwrap() []error`

How does `errors.Is` find _both_ joined errors when a chain is a tree, not a line? Go 1.20 taught the unwrap machinery a second shape. A single-wrap error implements `Unwrap() error`; a **multi**-wrap error implements `Unwrap() []error`. `errors.Is`/`As`/`AsType` try both, so they traverse the whole tree. `errors.Join` returns exactly such a value — and your own custom error can join, too, by implementing the slice form:

```go
type ValidationErrors struct {
    Errs []error
}

func (e *ValidationErrors) Error() string { /* join the messages */ return "..." }

// Return every wrapped error, so errors.Is/As/AsType can see them all.
func (e *ValidationErrors) Unwrap() []error { return e.Errs }
```

Note there's no `%w`-style ordering here: a `[]error` unwrap has no single "next" error, so `errors.Unwrap` (the singular function) returns `nil` for a multi-error — use `errors.Is`/`As`/`AsType` to inspect it, never a manual unwrap loop.

> Runnable demo: [`examples/join`](./examples/join/main.go).

---

## Panic and Recover (Use Rarely)

Go does have a mechanism that resembles exceptions: `panic` and `recover`. But they are **not** for normal error handling.

### panic — Unrecoverable Errors

```go
func mustParse(s string) int {
    n, err := strconv.Atoi(s)
    if err != nil {
        panic(fmt.Sprintf("mustParse: invalid number %q", s))
    }
    return n
}
```

`panic` stops normal execution, runs deferred functions, and crashes the program (printing a stack trace). Use it ONLY for:

- **Truly unrecoverable situations** (programmer errors, impossible states)
- **Initialization failures** where the program can't function (e.g., a required config is malformed at startup)

### recover — Catching Panics

`recover` stops a panic from crashing the program. It only works inside a deferred function.

```go
func safeRun(fn func()) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered from panic: %v", r)
        }
    }()
    fn()
    return nil
}
```

The most common legitimate use of `recover` is in long-running servers — you don't want one request's panic to crash the entire server. HTTP frameworks include recovery middleware for exactly this (we'll build one in Chapter 16).

### The Golden Rule

```
Normal errors → return them as values (error)
Catastrophic, unrecoverable, programmer-error situations → panic (rarely)
Server resilience (don't crash on one bad request) → recover in middleware
```

**Do not use panic/recover as a try/catch substitute.** That's the #1 mistake JS developers make in Go. If you find yourself using panic for expected error conditions (bad user input, missing files, failed network calls), you're fighting the language. Return errors instead.

```javascript
// JavaScript habit (DON'T translate this to Go):
function getUser(id) {
    if (!id) throw new Error("id required")   // throwing for expected condition
}
```

```go
// Go way — return the error, don't panic
func getUser(id string) (*User, error) {
    if id == "" {
        return nil, errors.New("id required")   // expected condition → return error
    }
    // ...
}
```

> Runnable demo: [`examples/panic-recover`](./examples/panic-recover/main.go).

---

## Error Handling Patterns in Practice

### Pattern 1: Wrap with Context as Errors Propagate

```go
func (s *Service) CreateUser(ctx context.Context, input UserInput) (*User, error) {
    if err := input.Validate(); err != nil {
        return nil, fmt.Errorf("validating user input: %w", err)
    }

    user, err := s.repo.Insert(ctx, input)
    if err != nil {
        return nil, fmt.Errorf("inserting user: %w", err)
    }

    if err := s.email.SendWelcome(user.Email); err != nil {
        // Maybe this isn't fatal — log it but don't fail the whole operation
        log.Printf("failed to send welcome email: %v", err)
    }

    return user, nil
}
```

Notice the last error is logged but not returned — sometimes a failure is non-fatal. You decide per error.

### Pattern 2: Handle at the Right Level

```go
// Low level: return errors, don't log them
func (r *Repo) GetUser(id int) (*User, error) {
    row := r.db.QueryRow("SELECT ... WHERE id = ?", id)
    var u User
    if err := row.Scan(&u.Name); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrUserNotFound       // translate to domain error
        }
        return nil, fmt.Errorf("scanning user: %w", err)
    }
    return &u, nil
}

// Top level (HTTP handler): decide the response, log once
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
    user, err := h.service.GetUser(id)
    if errors.Is(err, ErrUserNotFound) {
        http.Error(w, "user not found", http.StatusNotFound)
        return
    }
    if err != nil {
        log.Printf("getting user %d: %v", id, err)   // log ONCE at the top
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(user)
}
```

**Key principle: don't log AND return.** If you return an error, let the caller decide whether to log. Logging at every level produces duplicate log spam. Log once, at the level where you handle the error (usually the top).

---

## Common Mistakes

### Mistake 1: Ignoring errors

```go
// ❌ Silently ignoring — the linter (errcheck) flags this
data, _ := os.ReadFile("config.yaml")

// ✅ Handle it
data, err := os.ReadFile("config.yaml")
if err != nil {
    return fmt.Errorf("reading config: %w", err)
}
```

### Mistake 2: Using == instead of errors.Is

```go
// ❌ Breaks when the error is wrapped
if err == sql.ErrNoRows { }

// ✅ Works through wrapping
if errors.Is(err, sql.ErrNoRows) { }
```

### Mistake 3: Logging AND returning

```go
// ❌ Double handling — produces duplicate logs as it propagates
func doThing() error {
    if err := step(); err != nil {
        log.Println(err)        // logged here
        return err              // ...and will be logged again by caller
    }
}

// ✅ Just return (with context). Log once at the top.
func doThing() error {
    if err := step(); err != nil {
        return fmt.Errorf("doThing: %w", err)
    }
    return nil
}
```

### Mistake 4: Using panic for normal errors

```go
// ❌ Panicking for an expected condition
func parsePort(s string) int {
    n, err := strconv.Atoi(s)
    if err != nil {
        panic("bad port")     // NO — this is an expected failure
    }
    return n
}

// ✅ Return an error
func parsePort(s string) (int, error) {
    n, err := strconv.Atoi(s)
    if err != nil {
        return 0, fmt.Errorf("parsing port %q: %w", s, err)
    }
    return n, nil
}
```

### Mistake 5: Wrapping when you shouldn't

```go
// ⚠️ Sometimes you DON'T want to expose the underlying error
// (e.g., leaking internal DB errors to API responses)
func (s *Service) Login(user, pass string) error {
    err := s.db.checkPassword(user, pass)
    if err != nil {
        // Don't wrap and leak DB internals to the caller — return a generic error
        return ErrInvalidCredentials       // %w would leak the DB error
    }
    return nil
}
```

Wrap to add useful context. Don't wrap when it would leak internal details or when the caller only needs a generic error.

---

## Exercises

Stubs and tests live in [`exercises/`](./exercises/). Remove the `t.Skip` in each `_test.go`, implement the stub, and run `go test ./...` until it passes.

### Exercise 1: Sentinel Errors

In [`exercises/store.go`](./exercises/store.go), build a simple in-memory key-value store with sentinel errors. `Get` returns `ErrKeyNotFound` when the key is missing, `Create` returns `ErrKeyExists` when the key is already present, and `Delete` returns `ErrKeyNotFound` when there's nothing to delete. The test uses `errors.Is` to handle each case distinctly — proving callers can branch on the sentinel rather than string-matching messages.

### Exercise 2: Error Wrapping Chain

In [`exercises/chain.go`](./exercises/chain.go), implement three functions that call each other (`a` → `b` → `c`), where `c` fails with a sentinel. Each layer wraps the error with `%w` and adds its own context. The test checks the full chain appears in the final message **and** that `errors.Is` still detects the original sentinel from the top — the two things `%w` buys you at once.

### Exercise 3: Custom Error Type with errors.As

In [`exercises/parse.go`](./exercises/parse.go), create a `ParseError` custom type with fields `Line`, `Column`, and `Message`. Implement `Error()` and `Unwrap()`. Write a parser function that returns a `ParseError` wrapped in additional context; the test then uses `errors.AsType[*ParseError]` (Go 1.26) to extract it and read its line/column — the structured-data payload a plain sentinel can't carry.

### Exercise 4: Recovery Middleware (Preview of Part 2)

In [`exercises/recover.go`](./exercises/recover.go), implement `safeExecute(fn func() error) error` that runs `fn`, recovers a panic into an error (`"recovered from panic: ..."`), returns a normal error unchanged, and returns `nil` on success. This is the seed of the HTTP recovery middleware you'll build in Chapter 16 — one bad call shouldn't crash the whole server.

---

## Key Takeaways

1. **Errors are values, not exceptions.** Return them, check them with `if err != nil`. Explicit and visible beats implicit and surprising.

2. **Wrap with `%w` to add context** while preserving the original error for inspection. Build a chain of context as errors propagate.

3. **`errors.Is` for sentinel values, `errors.As`/`errors.AsType` for custom types.** Never use `==` on errors that might be wrapped. On Go 1.26+ prefer `errors.AsType[T](err)` — the type-safe generic form — over the older `errors.As(err, &target)`.

4. **Sentinel errors (`ErrXxx`) for known conditions; custom error types when you need structured data.**

5. **`errors.Join` (Go 1.20+) bundles multiple errors** into one inspectable value — and `fmt.Errorf` now takes multiple `%w` verbs. Custom types join by implementing `Unwrap() []error`; `errors.Is`/`As`/`AsType` walk the whole tree.

6. **panic/recover is NOT try/catch.** Use panic only for unrecoverable situations. Use recover only for server resilience (don't crash on one bad request). Return errors for everything expected.

7. **Log once, at the top.** Don't log and return — that creates duplicate logs. Return errors with context; let the top level decide how to handle and log them.

---

## 🧭 Navigation

| Direction    | Link                                                     |
| ------------ | -------------------------------------------------------- |
| **Previous** | [← Chapter 08: Pointers](../08-pointers/README.md)       |
| **Next**     | [Chapter 10: Collections →](../10-collections/README.md) |
