# Chapter 11: Packages & Modules

> **Package organization, visibility, go.mod, dependency management, and internal packages.**

## TL;DR

A package is Go's unit of code organization — every `.go` file belongs to exactly one package, and packages map to directories. A module is a collection of packages versioned together (defined by `go.mod`). The two rules that surprise JS developers most: visibility is controlled by capitalization (not `export`), and there are no circular imports allowed — ever.

---

## Packages: The Unit of Organization

Every Go file starts with a `package` declaration. All files in the same directory must declare the same package.

```go
// file: math/calculator.go
package math

func Add(a, b int) int { return a + b }
```

```go
// file: math/geometry.go — same directory, same package
package math

func Area(w, h int) int { return w * h }
```

Both files are part of package `math`. They can use each other's functions directly (no import needed — they're the same package).

### The package main Special Case

```go
// An executable program must have package main + func main()
package main

import "fmt"

func main() {
    fmt.Println("This is an executable")
}
```

`package main` with a `func main()` is the entry point of an executable. Every other package is a **library** — importable code that can't run on its own.

```
package main  → produces an executable binary (go build creates ./yourapp)
package xxx   → a library, imported by other packages
```

### JS Comparison

```javascript
// JavaScript: one file = one module, explicit exports
// file: calculator.js
export function add(a, b) { return a + b }
export function subtract(a, b) { return a - b }

// Importing
import { add, subtract } from './calculator.js'
```

```go
// Go: one DIRECTORY = one package, capitalization controls export
// file: math/calculator.go
package math

func Add(a, b int) int { return a + b }       // exported (capital A)
func subtract(a, b int) int { return a - b }  // unexported (lowercase s)

// Importing — import the package, use the qualified name
import "github.com/you/proj/math"
math.Add(1, 2)
```

Key differences:

- JS: module = file. Go: package = directory (can be many files).
- JS: explicit `export` keyword. Go: capitalization decides.
- JS: import specific names. Go: import the package, access members via `package.Name`.

---

## Visibility: Capitalization Is the Access Modifier

This is unique to Go. There's no `public`, `private`, `export`, or `default`. The first letter of an identifier's name controls its visibility:

```go
package store

// Exported (PUBLIC) — capitalized, accessible from other packages
type User struct {
    Name  string      // exported field
    Email string      // exported field
    token string      // unexported field — private to this package
}

func NewUser(name string) *User { }   // exported function
func validate(u *User) error { }       // unexported function

const MaxUsers = 1000                   // exported constant
var defaultTimeout = 30                 // unexported variable
```

The rule applies to **everything**: functions, types, struct fields, methods, constants, variables.

```go
// From another package:
import "github.com/you/proj/store"

u := store.NewUser("Alice")    // ✅ NewUser is exported
fmt.Println(u.Name)            // ✅ Name is exported
// fmt.Println(u.token)        // ❌ token is unexported — compile error
// store.validate(u)           // ❌ validate is unexported — compile error
```

### Why Capitalization?

It's a deliberate design choice that makes visibility **immediately visible** at the use site. When you read `store.User`, you instantly know it's exported. When you read `store.validate`, you know it's internal. No need to look up the definition or scan for `export` keywords. The information is in the name itself.

This takes a day to get used to. Then it feels natural — you can tell a type's visibility just by looking at how it's written anywhere in the code.

---

## Importing Packages

```go
import "fmt"                              // standard library
import "github.com/labstack/echo/v4"      // external package (a URL)
import "github.com/you/proj/internal/db"  // your own package

// Grouped imports (the common style)
import (
    "fmt"
    "net/http"
    "os"

    "github.com/labstack/echo/v4"
    "github.com/you/proj/internal/db"
)
```

`gofmt`/`goimports` automatically groups imports: standard library first, then a blank line, then third-party and local. You don't manage this manually — the tooling does.

### Using Imported Packages

You reference imported members with the package name (the last path segment, usually):

```go
import "net/http"

http.Get("https://example.com")   // package is "http" (last segment of net/http)
http.StatusOK                      // a constant from the package
```

### Import Aliases

When you have naming conflicts or want clarity:

```go
import (
    "math/rand"
    crand "crypto/rand"     // alias to avoid conflict with math/rand
)

rand.Intn(10)              // math/rand
crand.Read(buf)            // crypto/rand
```

### The Blank Import (Side Effects)

Sometimes you import a package only for its `init()` side effects, not to use its exported names:

```go
import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"   // blank import — registers the driver via init()
)

// You never call go-sqlite3 directly, but importing it registers
// the "sqlite3" driver with database/sql.
db, _ := sql.Open("sqlite3", "data.db")
```

The `_` blank import runs the package's `init()` but doesn't bind its name. Common for database drivers, image format decoders, and similar plugin-style registration. We'll use this in Part 2 for databases.

### Unused Imports Are Errors

```go
import (
    "fmt"
    "os"      // ❌ if you never use os, this is a COMPILE ERROR
)

func main() {
    fmt.Println("hello")
    // os is never used → won't compile
}
```

Like unused variables, unused imports are compile errors, not warnings. `goimports` automatically removes them on save, so in practice you rarely hit this — but know why it happens.

---

## Modules: Versioned Collections of Packages

A module is a collection of related packages released together, versioned as a unit. It's defined by a `go.mod` file at its root.

```
myproject/                    ← module root (has go.mod)
├── go.mod                     ← module github.com/you/myproject
├── go.sum
├── main.go                    ← package main
├── store/                     ← package store
│   ├── user.go
│   └── product.go
├── api/                       ← package api
│   ├── handler.go
│   └── middleware.go
└── internal/                  ← internal packages (special, see below)
    └── db/
        └── postgres.go
```

```
module = the whole project (one go.mod)
package = one directory within it
```

### go.mod Anatomy

```
// go.mod
module github.com/you/myproject     // the module path (import prefix for all packages)

go 1.26                              // minimum Go version

require (
    github.com/labstack/echo/v4 v4.13.3    // direct dependency
    github.com/jmoiron/sqlx v1.4.0
)

require (
    golang.org/x/crypto v0.31.0 // indirect    ← transitive dependency
)
```

The `module` line defines the import prefix. If your module is `github.com/you/myproject`, then the `store` package is imported as `github.com/you/myproject/store`.

### Importing Your Own Packages

```go
// In main.go (module github.com/you/myproject)
package main

import (
    "github.com/you/myproject/store"      // import your own store package
    "github.com/you/myproject/api"        // import your own api package
)

func main() {
    u := store.NewUser("Alice")
    api.StartServer()
}
```

Your own packages use the full module path. This is different from JS relative imports (`./store`). Go uses absolute, module-rooted paths everywhere. (gopls and goimports autocomplete these for you.)

---

## Dependency Management

### Adding a Dependency

```bash
# Add and download a dependency
go get github.com/labstack/echo/v4

# Specific version
go get github.com/labstack/echo/v4@v4.13.3

# Latest
go get github.com/labstack/echo/v4@latest
```

This updates `go.mod` and `go.sum`, and downloads the package to your module cache (`$GOPATH/pkg/mod/`, shared across all projects).

### go mod tidy — The Cleanup Command

```bash
go mod tidy
```

`go mod tidy` is essential. It:

- Adds any imports your code uses but `go.mod` is missing
- Removes any dependencies in `go.mod` your code no longer uses
- Updates `go.sum`

Run it whenever you add/remove imports. Think of it as `npm install` + `npm prune` combined, driven by your actual code rather than manual `package.json` edits.

### go.sum — Integrity Verification

```
github.com/labstack/echo/v4 v4.13.3 h1:pSc...
github.com/labstack/echo/v4 v4.13.3/go.mod h1:gB...
```

`go.sum` contains cryptographic hashes of every dependency (and their `go.mod` files). When you build, Go verifies that downloaded packages match these hashes. This prevents tampering — a dependency can't be silently swapped. It's like `package-lock.json` but with actual cryptographic verification, not just version pinning.

**Commit both `go.mod` and `go.sum`** to version control. Never edit `go.sum` by hand.

### Versioning and Major Versions

Go uses semantic versioning, but with a twist for major versions ≥ 2:

```go
import "github.com/labstack/echo/v4"   // note the /v4 — major version in the path!
```

Major versions 2+ include the version in the import path (`/v4`). This is "Semantic Import Versioning" — it lets you import two major versions of the same library simultaneously if needed. Unusual, but it solves real dependency conflicts.

---

## internal Packages — Enforced Privacy

Go has a special directory name: `internal`. Packages inside an `internal/` directory can only be imported by code rooted at `internal`'s parent.

```
myproject/
├── go.mod
├── api/
│   └── handler.go          ← CAN import internal/db
├── internal/
│   └── db/
│       └── postgres.go     ← only importable within myproject
└── ...
```

```go
// api/handler.go — within myproject, CAN import internal
import "github.com/you/myproject/internal/db"   // ✅ allowed

// A DIFFERENT module trying to import your internal package:
import "github.com/you/myproject/internal/db"   // ❌ compile error — internal
```

This is compiler-enforced encapsulation at the module level. Anything in `internal/` is private to your module — external code (and other modules) physically cannot import it. Use it for code you don't want to expose as public API: database layers, internal utilities, implementation details.

This is more powerful than capitalization (which is per-package). `internal/` controls visibility across your whole module boundary.

---

## No Circular Imports

Go absolutely forbids circular imports. If package A imports package B, then B cannot import A — directly or transitively.

```
package A imports package B   ✅
package B imports package A   ❌ — circular import, compile error
```

```go
// package user
import "github.com/you/proj/order"   // user imports order

// package order
import "github.com/you/proj/user"    // ❌ order imports user → CYCLE → compile error
```

JavaScript allows circular imports (with caveats and often subtle bugs). Go bans them outright. This forces cleaner architecture — you must think about dependency direction.

### Breaking a Cycle

When you hit a circular import, it's a design smell. Common fixes:

**1. Extract shared types to a third package:**

```
Before (cycle):
  user ←→ order        (they reference each other's types)

After (no cycle):
  user → types         (both depend on a shared types package)
  order → types
```

**2. Use interfaces to invert the dependency:**

```go
// Instead of order importing user directly, define an interface in order
package order

type UserGetter interface {
    GetUser(id int) (*User, error)
}

// user package implements it, order depends only on the interface
```

**3. Merge the packages** if they're truly that coupled — maybe they belong together.

Circular import errors are annoying at first but they're a feature: they prevent the tangled dependency graphs that plague large JS codebases.

---

## Package Naming Conventions

Go has strong naming conventions for packages:

```go
// ✅ Good package names — short, lowercase, single word, no underscores
package http
package json
package user
package auth

// ❌ Bad package names
package userManagement     // not lowercase, too long
package user_service       // no underscores
package utils              // too vague — what's in it?
package common             // meaningless
```

Guidelines:

- **Short and lowercase.** `http`, not `HTTPUtilities`.
- **No underscores or mixedCaps.** `package userauth`, not `user_auth` or `userAuth`.
- **Singular, usually.** `package user`, not `package users`.
- **Avoid generic names** like `util`, `common`, `helpers`, `misc`. They become dumping grounds. Name by what the package _does_.
- **The package name is part of the API.** Callers write `user.New()`, so avoid stutter: don't name a function `user.NewUser()` (it reads `user.NewUser` — redundant). Prefer `user.New()`.

### Avoiding Stutter

```go
// ❌ Stutter — "user.UserService" repeats "user"
package user
type UserService struct {}     // becomes user.UserService

// ✅ Clean — the package already says "user"
package user
type Service struct {}         // becomes user.Service — reads nicely
```

Since the package name qualifies every member, you don't repeat it in member names. `user.Service`, `http.Client`, `json.Decoder` — clean. This is a subtle but important idiom.

---

## A Realistic Package Layout

Here's how a small-to-medium Go project is typically organized (we'll go deeper in Chapter 17):

```
myapp/
├── go.mod
├── go.sum
├── cmd/                        ← entry points (executables)
│   └── server/
│       └── main.go             ← package main
├── internal/                   ← private application code
│   ├── user/                   ← package user (domain)
│   │   ├── user.go
│   │   └── service.go
│   ├── order/                  ← package order (domain)
│   │   └── order.go
│   └── store/                  ← package store (data access)
│       └── postgres.go
├── pkg/                        ← public libraries (if you intend others to import)
│   └── validator/
│       └── validator.go
└── api/                        ← API definitions, handlers
    └── handler.go
```

Conventions you'll see:

- **`cmd/`** — each subdirectory is an executable's `main` package. A project can have multiple binaries (`cmd/server`, `cmd/cli`, `cmd/worker`).
- **`internal/`** — private code, not importable outside the module.
- **`pkg/`** — public, reusable libraries (use only if you genuinely want external consumers; many projects skip it).

Don't over-engineer early. A small project can be a flat directory. Structure grows as the project does.

---

## Common Mistakes

### Mistake 1: Expecting JS-style relative imports

```go
// ❌ Go doesn't do relative imports
import "./store"
import "../utils"

// ✅ Use the full module path
import "github.com/you/myproject/store"
```

### Mistake 2: Capitalizing things that should be private

```go
// ❌ Exporting internal helpers pollutes your API
func ParseInternal() {}    // capital P — now part of public API

// ✅ Lowercase for internal helpers
func parseInternal() {}    // private to the package
```

Export only what callers genuinely need. Everything else stays lowercase. A smaller public API is easier to maintain.

### Mistake 3: Generic "utils" packages

```go
// ❌ A dumping ground — grows into chaos
package utils
func FormatDate() {}
func ParseJSON() {}
func HashPassword() {}
func ValidateEmail() {}

// ✅ Organize by purpose
package dateformat
package jsonutil
package auth
package validate
```

### Mistake 4: Stutter in names

```go
// ❌ http.HTTPClient, user.UserService — redundant
// ✅ http.Client, user.Service — the package name already provides context
```

### Mistake 5: Not running go mod tidy

After adding/removing imports, run `go mod tidy`. Otherwise your `go.mod` drifts out of sync with your actual dependencies, and builds may fail or carry phantom dependencies.

---

## Exercises

The two exercises in [`exercises/`](./exercises/) are cross-package by design — a
Packages chapter deserves exercises that actually cross a package boundary. Note
that both stubs ship **without their imports**: writing the `import` line yourself,
with the full module path, is part of the drill.

### Exercise 1: Visibility across a package boundary (`catalog`)

Implement the `catalog` package so its exported API (`New`, `SKU`, the `Product`
fields) works, while `sku` and `slugify` stay unexported. The test lives in the
parent `exercises` package and imports `catalog`, so it can touch only the exported
names — the compiler enforces the boundary for you.

### Exercise 2: One-way dependency direction (`metrics` → `report`)

Implement `metrics.Mean` (a leaf package that imports nothing from this project)
and `report.Summary` (one layer up, importing `metrics`). The arrow runs one way,
`report → metrics`; if `metrics` ever imported `report` back, the compiler would
reject the cycle. This is the muscle you'll use to keep import graphs acyclic.

The four hands-on tasks below can't be unit-tested — they're about creating real
modules and watching the toolchain react. Do them in a scratch directory outside
this repo.

### Hands-on A: Multi-Package Project

Create a module `github.com/you/calc` with this structure:

```
calc/
├── go.mod
├── main.go              ← package main, uses the operations
└── operations/
    ├── basic.go         ← package operations: Add, Subtract (exported)
    └── advanced.go      ← package operations: Power, sqrt (sqrt unexported)
```

Verify: `main` can call `operations.Add` and `operations.Power`, but NOT `operations.sqrt`.

### Hands-on B: Internal Package

Add an `internal/config` package to the above project. Put a `Load()` function in it. Import it from `main`. Verify it works within the module. Then reason about why a different module couldn't import it.

### Hands-on C: Break a Circular Import

Deliberately create a circular import: package `a` imports `b`, package `b` imports `a`. Observe the compile error. Then fix it by extracting the shared type into a package `c` that both depend on.

### Hands-on D: Add a Real Dependency

In a new module, `go get github.com/google/uuid`. Write a program that generates and prints a UUID:

```go
import "github.com/google/uuid"

func main() {
    id := uuid.New()
    fmt.Println(id.String())
}
```

Then run `go mod tidy` and inspect how `go.mod` and `go.sum` changed. Examine the module cache location with `go env GOMODCACHE`.

---

## Key Takeaways

1. **Package = directory, module = project.** Every file belongs to one package; packages map to directories. A module (one `go.mod`) groups packages and versions them together.

2. **Capitalization controls visibility.** Uppercase = exported (public), lowercase = unexported (private). No `export` keyword. The visibility is visible in the name everywhere it's used.

3. **Imports use full module paths**, not relative paths. Your own packages: `github.com/you/proj/store`. The tooling autocompletes these.

4. **`internal/` enforces module-level privacy.** Code in `internal/` can't be imported outside your module — compiler-enforced encapsulation.

5. **No circular imports, ever.** A↔B cycles are compile errors. Break them by extracting shared types or inverting with interfaces. This forces cleaner architecture.

6. **`go mod tidy` keeps dependencies in sync** with your code. Run it after changing imports. Commit both `go.mod` and `go.sum`.

7. **Package naming: short, lowercase, no stutter.** `user.Service`, not `user.UserService`. Avoid generic `utils`/`common` dumping grounds.

---

## 🧭 Navigation

| Direction    | Link                                                     |
| ------------ | -------------------------------------------------------- |
| **Previous** | [← Chapter 10: Collections](../10-collections/README.md) |
| **Next**     | [Chapter 12: Generics →](../12-generics/README.md)       |
