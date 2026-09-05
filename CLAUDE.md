# CLAUDE.md

Guidance for working in **justgo** — a personal Go learning repo (theory + runnable
code + exercises + questions). Read this before adding or changing content.

## What this repo is

A structured path from JS/TS frontend to Go backend. The **content is the product**;
the code exists to be learned from, not shipped. The full plan lives in
[`content/ROADMAP.md`](./content/ROADMAP.md) — treat it as the single source of truth
for what chapters/projects exist and their order.

## Language — use the glossary, always

[`CONTEXT.md`](./CONTEXT.md) is the glossary. The non-negotiable one:

> **"module" means exactly one thing: a Go module (`go.mod`).** It is NEVER a synonym
> for a unit of content. Content units are **Part** (the 4 stages) and **Chapter**
> (the numbered 01–42 units). Deliverables are **Milestone projects**. In-chapter
> coding tasks are **Exercises**; recall prompts are **Questions**.

If you catch yourself or the user saying "module" for a lesson, stop and disambiguate.

## Repository shape (see docs/adr/)

- **Monorepo**, no submodules — [ADR-0001](./docs/adr/0001-single-repo-for-theory-and-exercises.md).
- **Multi-module `go.work`** — [ADR-0002](./docs/adr/0002-multi-module-workspace.md):
  all Chapter code lives in the **root Go module**; **each Milestone project is its
  own module** (own `go.mod`), registered in `go.work`.
- **Node-free toolchain** — [ADR-0003](./docs/adr/0003-node-free-docs-toolchain.md):
  format docs with `dprint`, never introduce `package.json` / `node_modules` / pnpm.

## Chapter folder convention (the mould = chapter 03)

Every Chapter is a folder `content/part-N-slug/NN-kebab-slug/` (two-digit global
number 01–42). Inside:

```
NN-slug/
├── README.md       ← the theory. GitHub renders it as the folder index. Required.
├── QUESTIONS.md    ← interview-style recall questions, foldable <details> answers.
├── REFLECTIONS.md  ← worked answers to the Chapter's open-ended Reflection prompts
│                     (TL;DR + foldable deep dive). Optional — see ADR-0004.
├── examples/       ← runnable `package main` demos the README references (go run .)
├── exercises/      ← a package with stubs + _test.go; make `go test` pass.
└── archives/       ← optional personal scratch: a `package main` playground + raw notes.
```

- **`examples/`** are `package main`, complete and runnable.
- **`exercises/`** are stubs the learner completes. Ship each exercise with a
  table-driven `_test.go` and a `t.Skip(...)` at the top so `go test ./...` stays
  green until the learner removes the Skip and implements it.
- Chapters with no code (e.g. 01 History) omit `examples/` and `exercises/`.
- **`archives/`** (optional) is the owner's personal scratch — a `package main`
  playground plus raw study notes (`.md`) parked per chapter. It's committed but is
  **not** taught content; scaffolding never creates it (the owner adds it while
  studying). Like `exercises/` it's excluded from `golangci-lint`, but it's still
  compiled, vetted, and tested by `go ./...`, so keep it building.
- Copy chapter 03 verbatim as the starting skeleton for a new chapter.

## Content authorship

Chapter theory follows an **"agent drafts → owner refines"** model. If a
`.docs/NN-slug.md` draft exists, derive the `README.md` from it (verified against
current Go). If not, the **agent writes a complete first draft** as an expert Go
engineer — real prose, not `TODO`s — verified against primary Go sources
(pkg.go.dev, spec, go.dev), with a cited notes file in `.docs/`. The **owner owns
the final voice** and may rewrite freely; once a chapter is owner-edited, treat it as
owner-owned and don't rewrite it unprompted (draft only what's missing). See
[ADR-0007](./docs/adr/0007-agent-drafts-chapter-theory.md). Not every chapter is
drafted yet — the Part `README.md` index shows status.

## Adding things

- **New chapter** → mirror the chapter-03 folder; update the Part `README.md` table
  and, if the plan changes, `content/ROADMAP.md`. (A `scaffold-chapter` skill exists.)
- **New Milestone project** → new folder in its Part with its own `go.mod`, then add
  `use ./content/part-N-slug/<project>` to `go.work`.

## Commands

```sh
just fmt      # gofumpt + dprint, in place
just check    # what CI runs: gofumpt check, go vet, golangci-lint, dprint, test, secrets
just test     # go test -race ./...
```

Always run `just fmt` (or at least `gofumpt -w`) on Go you touch. Exercise and archive
dirs are excluded from `golangci-lint` on purpose — don't "fix" their intentional stubs
or scratch. `examples/` **are** linted, so keep demo code lint-clean.
