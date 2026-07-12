# Chapter 02 — Questions

Recall questions. Answer before unfolding; re-read [`README.md`](./README.md) if stuck.

<details>
<summary>1. What does <code>go.mod</code> declare, and what is a Go module?</summary>

A Go **module** is a collection of packages versioned together as a unit;
`go.mod` declares the module's import path, the Go version, and its dependencies. It
is the boundary of dependency management — the equivalent of a `package.json` root.

</details>

<details>
<summary>2. What's the difference between <code>gofmt</code> and <code>gofumpt</code>? Why is formatting not optional in Go?</summary>

`gofmt` is the standard formatter shipped with Go; `gofumpt` is a stricter superset
with extra rules. Formatting is not a matter of taste in Go — there is one canonical
style, so diffs stay about logic, not whitespace. This repo uses `gofumpt`.

</details>

<details>
<summary>3. What does <code>go.work</code> do, and why does this repo use one?</summary>

`go.work` defines a **workspace** — it lets several modules on disk resolve against
each other without publishing. This repo keeps Chapter code in the root module and
each Milestone project in its own module; `go.work` stitches them so the editor and
`go` commands see everything at once. (See ADR-0002.)

</details>
