# just as the task runner, not Makefile

Go has no `npm scripts`; the ecosystem's usual answer is a **Makefile** (or a
Taskfile). This repo uses **[just](https://github.com/casey/just)** instead — a
`justfile` at the repo root, with `just fmt` / `just check` / `just test` / `just
setup`, pinned through `mise` alongside the rest of the toolchain.

Why `just` over `make`:

- **Made for task running, not builds.** `make` is a build system: recipes are tied
  to file targets and timestamp checks, which we don't want for "run the tests". A
  target that shares a name with a real file (`test/`, `docs/`) silently misbehaves
  unless marked `.PHONY`. `just` has no file-target model — every recipe is just a
  named command.
- **No tab-vs-space footgun.** `make` requires literal tabs to indent recipes; a
  space-indented recipe fails cryptically. `just` doesn't care.
- **Readable recipes and shebang blocks.** `just` runs multi-line recipes as real
  shell scripts (`#!/usr/bin/env bash`), which `check` uses for a clean `set -euo
  pipefail` pipeline.
- **One tool, pinned.** `just` installs via `mise` like go, gofumpt, and
  golangci-lint — no separate system dependency ([ADR-0003](./0003-node-free-docs-toolchain.md)).

Chapter 02 still _teaches_ Makefile, because Go code you read in the wild overwhelmingly
ships one — you should recognize it. But for this repo, the canonical entry points are
the `just` recipes. Don't add a parallel `Makefile`.
