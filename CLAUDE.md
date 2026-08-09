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
└── exercises/      ← a package with stubs + _test.go; make `go test` pass.
```

- **`examples/`** are `package main`, complete and runnable.
- **`exercises/`** are stubs the learner completes. Ship each exercise with a
  table-driven `_test.go` and a `t.Skip(...)` at the top so `go test ./...` stays
  green until the learner removes the Skip and implements it.
- Chapters with no code (e.g. 01 History) omit `examples/` and `exercises/`.
- Copy chapter 03 verbatim as the starting skeleton for a new chapter.

## Content authorship

The **theory in each `README.md` is written by the repo owner** (composed with
claude.ai and pasted in). Do **not** fabricate or rewrite chapter teaching content
unless explicitly asked — scaffold the structure, write the code/exercises/questions
skeleton, and leave the prose to the owner. Only chapters up to 03 are drafted;
04–42 are not started yet.

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

Always run `just fmt` (or at least `gofumpt -w`) on Go you touch. Exercise dirs are
excluded from `golangci-lint` on purpose — don't "fix" their intentional stubs.
