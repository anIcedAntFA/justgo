# Chapter 11 — Questions

Recall questions. Try to answer out loud before unfolding. If you can't, re-read the
matching section of [`README.md`](./README.md).

<details>
<summary>1. What is the relationship between a <em>file</em>, a <em>package</em>, and a <em>module</em> in Go?</summary>

Every `.go` **file** declares exactly one **package** on its first line, and every
file in a directory must declare the same package — so a package maps to a directory.
A **module** is a collection of packages versioned together, defined by a single
`go.mod` at its root. In short: file → package (directory) → module (project).

</details>

<details>
<summary>2. How does Go decide whether an identifier is exported (public) or unexported (private)? Give the rule for a struct field.</summary>

**Capitalization of the first letter** — there is no `export` keyword. An uppercase
first letter (`Name`, `NewUser`, `MaxUsers`) is exported and visible to other
packages; a lowercase first letter (`token`, `validate`, `defaultTimeout`) is
unexported and private to its own package. The rule applies to everything, including
struct fields: `Name string` is readable/writable from outside, `token string` is
not.

</details>

<details>
<summary>3. What does <code>package main</code> mean, and how does it differ from every other package?</summary>

`package main` combined with a `func main()` is the entry point of an **executable** —
`go build` turns it into a binary. Every other package is a **library**: importable
code that can't run on its own. A project can have several `main` packages (e.g. under
`cmd/server`, `cmd/cli`), each producing its own binary.

</details>

<details>
<summary>4. Your module is <code>github.com/you/app</code>. How do you import your own <code>store</code> sub-package? Can you write <code>import "./store"</code>?</summary>

You import it by its **full module path**: `import "github.com/you/app/store"`. Go has
no relative imports — `"./store"` and `"../utils"` are not valid. The module path from
`go.mod` is the prefix for every package inside the module, and the tooling
(gopls/goimports) autocompletes these paths.

</details>

<details>
<summary>5. What are <code>go.mod</code> and <code>go.sum</code> for, and which one do you edit by hand? What's the JS analogy?</summary>

`go.mod` declares the module path, the Go version, and the `require` list of
dependencies. `go.sum` holds cryptographic hashes of every dependency (and their
`go.mod` files) so the toolchain can verify downloads haven't been tampered with —
like `package-lock.json` but with real integrity checking, not just version pinning.
Edit **neither by hand**: `go get` and `go mod tidy` maintain them. Commit both.

</details>

<details>
<summary>6. What does <code>go mod tidy</code> do, and when should you run it?</summary>

It syncs `go.mod`/`go.sum` with what your code actually imports: **adds** any
dependency your code uses but `go.mod` is missing, **removes** any dependency no
longer imported, and updates `go.sum`. Run it whenever you add or remove imports.
It's `npm install` + `npm prune` combined, driven by your source rather than manual
edits.

</details>

<details>
<summary>7. What is special about a directory named <code>internal/</code>, and why is it stronger than capitalization?</summary>

Packages under an `internal/` directory can only be imported by code rooted at
`internal`'s **parent** — the compiler physically rejects imports from anywhere else,
including other modules. Capitalization controls visibility **per package**;
`internal/` controls it across the whole **module boundary**, so it's the tool for
"private to this project" code like data layers and implementation details.

</details>

<details>
<summary>8. Package A imports package B. Can B import A? What is Go's rule, and how do you fix a violation?</summary>

**No** — Go forbids circular imports, directly or transitively. An A↔B cycle is a
compile error. Fixes: (1) extract the shared types into a third package both depend
on, (2) invert the dependency with an interface (define the interface in the
consumer, let the other package implement it), or (3) merge the two packages if
they're genuinely that coupled. The ban forces you to think about dependency
direction.

</details>

<details>
<summary>9. Why is <code>github.com/labstack/echo/v4</code> — with <code>/v4</code> in the path — imported that way?</summary>

**Semantic Import Versioning**: for major versions ≥ 2, the major version becomes part
of the import path (`/v2`, `/v4`, …). This lets a build import two major versions of
the same library at once when a dependency conflict requires it, and it makes the
major version explicit at every use site. Major version 0 and 1 have no suffix.

</details>

<details>
<summary>10. What is "stutter" in package/member naming, and how do you avoid it? Give an example.</summary>

Stutter is repeating the package name inside a member name, so the qualified use reads
redundantly: `user.UserService`, `http.HTTPClient`. Because the package name already
qualifies every member, drop the repetition: `user.Service`, `http.Client`,
`json.Decoder`. Constructors follow the same rule — prefer `user.New()` over
`user.NewUser()`.

</details>

<details>
<summary>11. What does a blank import (<code>_ "github.com/mattn/go-sqlite3"</code>) accomplish?</summary>

It imports the package **only for its `init()` side effects** without binding its name
— you never call the package directly. The classic case is registering a database
driver or an image-format decoder with a central registry (e.g. the `sqlite3` driver
registers itself with `database/sql`). The `_` avoids the "imported and not used"
compile error you'd otherwise get.

</details>

<details>
<summary>12. Why are generic packages like <code>utils</code>, <code>common</code>, or <code>helpers</code> discouraged?</summary>

They're named by nothing — the name doesn't say what's inside, so they become dumping
grounds that accumulate unrelated code and grow into a tangle (and often the seed of
circular-import problems). Name packages by what they **do**: `dateformat`,
`jsonutil`, `auth`, `validate`. A good package name is short, lowercase, single-word,
and descriptive.

</details>
