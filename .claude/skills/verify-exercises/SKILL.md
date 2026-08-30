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
exercise dirs excluded from `golangci-lint`) live in [`CLAUDE.md`](../../../CLAUDE.md);
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

5. **Teach the reasoning, then solve.** Lead the report with a one-line **Chapter
   verdict summary** — every exercise and its verdict on one line each, so the
   learner can triage before scrolling. Then, per exercise:

   - **Verdict** — solved / unsolved / non-idiomatic, with the evidence from
     steps 2–4. Keep it a crisp standalone line; don't bury it in the Guide.
   - **Guide** — a fixed five-stage walkthrough of *how to think*, not just what
     to type. The stages are always, in order:
     1. **Read the signature & table** — before any code, what do the func/type
        signature and the `_test.go` cases tell you? Inputs, outputs, and the edge
        rows the table quietly asserts (empty, `nil`, zero value, negative).
     2. **Mental model** — the one Go concept and the picture to hold in your head
        (this absorbs the old "Concept": "a struct can't hold itself by value → link
        via pointer").
     3. **Decision points** — the forks the solution hinges on (`nil` head vs not,
        pointer vs value receiver, *why return the head*).
     4. **Steps** — the ordered implementation plan in plain language, **no code**.
     5. **Idiomatic solution** — the exact code, in a code block.
   - **Exact solution** is stage 5 above.

   **Scale the Guide by verdict:**
   - **Unsolved** → all five stages; stage 5 is the exact implementation.
   - **Solved but non-idiomatic** → all five stages; stage 5 is the *corrected*
     implementation.
   - **Solved and idiomatic** → compress to **stages 2 and 3 only**, one line each
     ("your code returns the head because an empty-list append has no node to
     mutate — confirm that was your reasoning"). Drop stages 1, 4, and 5; stage 5
     shrinks to a half-line affirming the learner's existing code, not a fresh block.

   **Format budget** (so a 6-exercise Chapter doesn't become a wall): each stage is
   **1–2 sentences or bullets** — no prose paragraphs; code blocks appear **only** in
   stage 5.

## Done when

Every exercise in the Chapter has a verdict backed by a passing (or failing, and
named) test, and every one has a Guide scaled to its verdict — the full five stages
for anything unsolved or non-idiomatic (with an exact or corrected solution), the
compressed stages 2–3 for anything solved and idiomatic. A Chapter with a live
`t.Skip` anywhere is not done.
