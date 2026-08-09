# Reflection answers live in a per-Chapter REFLECTIONS.md

Some Chapters end with open-ended **Reflection Questions** (currently labelled
"Exercise 2" in Chapter 01's `README.md`) — essay-style prompts the learner answers
in prose, not code. The learner's worked answers go in a dedicated
**`REFLECTIONS.md`** in the Chapter folder, **not** inline in `README.md` and **not**
in `QUESTIONS.md`.

The reasons:

- **`README.md` is owner-authored teaching content** and should stay clean — injecting
  long personal answers (with tables and examples) would bloat the lesson.
- **`QUESTIONS.md` is for recall Questions** (short, interview-style, foldable
  `<details>`). Reflection answers are longer and comparative; mixing them dilutes
  both files.
- A separate file gives room for the intended **TL;DR → full** layering: a recallable
  summary up top, then a foldable deep dive with tables/examples.

Convention for `REFLECTIONS.md`: one `##` heading per question, a **TL;DR** bullet
list visible by default, then a `<details><summary>Full answer</summary>` block. The
`README.md` Exercise 2 section carries a one-line pointer to it. Chapters without
reflection prompts simply omit the file.

**Naming note (a glossary tension to fix later):** these prompts are really
**Questions** by the `CONTEXT.md` glossary (explain, no code) — "Exercise" is a
misnomer inherited from the drafted Chapter 01. We keep the file name `REFLECTIONS.md`
(distinct from the recall `QUESTIONS.md`) and leave the README label as-is for now to
avoid churning already-drafted prose; a future pass may relabel it. `CONTEXT.md` adds
a **Reflection** term to record this.
