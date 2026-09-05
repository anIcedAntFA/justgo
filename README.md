# justgo

My hands-on path from JS/TS frontend to Go backend — theory, runnable code,
exercises, and recall questions, all in one repo.

## Where to start

- **[The roadmap →](./content/ROADMAP.md)** — the full 42-chapter, 4-project plan.
- **[Part 1 · Foundation →](./content/part-1-foundation/README.md)** — start here.
- **[CONTEXT.md](./CONTEXT.md)** — the words this repo uses (Part, Chapter, Milestone
  project, Exercise, Question) and what they mean.
- **[docs/adr/](./docs/adr/)** — why the repo is shaped the way it is.
- **[docs/tooling.md](./docs/tooling.md)** — install and run the toolchain.

## How it's organized

```
content/                  the learning material (this is the product)
  ROADMAP.md              single source of truth for the plan
  part-1-foundation/      Part → Chapter folders + the gorg project
docs/                     ADRs + tooling how-to (how the repo runs)
CONTEXT.md                glossary
```

Each **Chapter** is a folder `NN-slug/` with a `README.md` (theory), `QUESTIONS.md`
(recall), and — where there's code — `examples/` (run with `go run .`) and
`exercises/` (test-driven; `go test` until green). Chapter 03 is the complete mould.

## Quick start

```sh
mise install     # or: just setup   (installs Go + tools, no Node needed)
just             # list all tasks
just fmt         # format Go + docs
just check       # everything CI runs
just test        # run the test suite
```

New here? Read [`CLAUDE.md`](./CLAUDE.md) for the conventions before adding content.
