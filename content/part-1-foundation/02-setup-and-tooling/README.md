# Chapter 02: Setup & Tooling

> **Install Go, configure your editor, understand the toolchain. Your Go environment from zero to productive.**

## TL;DR

Go's toolchain is remarkably self-contained compared to the JS ecosystem. One install gives you compiler, formatter, test runner, dependency manager, and more. No Webpack, no Babel, no Prettier config debates. This chapter gets your Arch Linux setup production-ready and explains every tool you'll use daily.

---

## Installing Go on Arch Linux

Arch keeps Go updated in the official repos. As of writing, `go` package is at version 1.26.3.

### Option A: pacman (Recommended for Arch)

```bash
# Install Go from official repos
sudo pacman -S go

# Verify
go version
# → go version go1.26.3 linux/amd64
```

That's it. No version manager, no nvm equivalent needed (yet). Arch repos track upstream closely.

### Option B: Official tarball (pin exact version)

If you want to control the exact version (useful if Arch repos lag behind):

```bash
# Download from go.dev
wget https://go.dev/dl/go1.26.3.linux-amd64.tar.gz

# Remove old installation (if any) and extract
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.26.3.linux-amd64.tar.gz

# Add to PATH (in your .zshrc or .bashrc)
export PATH=$PATH:/usr/local/go/bin

# Verify
go version
```

### Option C: mise (recommended for this repo)

[mise](https://mise.jdx.dev) is a polyglot version manager (think `nvm` + `asdf` for every tool). This repo pins Go **and its whole toolchain** in [`mise.toml`](../../../mise.toml), so one command installs everything at the exact versions:

```bash
mise install        # or: just setup  — installs go, gofumpt, golangci-lint, just, dlv, …
go version          # → go version go1.26.3 linux/amd64
```

Why this over `pacman` here: the version is pinned per-repo and reproducible across machines, and the same file also manages `gofumpt`, `golangci-lint`, `just`, `dlv`, and more — no global installs, no version drift. This is how the justgo repo is set up.

### Environment Variables

Add these to your shell config (`~/.zshrc`, `~/.bashrc`, or wherever you manage env):

```bash
# Go environment — add to your shell config
export GOPATH=$HOME/go                    # where go install puts binaries & cached modules
export GOBIN=$GOPATH/bin                  # where go install puts binaries
export PATH=$PATH:$GOBIN                  # so installed Go tools are in your PATH
```

Verify everything:

```bash
go env GOPATH GOBIN GOROOT
# GOPATH=/home/youruser/go
# GOBIN=/home/youruser/go/bin
# GOROOT=/usr/lib/go     (Arch) or /usr/local/go (tarball)
```

### JS Comparison: What Just Happened

```
JavaScript world:
  nvm install 20        → install Node runtime
  npm init              → create package.json
  npm install           → pull 500MB of node_modules
  npx webpack ...       → bundler
  npx prettier ...      → formatter
  npx eslint ...        → linter
  npx jest ...          → test runner
  (10+ separate tools before writing any code)

Go world:
  sudo pacman -S go     → install Go (compiler + ALL tools)
  go mod init           → create go.mod
  go build              → compile
  go run                → compile + run
  go test               → test
  go fmt                → format
  go vet                → lint (basic)
  (one install, everything included)
```

---

## The Go CLI — Your Swiss Army Knife

Unlike JS where you reach for npm/npx + separate tools, the `go` command handles almost everything. Here's every subcommand you'll actually use:

### Daily Commands

```bash
# Run a file directly (compile + execute, no binary saved)
go run main.go
go run .                      # run current package

# Build a binary
go build                      # outputs binary named after module
go build -o myapp             # custom output name

# Run tests
go test ./...                 # all tests in all packages
go test -v ./...              # verbose output
go test -run TestName ./pkg   # run specific test
go test -race ./...           # detect race conditions
go test -cover ./...          # show coverage percentage

# Format code (you'll rarely run this manually — editor does it on save)
go fmt ./...
gofmt -w .                    # same thing, different entry point

# Vet (static analysis — catches bugs the compiler misses)
go vet ./...
```

### Module Management

```bash
# Initialize a new module (like npm init)
go mod init github.com/yourname/projectname

# Add a dependency (like npm install)
go get github.com/some/package@latest
go get github.com/some/package@v1.2.3    # specific version

# Clean up unused deps (like npm prune)
go mod tidy

# Show dependency graph
go mod graph

# Vendor dependencies locally (like node_modules but explicit)
go mod vendor
```

### Useful Utilities

```bash
# Install a Go binary tool (like npx but permanent)
go install github.com/some/tool@latest

# View documentation in terminal
go doc fmt.Println
go doc net/http.HandleFunc

# Show all environment variables
go env

# Cross-compile (one of Go's superpowers)
GOOS=linux GOARCH=amd64 go build -o myapp-linux
GOOS=darwin GOARCH=arm64 go build -o myapp-mac
GOOS=windows GOARCH=amd64 go build -o myapp.exe
```

Cross-compilation deserves special attention. That last block builds binaries for Linux, macOS, and Windows — **from your Arch machine** — with zero extra setup. No Docker, no CI matrix. Just two environment variables. This is why Go is the #1 choice for CLI tools that need to ship cross-platform.

---

## Go Modules: go.mod Explained

Every Go project starts with `go mod init`. This creates `go.mod` — the equivalent of `package.json`.

```bash
mkdir myproject && cd myproject
go mod init github.com/yourname/myproject
```

This creates:

```
// go.mod
module github.com/yourname/myproject

go 1.26
```

When you add dependencies, it grows:

```
module github.com/yourname/myproject

go 1.26

require (
    github.com/labstack/echo/v4 v4.13.3
    github.com/mattn/go-sqlite3 v1.14.24
)

require (
    // indirect dependencies (auto-managed, don't touch)
    golang.org/x/crypto v0.31.0 // indirect
    golang.org/x/net v0.33.0 // indirect
)
```

Alongside `go.mod`, Go generates `go.sum` — cryptographic hashes of every dependency. Similar to `package-lock.json` but with actual security verification.

### Key Differences from package.json

| Aspect             | package.json                   | go.mod                                |
| ------------------ | ------------------------------ | ------------------------------------- |
| Package name       | Arbitrary string ("my-app")    | URL-based (`github.com/user/repo`)    |
| Version resolution | Semver ranges (`^1.2.3`)       | Exact versions (`v1.2.3`) only        |
| Lock file          | `package-lock.json` (versions) | `go.sum` (versions + hashes)          |
| Dependency storage | `node_modules/` per project    | `$GOPATH/pkg/mod/` shared cache       |
| Unused deps        | Warning (maybe)                | **Compile error**                     |
| Scripts            | `"scripts": {"test": "jest"}`  | No scripts — use Makefile or taskfile |

Notice: Go doesn't support version **ranges**. If you depend on `v1.2.3`, you get exactly `v1.2.3`. This eliminates an entire class of "works on my machine" bugs caused by different minor versions resolving differently.

---

## Editor Setup: VS Code

You already use VS Code for JS/TS. The Go setup is simpler — one extension does everything.

### Step 1: Install the Go Extension

```
Extension ID: golang.go
```

Install it. That's the only Go-specific extension you need. It bundles:

- **gopls** — the official Go language server (auto-complete, go-to-definition, rename, find references)
- **Delve** — Go debugger
- **Test runner** — run/debug individual tests from the editor
- **Code formatting** — auto-runs `gofmt` on save

When you first open a `.go` file, the extension will prompt you to install Go tools. Click "Install All". It installs `gopls`, `dlv` (Delve debugger), `staticcheck`, and others into your `$GOBIN`.

### Step 2: VS Code Settings for Go

Add to your `settings.json` (or workspace settings):

```jsonc
{
  // === Go-specific ===

  // Format on save (non-negotiable in Go)
  "[go]": {
    "editor.formatOnSave": true,
    "editor.defaultFormatter": "golang.go",
    "editor.codeActionsOnSave": {
      "source.organizeImports": "explicit"
    }
  },

  // Use gopls (should be default, but be explicit)
  "go.useLanguageServer": true,

  // Use golangci-lint instead of default linter
  "go.lintTool": "golangci-lint",
  "go.lintOnSave": "workspace",

  // Test flags
  "go.testFlags": ["-v", "-race", "-count=1"],

  // Show test coverage highlights in editor
  "go.coverOnSave": true,
  "go.coverageDecorator": {
    "type": "highlight",
    "coveredHighlightColor": "rgba(64,128,64,0.2)",
    "uncoveredHighlightColor": "rgba(128,64,64,0.2)"
  },

  // gopls settings
  "gopls": {
    "formatting.gofumpt": true, // stricter formatting than gofmt
    "ui.semanticTokens": true, // better syntax highlighting
    "ui.diagnostic.analyses": {
      "unusedwrite": true, // detect unused writes
      "useany": true, // suggest 'any' over 'interface{}'
      "nilness": true // nil pointer analysis
    }
  }
}
```

### Step 3: Keybindings You'll Use

These come from the Go extension — learn them early:

| Action             | Keybinding                                              | What it does                          |
| ------------------ | ------------------------------------------------------- | ------------------------------------- |
| Go to definition   | `F12`                                                   | Jump to function/type source          |
| Go back            | `Ctrl+-`                                                | Return from definition jump           |
| Find references    | `Shift+F12`                                             | Where is this used?                   |
| Rename symbol      | `F2`                                                    | Rename across entire project          |
| Run test at cursor | `Ctrl+Shift+P` → "Go: Test Function At Cursor"          | Run single test                       |
| Toggle test file   | `Ctrl+Shift+P` → "Go: Toggle Test File"                 | Jump between `foo.go` ↔ `foo_test.go` |
| Generate test      | `Ctrl+Shift+P` → "Go: Generate Unit Tests For Function" | Scaffolds test                        |
| Add struct tags    | `Ctrl+Shift+P` → "Go: Add Tags To Struct Fields"        | Adds `json:"field_name"`              |

### Alternative Editors

If VS Code isn't your thing:

**GoLand** (JetBrains) — the premium option. Best debugging, refactoring, and code intelligence. Paid ($249/yr first year). 28% market share among Go devs. If you use WebStorm for JS, GoLand is the natural choice.

**Neovim + gopls** — if you're already a Neovim user on Arch. gopls works with any LSP-compatible editor. Setup via `nvim-lspconfig` or `coc.nvim`. Fastest editing experience once configured.

**Zed** — fast Rust-based editor with native Go support. If you use Zed, this repo ships a working config — see the section below.

### Zed Setup

Zed has Go support built in: tree-sitter highlighting, **gopls** as the language server, and a native **Delve** debug adapter. Install `gopls` and `dlv` yourself — `go install golang.org/x/tools/gopls@latest` (here they come from `mise`, via `just setup`) — and make sure `$HOME/go/bin` is on `PATH`. Zed's docs specifically recommend the `go install` build over a distro/Homebrew package.

Unlike VS Code, **none** of the formatting/analysis behavior is on by default — you configure it. This repo already does, in [`.zed/settings.json`](../../../.zed/settings.json) and [`.zed/debug.json`](../../../.zed/debug.json):

- **Format on save** via a `formatter` array that runs in order: the `source.organizeImports` code action, then gopls (which applies **gofumpt** because `"gofumpt": true` is set). Note: the old `code_actions_on_format` key is gone — the `formatter` array replaces it, and inside it you use the object form `{ "language_server": { "name": "gopls" } }`.
- **gopls settings** under `lsp.gopls.initialization_options`: `gofumpt`, `local` (import grouping by module path), `staticcheck`, `analyses` (`nilness`, `unusedwrite`, `useany`), and inlay-hint content.
- **Inlay hints** need the editor-level `"inlay_hints": { "enabled": true }` **and** the gopls `hints` map — both, or nothing shows.
- **Debugging**: `.zed/debug.json` has Delve configs (debug the package / debug tests); open with the command palette → `debugger: start`.

**golangci-lint** is not native to Zed and gopls doesn't run it — it needs a separate extension plus the `golangci-lint-langserver` binary. This repo skips that: golangci-lint runs via `just check` and CI, and gopls' `staticcheck` covers the common cases in the editor.

> Heads-up: organize-imports-on-save has had reliability bugs in Zed. Ordering the code action **before** gopls in the `formatter` array is the arrangement that works.

Sources: [Zed Go docs](https://zed.dev/docs/languages/go) · [configuring languages](https://zed.dev/docs/configuring-languages) · [gopls settings](https://github.com/golang/tools/blob/master/gopls/doc/settings.md).

---

## Essential Go Tools

### gofmt / goimports — Formatting

```bash
# gofmt is built-in, formats Go code
gofmt -w .          # format all files in current directory

# goimports does gofmt + manages imports automatically
go install golang.org/x/tools/cmd/goimports@latest
goimports -w .
```

You won't run these manually — your editor does it on save. But know they exist.

**gofumpt** is a stricter version of gofmt (extra formatting rules). Recommended — we enabled it in the VS Code config above.

```bash
go install mvdan.cc/gofumpt@latest
```

### golangci-lint — Linting

The standard linting tool for Go. It runs 100+ linters in parallel. Think of it as ESLint but for Go, except it comes pre-configured with sane defaults.

```bash
# Install
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run
golangci-lint run          # lint current project
golangci-lint run --fix    # auto-fix where possible
```

Create `.golangci.yml` in your project root for configuration:

```yaml
# .golangci.yml — starter config
version: "2"

run:
  timeout: 5m

linters:
  default: standard
  enable:
    - errcheck        # unchecked errors
    - govet           # suspicious constructs
    - staticcheck     # advanced static analysis
    - unused          # unused code
    - gosimple        # simplification suggestions
    - ineffassign     # ineffectual assignments
    - gocritic        # opinionated style checks
    - errname         # error naming conventions
    - gofumpt         # strict formatting

formatters:
  enable:
    - gofumpt

linters-settings:
  gocritic:
    enabled-tags:
      - diagnostic
      - style
      - performance
```

### Delve — Debugger

```bash
# Already installed by VS Code Go extension, but if needed:
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug from command line
dlv debug ./main.go

# Attach to running process
dlv attach <pid>
```

VS Code integrates Delve seamlessly — set breakpoints, inspect variables, step through code. Works like Chrome DevTools but for Go.

### Air — Live Reload (for development)

Go doesn't have nodemon, but Air fills that role:

```bash
go install github.com/air-verse/air@latest

# Run in project root — watches files, rebuilds on change
air
```

Configure with `.air.toml`:

```toml
# .air.toml
root = "."
tmp_dir = "tmp"

[build]
cmd = "go build -o ./tmp/main ."
bin = "./tmp/main"
delay = 1000 # ms after file change before rebuild
exclude_dir = ["tmp", "vendor", ".git"]
exclude_regex = ["_test.go"]
include_ext = ["go", "html", "yaml"]

[log]
time = true
```

Useful during Part 2 when building HTTP servers. Not needed for CLI tools in Part 1.

---

## Your First Go Program

Let's verify everything works end-to-end.

### Create the project

**In this repo**, your first program already lives as a Chapter example — no
`go mod init` needed, because all Chapter code shares the **root module** via
`go.work` ([ADR-0002](../../../docs/adr/0002-multi-module-workspace.md)):

```bash
cd content/part-1-foundation/02-setup-and-tooling/examples/hello
# main.go + main_test.go are already here — run: go run .
```

If you'd rather scratch **outside** the repo (a throwaway you won't commit), that
needs its own module:

```bash
mkdir ~/projects/hello-go && cd ~/projects/hello-go
go mod init example.com/hello-go
```

### Write the code

`examples/hello/main.go` (already created for you):

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello from Go!")
	
	// Multiple return values — you'll see this everywhere
	name, age := "Go Developer", 2026-2007
	fmt.Printf("I'm a %s, Go is %d years old\n", name, age)
}
```

### Run it

```bash
# Run directly (compile + execute, binary discarded)
go run .
# → Hello from Go!
# → I'm a Go Developer, Go is 19 years old

# Build a binary
go build -o hello
./hello
# → same output

# Check the binary size
ls -lh hello
# → ~1.8MB — that's the entire runtime, no external dependencies

# Cross-compile for other platforms
GOOS=windows GOARCH=amd64 go build -o hello.exe
GOOS=darwin GOARCH=arm64 go build -o hello-mac
ls -lh hello*
# Three binaries for three platforms, built from one machine
```

### Format and lint

```bash
# Format (should be no-op if editor already formatted on save)
gofmt -d .             # show diff (should be empty)

# Vet
go vet ./...           # should show no issues

# Lint
golangci-lint run      # should pass
```

### Test (preview — we'll cover testing properly in Chapter 13)

Create `main_test.go`:

```go
package main

import "testing"

func TestSomething(t *testing.T) {
	expected := 19
	got := 2026 - 2007
	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
```

```bash
go test -v ./...
# === RUN   TestSomething
# --- PASS: TestSomething (0.00s)
# PASS
```

No test framework to install. No config file. `go test` just works.

---

## Project Structure: Where Files Go

For now, keep it simple. You'll learn advanced project layout in Chapter 17. Here's what a minimal Go project looks like:

```
hello-go/
├── go.mod              ← module definition (like package.json)
├── go.sum              ← dependency checksums (like package-lock.json)
├── main.go             ← entry point (package main, func main)
└── main_test.go        ← tests (same package, _test.go suffix)
```

Rules to know now:

- **`package main`** + **`func main()`** = executable entry point
- **`_test.go`** suffix = test file, only compiled during `go test`
- **One package per directory** (unlike JS where one file can export anything)
- **Capitalized names are exported** (public), lowercase are unexported (private)

That last one is unique to Go — no `export` keyword, no `public`/`private`:

```go
func DoSomething() {}  // Exported — other packages can call this (uppercase D)
func doSomething() {}  // Unexported — only this package can call this (lowercase d)

type User struct {
    Name  string       // Exported field
    email string       // Unexported field
}
```

Coming from JS where you explicitly `export`, this feels strange at first. You get used to it in a day.

---

## The go.env and GOPATH Explained

New Go devs (especially from JS) often confuse these. Here's the clarity:

```
GOROOT  → where Go is installed (/usr/lib/go on Arch)
          You never touch this. It's like /usr/bin/node.

GOPATH  → your Go workspace (default: ~/go)
          ├── bin/     → installed binaries (golangci-lint, air, etc.)
          ├── pkg/     → cached compiled packages + module cache
          └── src/     → (legacy, not used with modules)
          
          Think of GOPATH/pkg/mod as a global node_modules cache
          that's shared across ALL your projects.
```

With Go modules (since Go 1.11, 2018), you **don't need to work inside GOPATH**. Put your projects anywhere. The module system handles everything.

---

## Makefile: Your Task Runner

Go doesn't have `npm scripts`. The standard replacement is a `Makefile` (or `Taskfile.yml` if you prefer). Here's a starter:

> **This repo uses [`just`](https://github.com/casey/just), not Make** — see the
> [`justfile`](../../../justfile) and [ADR-0005](../../../docs/adr/0005-just-as-task-runner.md).
> `just` is a task runner (not a build system), so no `.PHONY` boilerplate and no
> tab-vs-space traps. Learn to _read_ a Makefile — Go code in the wild ships one
> constantly — but reach for `just fmt` / `just check` / `just test` here.

```makefile
# Makefile

.PHONY: run build test lint clean

# Default target
all: lint test build

# Run the application
run:
	go run .

# Build binary
build:
	go build -o bin/app .

# Run all tests with race detection
test:
	go test -v -race -cover ./...

# Run linter
lint:
	golangci-lint run

# Format code
fmt:
	gofumpt -w .
	goimports -w .

# Clean build artifacts
clean:
	rm -rf bin/ tmp/

# Install dev tools
tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install mvdan.cc/gofumpt@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/air-verse/air@latest
```

Usage:

```bash
make run          # run the app
make test         # run tests
make lint         # run linter
make build        # build binary
make tools        # install all dev tools (run once)
```

---

## Common Mistakes

### Mistake 1: Installing Go tools with `go get` instead of `go install`

```bash
# Wrong — go get is for adding dependencies to your project
go get github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Right — go install is for installing binary tools
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

`go get` modifies your `go.mod`. `go install` installs a binary to `$GOBIN`. Different purposes.

### Mistake 2: Not having $GOBIN in PATH

You install a tool, then `command not found`. Fix:

```bash
# Make sure this is in your shell config
export PATH=$PATH:$(go env GOPATH)/bin
```

### Mistake 3: Ignoring `go vet`

`go vet` catches real bugs that compile fine but are wrong at runtime — format string mismatches, unreachable code, incorrect struct tags. Run it. Always.

### Mistake 4: Fighting gofmt

Don't. There's no config. There's no debate. Accept the One True Style and move on. The mental energy you save is enormous.

---

## Exercises

### Exercise 1: Complete Setup Verification

Run through the full setup on your Arch machine. Verify each step:

```bash
# 1. Go is installed
go version                    # should show 1.26.x

# 2. Environment is correct  
go env GOPATH                 # should show ~/go or similar
echo $PATH | grep -o 'go/bin' # should find it

# 3. Tools are installed
golangci-lint --version       # should show version
gofumpt --version             # should show version

# 4. Editor works
# Open a .go file in VS Code
# Type fmt. — auto-complete should appear
# Save the file — it should auto-format
# Introduce a bug — red squiggly should appear
```

### Exercise 2: Cross-Compile

Build your hello-go program for 3 platforms. Verify each binary:

```bash
GOOS=linux GOARCH=amd64 go build -o hello-linux
GOOS=darwin GOARCH=arm64 go build -o hello-mac
GOOS=windows GOARCH=amd64 go build -o hello.exe

file hello-linux hello-mac hello.exe
# hello-linux:   ELF 64-bit LSB executable, x86-64 ...
# hello-mac:     Mach-O 64-bit arm64 executable ...
# hello.exe:     PE32+ executable (console) x86-64 ...
```

Imagine shipping a CLI tool to your team across 3 platforms from one `go build` command. No Docker multi-arch builds. No CI matrix. This is why Go dominates CLI tooling.

### Exercise 3: Explore `go doc`

Use `go doc` to explore the standard library from your terminal:

```bash
go doc fmt                    # overview of fmt package
go doc fmt.Println            # specific function
go doc net/http               # the HTTP package you'll use in Part 2
go doc net/http.ListenAndServe
```

Compare this to MDN or npmjs.com — documentation is built into the toolchain, no browser needed.

---

## Key Takeaways

1. **One install, everything included.** `go` CLI replaces npm, npx, jest, prettier, eslint, webpack. Fewer tools = fewer things that break.

2. **`go.mod` is your `package.json`** but simpler — exact versions only, no ranges, unused deps are compile errors.

3. **VS Code + Go extension** is all you need. gopls handles everything. Format on save is mandatory, not optional.

4. **golangci-lint** is the standard linter. Set it up from day one.

5. **Cross-compilation is trivial.** Two env vars, zero extra tooling. This is a Go superpower.

6. **Use a Makefile** for common tasks. Go doesn't have npm scripts, and that's fine.

---

## 🧭 Navigation

| Direction    | Link                                                                 |
| ------------ | -------------------------------------------------------------------- |
| **Previous** | [← Chapter 01: History & Philosophy](./01-history-and-philosophy.md) |
| **Next**     | [Chapter 03: Types & Variables →](./03-types-and-variables.md)       |
