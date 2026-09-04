# HANDOFF — Part 1 review + `gorg` milestone

> Snapshot for continuing in a fresh session (owner is switching machines). Created
> 2026-09-04. Delete this file once the follow-up doc edits below are merged.

## What this session did

A `/grill-with-docs` session that (1) verified Part 1 Foundation for gaps and
(2) chose and specified the Part 1 milestone project.

**All decisions are recorded in the new plan file — read it first:**
[`content/part-1-foundation/gorg/PLAN.md`](./content/part-1-foundation/gorg/PLAN.md)
(see its §12 "Decisions locked").

### Decisions locked (summary — full context in PLAN.md)

1. **Part 1 chapters 01–13 are complete** as a language syllabus. No missing concept.
2. Three project-support gaps found and resolved by decision:
   - ➕ Add **Chapter 14 — Building CLIs & JSON** (`flag`/subcommands + `encoding/json`).
   - 📈 Deepen **Chapter 03** with a strings/bytes/runes section (UTF-8, `strings.Builder`, `strconv`).
   - 🔢 Renumber Part 2+ in `ROADMAP.md` (14→15 … 42→43). Cheap now — those folders
     don't exist on disk yet, it's a pure roadmap-table edit.
3. **Milestone project = `gorg`** (file organizer CLI), **replacing `gitm`**. Reasons:
   higher daily-use for the owner (Arch), and it forces interfaces + generics (which
   `gitm` under-exercised). `httpping` from research was rejected (needs `net/http` →
   Part 2).
4. **Structure:** build `gorg` at the END of Part 1 as its own Go module (per
   ADR-0002). Chapters 03–13 are NOT rewritten to build it incrementally — the
   "grows over time" feel lives in `gorg`'s own Phase 0→5 git history instead.
5. **Part 1 = 100% standard library.** `flag` not cobra; JSON not TOML. Cobra, TOML,
   `fsnotify` watch daemon, bubbletea TUI, OSS release are explicit **post-Part-1**
   evolutions (PLAN.md §9/§11).
6. Doc language for committed content = **English**. Project name = **`gorg`**
   (go + organize; avoids the generic `organize` / Python `organize-tool` clash).

## Immediate next step (was awaiting owner's go-ahead)

Propagate the `gitm → gorg` swap + the ch14/ch03 changes into the three existing
committed docs that still say `gitm` / still show 13 chapters:

- `content/ROADMAP.md` — Part 1 milestone section, timeline table, **Part 2+ chapter
  renumber**.
- `content/part-1-foundation/README.md` — chapter table (add row 14) + milestone line.
- `CONTEXT.md` — the _Milestone project_ definition lists `gitm` as an example → swap
  to `gorg`.

Owner may want a one-line "alternative project idea" note preserving `gitm` — confirm.

## Then / later

- Scaffold Chapter 14 and the Chapter 03 strings section (there is a `scaffold-chapter`
  skill; chapter mould = chapter 03).
- Scaffold the `gorg` module (`go.mod` + `cmd/gorg` + `internal/*`, register in
  `go.work`) when ready to start building — follow PLAN.md's Phase 0.
- Owner authors the teaching prose (README theory); the agent scaffolds
  structure/code/exercises only, per CLAUDE.md — do NOT write chapter teaching content
  unprompted.

## Notes / state

- Owner's research inputs are in `temp/gemini.md` and `temp/gpt.md` (untracked scratch;
  `gpt.md` is the basis of the `gorg` design). Decide whether to keep or delete `temp/`
  before/at commit — probably NOT committed.
- Repo is otherwise clean. Toolchain: `just fmt`, `just check`, `just test`
  (node-free — dprint for docs, per ADR-0003). PLAN.md is already dprint-formatted.
- Glossary discipline (CLAUDE.md / CONTEXT.md): "module" = Go module only; content
  units are Part / Chapter; deliverables are Milestone projects.

## Suggested skills for the next session

- `scaffold-chapter` — to stub Chapter 14 and (structurally) touch Chapter 03.
- `git-workflow` — for staging/committing the doc changes as atomic conventional commits.
- `grill-with-docs` — if the owner wants to keep stress-testing the ROADMAP renumber
  or the `gorg` phase plan before implementing.
