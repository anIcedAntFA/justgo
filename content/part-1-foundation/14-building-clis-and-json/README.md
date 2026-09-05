# Chapter 14: Building CLIs & JSON

> **Parse command-line flags and subcommands with the `flag` package, and read/write
> structured data with `encoding/json` — the two stdlib skills the `gorg` milestone
> needs.**

> [!NOTE]
> This chapter is **scaffolded, not yet written**. The section headings and the
> runnable [`examples/`](./examples/) are in place; the teaching prose (marked
> `TODO`) is authored by the repo owner. The [`exercises/`](./exercises/) ship as
> stubs with skipped tests — see [the exercises note](#exercises).

## TL;DR

> **TODO:** 2–3 lines. A real CLI needs two things Part 1 hasn't covered yet: turning
> `os.Args` into typed options (the `flag` package + manual subcommand dispatch), and
> moving structured data in and out of files as JSON (`encoding/json`). Together they
> are everything `gorg` needs for its flags, its config file, and its undo journal —
> no third-party libraries.

---

## `os.Args` and exit codes

> **TODO:** the raw input — `os.Args[0]` is the program name, `os.Args[1:]` the
> arguments. Exiting: `os.Exit(code)` (and why `defer` does **not** run after it, so
> prefer returning an error up to `main`). JS comparison: `process.argv`,
> `process.exit()`.

## The `flag` package

> **TODO:** define flags with `flag.String`/`flag.Int`/`flag.Bool` (and the `*Var`
> forms), `flag.Parse()`, reading `flag.Args()` for positional args, default values,
> and the auto-generated `-h`/`--help`. Contrast with reaching for `cobra` (that's a
> post-Part-1 upgrade — see [`gorg/PLAN.md`](../gorg/PLAN.md) §11). JS comparison:
> `commander` / `yargs` vs the stdlib.
>
> Runnable demo: [`examples/flags/`](./examples/flags/) — `go run . --name Go --loud`.

## Subcommand dispatch with `flag.FlagSet`

> **TODO:** git-style subcommands (`gorg stats`, `gorg undo`) with **no framework**:
> switch on `os.Args[1]`, then give each subcommand its own `flag.NewFlagSet` and
> `Parse(os.Args[2:])`. Why a `FlagSet` per subcommand instead of package-level flags.
>
> Runnable demo: [`examples/subcommands/`](./examples/subcommands/) —
> `go run . stats ./x` / `go run . undo`.

## `encoding/json` — Marshal & Unmarshal

> **TODO:** `json.Marshal` (Go → bytes) and `json.Unmarshal` (bytes → Go). Zero values
> vs missing fields. Errors are values here too — always check them. JS comparison:
> `JSON.stringify` / `JSON.parse`, but type-checked and error-returning.

## Struct tags

> **TODO:** `json:"name,omitempty"` — renaming fields, `omitempty`, `-` to skip,
> exported-fields-only rule (unexported fields are invisible to `encoding/json`). This
> is Go's answer to the FE habit of camelCase JSON over PascalCase-ish Go fields.

## JSON round-trip to files (config + journal)

> **TODO:** streaming with `json.NewEncoder(w).Encode` / `json.NewDecoder(r).Decode`
> vs `Marshal` + `os.WriteFile`; `MarshalIndent` for human-readable config;
> `os.UserConfigDir()` for where a config file lives. This is exactly how `gorg` loads
> its rules and reads/writes its undo journal.
>
> Runnable demo: [`examples/json-roundtrip/`](./examples/json-roundtrip/) — encode a
> value to a temp file and decode it back.

---

## Common Mistakes

> **TODO:** collect the JS-dev traps, e.g.
>
> - Unexported struct fields silently omitted from JSON (must be capitalized).
> - Calling `flag.Parse()` before defining flags, or reading a flag before `Parse`.
> - Using package-level `flag` for subcommands instead of a `FlagSet` each.
> - `os.Exit` in a helper skipping every `defer` (including flushes/cleanup).
> - Ignoring the `error` from `Unmarshal`/`Decode` (there is no silent `NaN` here).

## Exercises

Coding exercises live under [`exercises/`](./exercises/) as stubs with table-driven
tests. Each test starts with a `t.Skip(...)`; delete it once you've implemented the
function, then run `go test ./...` until green.

- **`greetflags`** — parse `-name` / `-loud` flags into a greeting.
- **`config`** — round-trip a `Config` struct through JSON (tags + `omitempty`).
- **`dispatch`** — route a subcommand name to the right handler.

## Key Takeaways

> **TODO:** the bullets to remember. Draft:
>
> 1. `flag` covers real CLIs; reach for `cobra` only when you outgrow it.
> 2. One `flag.FlagSet` per subcommand — that's the whole "framework".
> 3. `encoding/json` only sees **exported** fields; struct tags control the wire shape.
> 4. Prefer `Encoder`/`Decoder` for streams and files; `Marshal` for in-memory bytes.
> 5. Every JSON and parse call returns an `error` — handle it, no silent coercion.

---

## 🧭 Navigation

| Direction    | Link                                                   |
| ------------ | ------------------------------------------------------ |
| **Previous** | [← Chapter 13: Testing](../13-testing/README.md)       |
| **Next**     | `gorg` — File Organizer CLI _(milestone, not started)_ |
| **Overview** | [Part 1 — Foundation](../README.md)                    |
