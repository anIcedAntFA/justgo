# Chapter 01 — Reflection Answers

My worked answers to **Exercise 2: Reflection Questions** in
[`README.md`](./README.md). Each answer leads with a **TL;DR** you can recall from
memory, then a foldable **Full answer** with tables and examples.

> These are study notes, not canonical theory — the teaching content stays in
> `README.md`. Write your own answer first, _then_ unfold to compare.

---

## 1. Name 3 features from TypeScript that Go deliberately excluded — and why.

**TL;DR**

- **Exceptions (`try/catch`)** → Go returns **errors as values** so every failure is
  handled explicitly and can't slip through.
- **Classes / inheritance** → Go prefers **composition** (struct embedding +
  small interfaces) over class hierarchies.
- **Ternary operator `? :`** → Go keeps **one obvious way**: `if/else`.

<details>
<summary>Full answer</summary>

Each omission is a design decision, not an oversight — the point is to keep the
language small so features don't multiply complexity ("less is exponentially more").

| Excluded TS feature      | TS style                   | Go's replacement                                  | Why Go said no                                                                            |
| ------------------------ | -------------------------- | ------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| Exceptions (`try/catch`) | `throw` / `catch (e)`      | Multi-value return `(T, error)` + `if err != nil` | Error paths stay visible next to the happy path; the compiler makes you acknowledge them. |
| Classes & inheritance    | `class Dog extends Animal` | `struct` embedding + implicit interfaces          | Composition is flatter and avoids the "fragile base class" problem of deep hierarchies.   |
| Ternary operator `? :`   | `cond ? a : b`             | `if/else`                                         | "Clear is better than clever" — one readable form, no nested-ternary golf.                |

```go
// Composition, not inheritance: Dog "has an" Animal.
type Animal struct{ Name string }
type Dog struct {
    Animal // embedded — Animal's fields/methods are promoted onto Dog
}
```

Other valid picks from the same family: **generics** (deliberately waited 13 years,
until 1.18), **enums** (const + `iota` instead), **decorators/annotations** (explicit
code over magic), **default parameter values** (functional-options pattern instead),
and **method overloading** (one name, one behavior).

</details>

---

## 2. How does Go's error handling philosophy differ from JavaScript's? What are the tradeoffs?

**TL;DR**

- **Go**: errors are ordinary **values returned** from a function and handled right
  there (`if err != nil { ... }`).
- **JS**: errors are **exceptions** thrown and caught elsewhere (`throw` /
  `try...catch`), propagating implicitly up the stack.
- **Tradeoff**: Go buys **explicitness & predictability** at the cost of verbosity;
  JS buys a **clean happy path** at the cost of invisible control flow.

<details>
<summary>Full answer</summary>

In Go, a function's signature _tells you it can fail_ (`(T, error)`), and you decide
what to do at every step. In JS, any function can `throw` without that showing up in
its type, so failure is easy to forget until it blows up at runtime.

```go
// Go: the failure path is right next to the happy path, always visible.
data, err := fetchUser(id)
if err != nil {
    return nil, fmt.Errorf("fetching user %d: %w", id, err)
}
```

```javascript
// JS: which of the two calls threw? The type didn't warn you.
try {
  const data = await fetchUser(id)
  return await fetchProfile(data.profileId)
} catch (err) {
  console.error(err) // 🤷 where did this come from?
}
```

| Dimension       | Go (errors as values)                 | JS (exceptions)                                 |
| --------------- | ------------------------------------- | ----------------------------------------------- |
| Visibility      | In the signature — you can't miss it  | Hidden; a `throw` isn't in the return type      |
| Control flow    | Explicit, local, predictable          | Non-local jump up the stack                     |
| Easy to ignore? | Hard — unused values nag you          | Easy — a forgotten `catch` swallows it silently |
| Verbosity       | Higher (`if err != nil` repeats)      | Lower — happy path reads cleanly                |
| Best when       | Reliability matters (backends, infra) | Quick scripts, UI flows                         |

**Net**: explicitness/predictability vs conciseness/implicit propagation. Go trades
a few extra lines for the guarantee that errors can't silently slip through.

</details>

---

## 3. What's the difference between concurrency and parallelism? Which does JavaScript support? Which does Go support?

**TL;DR**

- **Concurrency** = _structuring_ a program to **deal with** many tasks at once
  (interleaving). **Parallelism** = _actually executing_ many tasks
  **simultaneously** (needs multiple cores). Rob Pike: _"Concurrency is not
  parallelism."_
- **JavaScript**: concurrency by default (single-threaded event loop); parallelism
  only via **Web Workers / worker threads**.
- **Go**: **both** — goroutines express concurrency, and the runtime schedules them
  across CPU cores (controlled by `GOMAXPROCS`) for real parallelism.

<details>
<summary>Full answer</summary>

Concurrency is about **structure and coordination** — breaking work into independent
tasks that _can_ make progress in overlapping time windows. Parallelism is about
**execution resources** — running tasks at the literal same instant, which requires
more than one core. You can have concurrency on a single core (interleaving);
parallelism needs the hardware.

| Aspect              | Concurrency                         | Parallelism                   |
| ------------------- | ----------------------------------- | ----------------------------- |
| Question it answers | _How do I structure many tasks?_    | _How do I run tasks at once?_ |
| Needs many cores?   | No                                  | Yes                           |
| Analogy             | One barista juggling several orders | Several baristas, one each    |

**JavaScript** is single-threaded at heart. The **event loop** gives concurrency:
while an `await`ed network request is pending, the thread runs other work — great for
I/O-bound tasks. True parallelism needs extra threads: **Web Workers** (browser) or
**worker threads** (Node.js).

**Go** builds both into the language. **Goroutines** (~2KB each, millions possible)
express concurrency cheaply; the runtime multiplexes them onto OS threads and, when
`GOMAXPROCS > 1` (default = number of CPU cores), runs them **in parallel** across
cores.

```go
// Concurrency (structure) that the runtime can also run in parallel.
var wg sync.WaitGroup
for _, url := range urls {
    wg.Add(1)
    go func() { defer wg.Done(); fetch(url) }() // each in its own goroutine
}
wg.Wait()
```

**Remember**: JS = concurrency by default, parallelism when you reach for workers.
Go = concurrency in the language, parallelism for free from the runtime.

</details>

---

## 4. Why does Go culture discourage heavy dependency usage? How does this compare to the npm ecosystem?

**TL;DR**

- Go proverb: **"A little copying is better than a little dependency."** A dependency
  is a liability (security, breaking changes, supply-chain, maintenance).
- Go's **standard library is deliberately comprehensive**, so you rarely need
  external packages.
- **npm** embraces many tiny composable packages (`left-pad`, `is-odd`), so even
  simple apps drag in hundreds/thousands of transitive deps.
- **Tradeoff**: npm gives ecosystem reach & speed; Go gives a smaller, more auditable,
  longer-lived dependency tree.

<details>
<summary>Full answer</summary>

Go values simplicity, maintainability, and long-term stability. Adding a dependency
isn't free — it's surface area for supply-chain attacks, unexpected breaking changes,
version churn, and an ongoing maintenance burden. If you need 20 lines of utility,
Go culture says **copy them** rather than take on that liability. A strong standard
library (`net/http`, `encoding/json`, `testing`, `log/slog`) means you can go far
without third parties.

npm optimizes for the opposite: small, sharply-scoped packages you compose. That
enables enormous velocity and reuse, but a "simple" app can end up with a huge
transitive tree — and each node is a potential risk.

| Aspect            | Go                                                    | npm                                     |
| ----------------- | ----------------------------------------------------- | --------------------------------------- |
| Default instinct  | Stdlib first; copy small helpers                      | `npm install` a package for it          |
| Typical tree size | Small, shallow                                        | Large, deep (transitive deps explode)   |
| Guiding proverb   | "A little copying is better than a little dependency" | "Don't reinvent the wheel"              |
| Cost model        | Fewer deps → less risk, more code to own              | More deps → faster, more risk to manage |
| Verification      | `go.sum` cryptographic hashes                         | `package-lock.json`                     |

**Cultural one-liner**: Go leans "use the stdlib or write it yourself unless a
dependency clearly earns its place"; npm leans "compose functionality from many small
packages."

</details>
