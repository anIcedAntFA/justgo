# Chapter 14: Building CLIs & JSON

> **Turn `os.Args` into typed options with the `flag` package and git-style
> subcommands, and move structured data in and out of files with `encoding/json` —
> the two standard-library skills the `gorg` milestone runs on. No third-party
> dependencies.**

## TL;DR

A real command-line tool needs two things the language chapters haven't covered yet.
First, a way to turn the raw `os.Args` slice into typed options and subcommands — that
is the `flag` package, and it's enough for a serious CLI without reaching for `cobra`.
Second, a way to persist structured data — that is `encoding/json`, Go's built-in,
reflection-driven serializer. Put together they are everything `gorg` needs: its
`--dry-run` / `--recursive` flags, its `stats` / `undo` / `rules` subcommands, its JSON
rules config, and its JSON undo journal. This chapter is stdlib-only on purpose — the
day you outgrow it is the day Chapter 11 (dependency management) becomes real.

---

## `os.Args` and exit codes

Every program starts from one slice of strings. `os.Args[0]` is the program name (the
path it was invoked as); `os.Args[1:]` is everything the user typed after it.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("program:", os.Args[0])
	fmt.Println("arguments:", os.Args[1:])
}
```

Coming from Node this is `process.argv`, with one difference worth internalizing:
Node's `process.argv[0]` is the `node` binary and `[1]` is your script — Go's `Args[0]`
is _your_ program directly, and the real arguments start at `[1]`.

### Exit codes and the `os.Exit` trap

A CLI communicates success or failure through its **exit code**: `0` means success,
non-zero means failure (the convention shells and CI rely on). `os.Exit(code)` sets it
and terminates immediately — and _immediately_ is the catch:

```go
func run() {
	f, _ := os.Open("data.txt")
	defer f.Close()      // ⚠️ this defer will NOT run...
	os.Exit(1)           // ...because os.Exit skips every pending defer
}
```

`os.Exit` does not unwind the stack, so **no deferred call fires** — no file flush, no
unlock, no cleanup. The idiom is to keep `main` thin: do the work in a function that
**returns an `error`**, and call `os.Exit` (or `log.Fatal`) in exactly one place, after
the deferred cleanup has already run.

```go
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "gorg:", err)
		os.Exit(1)
	}
}
```

Note `os.Stderr` for errors and diagnostics; keep `os.Stdout` for the program's real
output, so a user can pipe the output without the error text polluting it.

## The `flag` package

Parsing `os.Args` by hand gets miserable fast. The `flag` package does it for you.
Define each flag with its type, name, default, and help text — each constructor returns
a **pointer** to where the parsed value will land:

```go
name := flag.String("name", "World", "who to greet")
loud := flag.Bool("loud", false, "shout the greeting")
retries := flag.Int("retries", 3, "how many times to retry")

flag.Parse() // parse os.Args[1:] — AFTER all flags are defined, BEFORE you read them

fmt.Println(*name, *loud, *retries) // dereference the pointers to read values
```

The ordering rule is not optional: **every flag must be defined before `flag.Parse()`,
and no flag may be read before it.** Read a `*flag` before `Parse` and you get its
default, silently.

If you'd rather bind to an existing variable than juggle pointers, use the `*Var` forms:

```go
var cfg struct {
	Name string
	Loud bool
}
flag.StringVar(&cfg.Name, "name", "World", "who to greet")
flag.BoolVar(&cfg.Loud, "loud", false, "shout the greeting")
flag.Parse()
```

### Positional arguments and help

After `Parse`, the **non-flag** arguments are available via `flag.Args()` (and
`flag.Arg(i)`, `flag.NArg()`). That's how `gorg <dir> --dry-run` gets its `<dir>`:

```go
dryRun := flag.Bool("dry-run", false, "preview only")
flag.Parse()

dir := "."
if flag.NArg() > 0 {
	dir = flag.Arg(0) // the positional <dir>
}
```

Flags accept several forms — `-flag`, `--flag`, `-flag=value`, and `-flag value` — and
`-h` / `-help` prints auto-generated usage from your help strings. **One booleans-only
gotcha:** the space form does _not_ work for a `bool`. `-loud` (bare) means true, and
`-loud=false` sets it explicitly, but `-loud false` treats `false` as a **positional
argument**, not the flag's value.

> Runnable demo: [`examples/flags/`](./examples/flags/) —
> `go run . --name Go --loud extra`.

### Why not `cobra`?

If you've seen Go CLIs before, you've seen `spf13/cobra`. It's excellent, and it's also
a third-party dependency with its own concepts. For Part 1 the whole point is to see how
little you need: `flag` covers flags, defaults, help, and (below) subcommands with zero
imports beyond the stdlib. Reaching for `cobra`/`viper` is a deliberate _post_-Part-1
evolution of `gorg` — see [`gorg/PLAN.md`](../gorg/PLAN.md) §11. Learn the floor before
you buy the elevator.

## Subcommand dispatch with `flag.FlagSet`

`gorg stats`, `gorg undo`, `gorg rules` — a tool with subcommands (like `git`) needs a
router. You don't need a framework for it. The package-level `flag` functions all live
on a hidden default `FlagSet`; for subcommands you make **one `FlagSet` per command** so
each has its own isolated flags, and you switch on the first argument:

```go
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: gorg <organize|stats|undo> [flags]")
		os.Exit(2)
	}

	switch os.Args[1] { // the subcommand name
	case "stats":
		stats(os.Args[2:]) // hand the REST of the args to the subcommand
	case "undo":
		undo(os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "unknown subcommand %q\n", os.Args[1])
		os.Exit(2)
	}
}

func stats(args []string) {
	fs := flag.NewFlagSet("stats", flag.ExitOnError)
	asJSON := fs.Bool("json", false, "emit stats as JSON")
	_ = fs.Parse(args) // parse only this subcommand's args
	// ... use *asJSON and fs.Args() ...
}
```

The `ErrorHandling` argument to `NewFlagSet` decides what a parse error does:
`ExitOnError` calls `os.Exit(2)` (and `os.Exit(0)` for `-h`) — convenient for a
top-level command; `ContinueOnError` returns the error so you can handle it — better
when you want control; `PanicOnError` panics. Each subcommand's `-h` prints only _its_
flags, which is exactly the git-style help users expect.

> Runnable demo: [`examples/subcommands/`](./examples/subcommands/) —
> `go run . stats --json ./dir` and `go run . undo`.

## `encoding/json` — Marshal & Unmarshal

`encoding/json` is Go's built-in serializer. Two functions cover the in-memory case:

```go
type Rule struct {
	Category   string
	Extensions []string
}

r := Rule{Category: "Images", Extensions: []string{".jpg", ".png"}}

data, err := json.Marshal(r)     // Go value → JSON bytes
// data == {"Category":"Images","Extensions":[".jpg",".png"]}

var back Rule
err = json.Unmarshal(data, &back) // JSON bytes → Go value (note the &: needs a pointer)
```

This is `JSON.stringify` / `JSON.parse` from JS — with two Go differences that matter.
First, **it's type-checked**: you unmarshal into a concrete struct, not a free-form
object, so the shape is known at compile time. Second, **errors are values**: every call
returns an `error` you must check. There is no silent `NaN` and no thrown exception to
forget — `Unmarshal` into a non-pointer, or malformed JSON, comes back as an `error`.

## Struct tags

Notice the JSON above used `Category` and `Extensions` — capitalized, because of Go's
single most important JSON rule:

> **`encoding/json` only sees exported (capitalized) fields.** An unexported field is
> invisible — never marshaled, never unmarshaled.

It uses reflection, and reflection can't reach unexported fields. That collides with the
FE habit of `snake_case` or `camelCase` JSON keys, so you bridge the two with **struct
tags**:

```go
type Rule struct {
	Category   string   `json:"category"`
	Extensions []string `json:"extensions"`
	Note       string   `json:"note,omitempty"` // dropped from output when ""
	internal   string   `json:"-"`              // unexported anyway; also, "-" = never
}
```

- `json:"category"` renames the key on the wire while the Go field stays idiomatic.
- `omitempty` omits the field when it holds the type's empty value (`false`, `0`, `nil`,
  or an empty string/slice/map) — perfect for optional config keys so the file stays
  clean.
- `json:"-"` always omits a field, even an exported one (secrets, computed fields).

On the way _in_, `Unmarshal` matches keys to fields preferring an exact match but also
accepting a **case-insensitive** one, and it simply **ignores** JSON keys with no
matching field (unless you opt into strictness — see below). Missing keys leave their Go
field at its zero value, no error.

## JSON round-trip to files (config + journal)

`Marshal`/`Unmarshal` work on `[]byte` in memory. When the source or destination is a
**stream** — a file, an HTTP body, a socket — use an `Encoder`/`Decoder` so you don't
buffer the whole payload:

```go
// Write: stream a value straight to a file.
f, err := os.Create(path)
if err != nil { /* handle */ }
defer f.Close()

enc := json.NewEncoder(f)
enc.SetIndent("", "  ")  // pretty-print for a human-editable config
if err := enc.Encode(rules); err != nil { /* handle */ }
// Encode writes the value FOLLOWED BY A NEWLINE — handy for append-only journals.

// Read: stream it back.
f2, err := os.Open(path)
defer f2.Close()
var loaded []Rule
if err := json.NewDecoder(f2).Decode(&loaded); err != nil { /* handle */ }
```

For a small config you already hold in memory, the pair `json.MarshalIndent(v, "", "  ")`

- `os.WriteFile(path, data, 0o644)` is just as good and reads a little plainer. Either
  way, `gorg` puts these files under `os.UserConfigDir()` — `~/.config` on Linux — so the
  rules and the undo journal live where a user expects them.

That's the whole persistence story for the milestone: the **rules config** is a JSON
file you `MarshalIndent` and hand-edit; the **undo journal** is a JSON record you
`Encode` after a run and `Decode` to reverse it.

> Runnable demo: [`examples/json-roundtrip/`](./examples/json-roundtrip/) — encode
> rules to a temp file with indentation, decode them back.

### One strictness knob

By default an unknown JSON key is silently ignored. When a typo in a config file should
be an error rather than a no-op, opt in:

```go
dec := json.NewDecoder(f)
dec.DisallowUnknownFields() // now an unexpected key fails the decode
```

> **Ecosystem note (not for Part 1):** Go 1.25 introduced an experimental
> `encoding/json/v2` with revised behavior. `v1` — the package this chapter teaches —
> stays the default. Good to know the ground is shifting; nothing to act on yet.

---

## Common Mistakes

### Mistake 1: Lowercase struct fields vanish from JSON

```go
type File struct {
	name string `json:"name"` // ❌ unexported — encoding/json can't see it; always ""
	Size int64  `json:"size"` // ✅ exported
}
```

The tag doesn't help — the field must be **capitalized** to cross the wire at all. This
is the number-one JSON surprise for people coming from languages where every field is
public.

### Mistake 2: Reading a flag before `flag.Parse()`

```go
verbose := flag.Bool("v", false, "verbose")
if *verbose { ... }   // ❌ always the default — Parse hasn't run yet
flag.Parse()
```

Define all flags, `Parse()`, _then_ read.

### Mistake 3: `-boolflag value` with a space

```go
// gorg --dry-run true ./dir
```

For a `bool`, `true` is not consumed by `--dry-run` — it becomes a positional argument,
and `flag.Arg(0)` is `"true"`, not `"./dir"`. Use `--dry-run` alone or `--dry-run=true`.

### Mistake 4: `os.Exit` inside a helper, skipping cleanup

`os.Exit` runs **no** deferred calls. A `defer f.Close()` or `defer mu.Unlock()` above it
never fires. Return an error up to `main` and exit there.

### Mistake 5: Ignoring the error from `Unmarshal`/`Decode`

```go
json.Unmarshal(data, &cfg) // ❌ dropped error — malformed input leaves cfg half-set
```

There's no exception to catch and no `NaN` to notice later; the error _is_ the signal.
Check it every time.

## Exercises

Coding exercises live under [`exercises/`](./exercises/) as stubs with table-driven
tests. Each test opens with a `t.Skip(...)`; delete it once you've implemented the
function, then run `go test ./...` until it's green.

- **`greetflags`** — parse `-name` / `-loud` from an args slice with a `flag.FlagSet`
  and build the greeting. (flag basics, `ContinueOnError`.)
- **`config`** — round-trip a `Config` struct through JSON: `EncodeConfig` (indented)
  and `DecodeConfig`. (tags, `omitempty`, Marshal/Unmarshal.)
- **`dispatch`** — route a subcommand name (`""`/`organize`/`stats`/`undo`/`rules`) to
  its action, wrapping `ErrUnknownCommand` for anything else. (subcommand routing,
  error wrapping.)

## Key Takeaways

1. **`os.Args[1:]` are the real arguments**, and a CLI signals outcome through its exit
   code — `0` success, non-zero failure.
2. **`flag` is enough for a serious CLI.** Define flags (they return pointers) _before_
   `flag.Parse()`, read positionals with `flag.Args()`. Reach for `cobra` only when you
   truly outgrow it.
3. **One `flag.FlagSet` per subcommand** — switch on `os.Args[1]`, hand `os.Args[2:]` to
   the chosen command. That's the entire "router".
4. **`encoding/json` only sees exported fields;** struct tags (`json:"name,omitempty"`,
   `json:"-"`) control the wire shape.
5. **`Encoder`/`Decoder` for streams and files** (`Encode` appends a newline);
   `Marshal`/`MarshalIndent` for in-memory bytes.
6. **Every flag-parse, `Unmarshal`, and `Decode` returns an `error`.** Handle it — Go
   gives you no silent coercion and no thrown exception to forget.
7. **`os.Exit` skips `defer`s.** Keep `main` thin: return errors, exit in one place.

---

## 🧭 Navigation

| Direction    | Link                                                   |
| ------------ | ------------------------------------------------------ |
| **Previous** | [← Chapter 13: Testing](../13-testing/README.md)       |
| **Next**     | `gorg` — File Organizer CLI _(milestone, not started)_ |
| **Overview** | [Part 1 — Foundation](../README.md)                    |
