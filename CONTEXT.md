# justgo — Context

Glossary for this Go learning repo. Definitions only — the actual teaching content
lives under `content/`, the tooling _how_ lives in `docs/`. Terms here are the ones
easy to conflate; pick the listed word, avoid the rest.

## Content taxonomy

**Part**:
One of the four stages of the learning path — Foundation, Web & APIs, Advanced,
Production. The top grouping; each Part ends with a Milestone project.
_Avoid_: phase, section, module.

**Chapter**:
One numbered unit of theory (01–42) inside a Part. The atomic thing you "study" —
one concept, one folder.
_Avoid_: lesson, unit, module.

**Milestone project** (short: _project_):
A runnable, genuinely usable program built at the end of a Part — `gitm`,
`dropshare`, `goproxy`, `gochat`, `gogate`. Each is its own **Go module**.
_Avoid_: app, exercise.

**Exercise**:
A small coding challenge that lives _inside_ a Chapter to reinforce that Chapter's
concept. Much smaller than a Milestone project.
_Avoid_: task, challenge, kata.

**Question**:
A theory/recall prompt inside a Chapter (interview-style, no code) — "explain the
difference between a zero value and nil". Lives in the Chapter's `QUESTIONS.md`.
An Exercise makes you _write_ code; a Question makes you _explain_.
_Avoid_: quiz, FAQ.

**Reflection**:
An open-ended, essay-style prompt at the end of a Chapter — the learner answers in
prose (compare Go vs JS/TS, explain a tradeoff), longer and more comparative than a
recall Question. Worked answers live in the Chapter's `REFLECTIONS.md`
(TL;DR + foldable deep dive), never inline in `README.md`. See
[ADR-0004](./docs/adr/0004-reflection-answers-in-separate-file.md).
_Note_: Chapter 01 labels these "Exercise 2" in its README — a misnomer; they are
Reflections/Questions, not coding Exercises.
_Avoid_: essay, journal.

**Checkpoint**:
A weekly, verifiable "can you do X yet?" success criterion from the roadmap. A
learning assessment, not a deliverable.
_Avoid_: milestone (that word belongs to Milestone project).

## Flagged ambiguities

**"module" means exactly one thing here: a Go module** (`go.mod`, the unit of
dependency management). It is **never** a synonym for a Part, Chapter, or lesson.
When grouping content, say Part or Chapter. When talking about `go.mod` / `go.work`,
say Go module.

**"project"** unqualified means a **Milestone project**. The whole repository is the
"repo", not "the project".

## Example dialogue

> **Learner:** Which module is the `use` command in?
> **Mentor:** Careful — do you mean which Go module, or which Chapter? The `use`
> command is in the **`gitm` Milestone project**, which is its own Go module under
> `content/part-1-foundation/`. The Chapter that _teaches_ the concepts it uses is
> Chapter 06, Structs & Methods.
> **Learner:** Got it. And the little slice-append thing I did earlier?
> **Mentor:** That's an **Exercise** inside Chapter 10, not a project — it stays in
> the root Go module with the other Chapter code.
