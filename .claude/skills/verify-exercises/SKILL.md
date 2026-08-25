---
name: verify-exercises
description: Verify a justgo Chapter's exercise solutions — confirm each is genuinely solved (t.Skip removed, stub gone, tests green, idiomatic Go), then teach the concept and give the exact solution. Use when the user wants to check, verify, review, or grade the exercises in a Chapter, or asks whether a Chapter's exercises are done.
---

# Verify justgo exercise solutions

Verify every exercise in a Chapter's `exercises/` folder is genuinely solved, then
**teach before you solve**: for each one, explain the concept it drills and any
non-idiomatic bits, *then* give the exact idiomatic solution. This is a learning
repo — the review is a lesson, not a patch.

The conventions behind an exercise (table-driven `_test.go`, the `t.Skip` pattern,
exercise dirs excluded from `golangci-lint`) live in [`CLAUDE.md`](../../CLAUDE.md);
read it as the source of truth. This skill caches only the traps that trip
verification.

## Steps

1. **Locate.** Resolve the Chapter folder from the user's reference (number or
   slug) under `content/part-N-*/NN-*/` and list `exercises/*.go`. Pair each
   `X.go` with its `X_test.go`. If the folder has no `exercises/`, the Chapter is
   theory-only — say so and stop.

2. **Confirm solved — every exercise, no exceptions.** For each pair:
   - The `_test.go` has **no active `t.Skip`** (commented-out is fine). A live
     `t.Skip` means the exercise is unsolved and its green test is a false pass.
   - The `X.go` has **no leftover stub**: no `// TODO`, no placeholder return
     (`return "", false // TODO`, `return nil // TODO`, etc.).

3. **Run the checks.** From the repo root:
   - `go test -race ./content/part-N-*/NN-*/...` — must pass with the skips gone.
   - `gofumpt -l content/part-N-*/NN-*/exercises/` — must print nothing.
   - `go vet ./content/part-N-*/NN-*/...`.

   Do **not** run `golangci-lint` on `exercises/` — those dirs are excluded on
   purpose; its stub-oriented warnings are noise here.

4. **Review for idiomatic Go.** Read each solution against its `TODO` hint and the
   Chapter `README.md` concept. Look for: receiver consistency (a type with any
   pointer-receiver method takes pointer receivers everywhere), comma-ok reads,
   error-wrapping style, zero-value correctness, unnecessary allocation.

5. **Teach, then solve.** Report per exercise:
   - **Concept** — the one thing this exercise drills (one or two sentences).
   - **Verdict** — solved / unsolved / non-idiomatic, with the evidence from
     steps 2–4.
   - **Guide** — if unsolved or non-idiomatic, walk through the reasoning so the
     learner sees *why*, before the answer.
   - **Exact solution** — the idiomatic implementation, in a code block.

## Done when

Every exercise in the Chapter has a verdict backed by a passing (or failing, and
named) test, and every unsolved or non-idiomatic one has a guide **and** an exact
solution. A Chapter with a live `t.Skip` anywhere is not done.
