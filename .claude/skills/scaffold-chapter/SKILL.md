---
name: scaffold-chapter
description: Scaffold a new Chapter (or Part) folder in the justgo learning repo, mirroring the chapter-03 mould — README + QUESTIONS, plus examples/ and exercises/ where the chapter has code. Use when adding, starting, or stubbing a new chapter or part in justgo.
---

# Scaffold a justgo Chapter

Create a Chapter that matches this repo's convention exactly. **Chapter 03
(`content/part-1-foundation/03-types-and-variables/`) is the mould — open it and
mirror its shape.** The rules behind the mould live in
[`CLAUDE.md`](../../../CLAUDE.md), [`CONTEXT.md`](../../../CONTEXT.md), and
`docs/adr/`; the plan (numbers, slugs, key topics per chapter) lives in
[`content/ROADMAP.md`](../../../content/ROADMAP.md). Follow those — don't restate them.

## Inputs to confirm

1. **Number + slug** — two-digit global number (01–42) and kebab title, e.g. `05` +
   `functions` → folder `05-functions`. Take both from the ROADMAP.
2. **Which Part** — `content/part-N-slug/`. If the Part folder doesn't exist yet,
   create it plus its `README.md` chapter-table index.
3. **Has code?** — history/philosophy-style chapters skip `examples/` and
   `exercises/`; concept chapters have both.

## Steps

1. Create `content/part-N-slug/NN-slug/`.
2. **README.md** — the theory file. Two paths, depending on whether a local
   reference draft exists (the `.docs/` folder is git-ignored and personal):
   - **`.docs/NN-slug.md` exists** → derive the README from it, but **verify every
     claim against current Go first** and modernize stale idioms before pasting.
     This is the [`/grill-with-docs`](../grill-with-docs/) flow.
   - **No reference** → build a detailed README **outline**: section headings drawn
     from the chapter's ROADMAP key-topics row, each with a short `TODO`. Leave the
     teaching prose to the owner (CLAUDE.md) — scaffold the skeleton, don't fabricate
     the lesson.
3. **QUESTIONS.md** — interview-style recall questions with foldable `<details>`
   answers, mirroring ch03. Write as many as the chapter's concepts warrant.
4. **If the chapter has code** — mirror ch03's file shapes, and let the chapter's
   scope set the count (a narrow chapter may need one of each; a broad one several):
   - `examples/<name>/main.go` — runnable `package main`, one concept per folder,
     verified with `go run .`. Stays in the root Go module (no new `go.mod`).
   - `exercises/<name>.go` — a stub the learner completes, paired with a table-driven
     `<name>_test.go` that has `t.Skip(...)` at the top, so `go test ./...` stays
     green until the learner removes the Skip and implements it.
5. **Wire it in** — update the Part `README.md` chapter table (status → drafted +
   link the title) and the new README's nav footer (prev/next), same as ch03/ch04.
6. **Verify** — run `just fmt`, then `just check`. Done when `just check` is green.

## Guardrails

- A **Milestone project** is not a Chapter: it gets its own `go.mod` registered in
  `go.work`. Don't scaffold one with this flow — this flow only touches the root
  Go module.
- Terminology and the Node-free rule are set in CLAUDE.md / CONTEXT.md — the one that
  bites here: never write "module" for a Chapter.
