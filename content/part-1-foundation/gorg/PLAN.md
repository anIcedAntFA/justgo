# `gorg` — Milestone Project Plan (Part 1 · Foundation)

> **`gorg`** = _go + organize_. A fast, safe, testable **file organizer CLI** built
> from the Go **standard library only**. This is the capstone of Part 1: every
> chapter's concept lands here for a real reason, not a contrived demo.

> [!NOTE]
> This is a **plan / specification**, not an implementation guide. It defines _what_
> we build, _why_, and _in what order_ — mapping each phase to the Part 1 concepts and
> the standard-library packages it exercises. It intentionally contains **no solution
> source code**. You write the code; this document is the map.

---

## TL;DR

`gorg <dir>` scans a messy directory (Downloads, screenshots, a project dump), sorts
files into category folders (`Images/`, `Documents/`, …) by configurable rules, and
does it **safely** — preview with `--dry-run`, resolve name clashes deliberately, and
reverse the whole run with `gorg undo`. Pure stdlib. Fully unit-testable via a fake
filesystem. Built incrementally in 5 phases that track chapters 03 → 14.

---

## 1. What we build (scope)

**In scope (Part 1, stdlib only):**

- Scan a directory (flat or `--recursive`), collect file metadata.
- Classify each file into a category via rules (extension-based by default,
  overridable by a JSON config).
- Build a **plan** of move operations, then execute it.
- Safety features: `--dry-run` (preview, zero mutations), conflict resolution
  (overwrite / skip / rename), and an **undo** journal that reverts the last run.
- A JSON config file for custom rules and destination folders.
- A `stats` report: counts and total size per category.
- (Stretch, still stdlib) duplicate detection by content hash.

**Explicitly NOT in scope for Part 1** (these come _after_ Part 1 as the "real tool"
evolution — see §9):

- ❌ No database, no HTTP server, no network.
- ❌ No concurrency (no goroutines/channels) — a watch daemon is a Part 3 upgrade.
- ❌ No TUI — a bubbletea front-end is a post-Part-1 bonus.
- ❌ No third-party packages — `flag`, not cobra; JSON, not TOML/YAML. Adding these
  later is the point at which Chapter 11 (dependency management) becomes real.

> **Why the constraint?** The goal of Part 1 is _"think in Go"_. A pile of features
> makes the project bigger, not the learning deeper. CLI + filesystem + config +
> dry-run + undo + tests is the sweet spot that exercises the whole language surface
> while staying coherent.

## 2. Usage / value model (the "why you'll actually run it")

| Who                 | When                                                          | Value                                            |
| ------------------- | ------------------------------------------------------------- | ------------------------------------------------ |
| You, daily, on Arch | `~/Downloads` and `~/Pictures/Screenshots` fill up with noise | One command tidies them, reversibly              |
| You, on a project   | A folder accumulates exports, PDFs, assets                    | Rule-based sort into sane subfolders             |
| Future you          | Publish as OSS                                                | A small, safe, well-tested tool others can trust |

The defining virtue is **safety**: a file organizer that can silently overwrite or
lose data is worse than useless. `--dry-run` + explicit conflict strategy + `undo` are
not extras — they are the product.

## 3. Domain model (concepts, not code)

These are the nouns the program reasons about. Each becomes a type; responsibilities
listed, no implementation shown.

| Concept              | Responsibility                                                                                                                                          |
| -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **File**             | An observed file: path, name, extension, size, hidden?, (later) content hash                                                                            |
| **Category**         | A named bucket (`Images`, `Documents`, `Videos`, `Music`, `Archives`, `Other`, or custom) — a defined enum-like string type                             |
| **Rule**             | Maps a file characteristic (extension today; mime/size later) to a Category                                                                             |
| **Classifier**       | _Behaviour_ that decides a file's Category — an **interface**; default impl is extension-based, others plug in later                                    |
| **FileSystem**       | _Behaviour_ for `ReadDir` / `Move` / `Mkdir` / `Stat` — an **interface** so the real OS and a fake can be swapped (this is what makes the app testable) |
| **Operation**        | A single intended change: move `Source` → `Destination`                                                                                                 |
| **Plan**             | An ordered set of Operations produced _before_ any mutation (enables `--dry-run`)                                                                       |
| **ConflictStrategy** | _Behaviour_ that resolves a destination-already-exists clash: overwrite / skip / rename                                                                 |
| **Transaction**      | The executed Operations of one run, persisted as the **undo journal**                                                                                   |
| **Config**           | User-defined rules + destination folder names, loaded from JSON                                                                                         |
| **Report**           | Aggregated stats (counts, sizes) per Category for `gorg stats`                                                                                          |

The two **interface** seams — `Classifier` and `FileSystem` — are the heart of the
design and the reason the whole thing is unit-testable. Treat them as the primary
learning target of the project.

## 4. Command surface (input) & output

```text
gorg <dir>                      organize a directory (default command)
gorg <dir> --dry-run            preview operations; change nothing
gorg <dir> --recursive          descend into subdirectories
gorg <dir> --interactive        prompt on each conflict
gorg <dir> --config <path>      use a specific rules file
gorg <dir> --duplicates         (stretch) detect & report duplicate content
gorg stats <dir>                report counts/sizes per category
gorg undo                       revert the most recent run
gorg rules                      print the effective ruleset
```

**Example — `--dry-run` output (no mutation):**

```text
Plan for ~/Downloads (dry-run, 4 operations):
  IMG_1234.jpg   → Images/IMG_1234.jpg
  report.pdf     → Documents/report.pdf
  song.mp3       → Music/song.mp3
  archive.zip    → Archives/archive.zip
Nothing was moved. Re-run without --dry-run to apply.
```

**Example — `stats` output:**

```text
~/Downloads
  Images       124 files    412 MB
  Documents     32 files     88 MB
  Videos         8 files    1.2 GB
  Archives      15 files    340 MB
  Other          6 files      4 MB
```

## 5. Functional requirements

- **FR-1** Scan a directory and produce a list of Files with metadata.
- **FR-2** `--recursive` descends into subdirectories; without it, only the top level.
- **FR-3** Classify each File into exactly one Category using the effective ruleset.
- **FR-4** Default rules are built in; a JSON config file overrides/extends them.
- **FR-5** Build a Plan of Operations _before_ mutating anything.
- **FR-6** `--dry-run` prints the Plan and performs zero filesystem changes.
- **FR-7** Execute the Plan: create destination folders as needed, then move files.
- **FR-8** On a destination clash, resolve via the selected ConflictStrategy
  (default: skip; `--interactive` prompts; config may set overwrite/rename).
- **FR-9** Every executed run writes a Transaction (undo journal) to disk.
- **FR-10** `gorg undo` reverses the most recent Transaction.
- **FR-11** `gorg stats` reports counts and total size per Category.
- **FR-12** _(Stretch)_ `--duplicates` detects same-content files by hash.

## 6. Non-functional requirements

- **NFR-1 Safety first** — never lose data. No silent overwrite; `--dry-run` must be
  a true no-op; a partial failure must leave a recoverable journal.
- **NFR-2 Testable** — all filesystem access goes through the `FileSystem` interface,
  so the core logic tests against a fake with **no real disk I/O** (`t.TempDir()` used
  only for the thin real-OS adapter tests).
- **NFR-3 Deterministic** — same input + same rules ⇒ same Plan.
- **NFR-4 Idempotent-ish** — re-running on an already-organized dir does nothing harmful.
- **NFR-5 Zero external dependencies** in Part 1 — stdlib only.
- **NFR-6 Clear errors** — wrap with context (`%w`), define sentinel/domain errors,
  aggregate per-file failures instead of aborting the whole run on the first error.
- **NFR-7 Portable** — build to a single static binary; behave under Linux paths.
- **NFR-8 Reasonable coverage** — ≥ 80% on the core packages (classifier, planner,
  organizer).

## 7. Proposed module structure

Its own Go module (per ADR-0002), registered in `go.work`. **Do not create all of this
on day one** — it emerges as complexity appears (that _is_ the Chapter 11 lesson).

```text
content/part-1-foundation/gorg/
├── go.mod
├── README.md
├── PLAN.md                     ← this file
├── cmd/gorg/                   ← main: flag parsing, subcommand dispatch, wiring
└── internal/
    ├── file/                   ← File type + metadata helpers
    ├── classifier/             ← Classifier interface + extension impl
    ├── planner/                ← Plan + Operation, conflict detection, dry-run
    ├── organizer/              ← orchestrates scan → classify → plan → execute
    ├── filesystem/             ← FileSystem interface, OS impl, fake impl (tests)
    ├── history/                ← Transaction journal + undo
    └── config/                 ← JSON rules load/merge
```

## 8. Standard-library map (the only "libraries" in Part 1)

| Package           | Used for                                                                       |
| ----------------- | ------------------------------------------------------------------------------ |
| `flag`            | Flags and (manual) subcommand dispatch                                         |
| `os`              | `Open`, `Rename`/move, `MkdirAll`, `Stat`, `Args`, exit codes, `UserConfigDir` |
| `io` / `io/fs`    | `fs.FileInfo`, `fs.DirEntry`, the `FileSystem` abstraction                     |
| `path/filepath`   | `Join`, `Ext`, `Base`, `WalkDir` (recursive scan), clean paths                 |
| `encoding/json`   | Load config; read/write the undo journal                                       |
| `strings`         | Extension lowercasing, matching, building output                               |
| `strconv`         | Human-readable sizes / parsing                                                 |
| `fmt`             | Output, and error wrapping with `%w`                                           |
| `errors`          | `errors.New`, `Is`, `As`, `Join` (aggregate per-file failures)                 |
| `sort` / `slices` | Deterministic ordering of the Plan and the stats report                        |
| `crypto/sha256`   | _(stretch)_ content hashing for `--duplicates`                                 |
| `testing`         | Table-driven tests against the fake filesystem                                 |

## 9. Concept coverage — where each chapter lands

| Chapter                                     | Exercised by                                                                      |
| ------------------------------------------- | --------------------------------------------------------------------------------- |
| 03 Types & Variables (+strings/bytes/runes) | `File`, `Category` enum, extension parsing, size formatting                       |
| 04 Control Flow                             | classification `switch`, scan loops, `defer` on file handles                      |
| 05 Functions                                | small pure functions (`scan`, `classify`, `destination`), function-value rules    |
| 06 Structs & Methods                        | `File`, `Operation`, `Plan` + value/pointer receiver decisions                    |
| 07 Interfaces ⭐                            | `FileSystem`, `Classifier`, `ConflictStrategy` — decoupling behaviour             |
| 08 Pointers                                 | mutable `Plan`/`Transaction`, `*Config` shared state                              |
| 09 Error Handling ⭐                        | wrapped move errors, `ErrDestinationExists`, `FileConflictError`, `errors.Join`   |
| 10 Collections ⭐                           | `[]File`, `map[Category][]File`, stats maps, sort/filter                          |
| 11 Packages & Modules                       | refactor `main.go` → `internal/*` as it grows; own `go.mod`                       |
| 12 Generics                                 | `Filter[T]` / `GroupBy[T,K]` **only if** they genuinely read better — else delete |
| 13 Testing ⭐                               | fake filesystem; table-driven classifier/planner/undo tests                       |
| 14 Building CLIs & JSON _(new)_             | `flag` subcommands, JSON config + journal round-trip                              |

## 10. Build roadmap (guided phases — the through-line)

Build one coherent codebase in phases; keep each phase's `git` history so that
`git diff v0.1 → v1.0` tells the story of learning to think in Go.

### Phase 0 — Walking skeleton `v0.1`

- **Goal:** `gorg <dir>` scans and _prints_ what it finds. No moving yet.
- **Build:** arg handling; scan top-level; print files + guessed category.
- **Concepts:** 03, 04, 05. **Libs:** `os`, `path/filepath`, `flag`, `fmt`.
- **Done when:** it lists a real directory's files with a category each.

### Phase 1 — Move for real `v0.2`

- **Goal:** actually sort files into category folders.
- **Build:** `Operation`, `MkdirAll`, `Rename`; the happy path only.
- **Concepts:** 06 (methods), early 09 (basic errors). **Libs:** `os`.
- **Done when:** a messy folder gets organized; errors are surfaced, not swallowed.

### Phase 2 — Domain & seams `v0.3`

- **Goal:** introduce the interfaces that make it testable and extensible.
- **Build:** `FileSystem` interface + OS impl; `Classifier` interface + extension
  impl; `Plan` built before execution.
- **Concepts:** 07 ⭐, 08, 09 ⭐, 10 ⭐. **Libs:** `io/fs`, `errors`.
- **Done when:** the organizer depends on interfaces, not concrete OS calls.

### Phase 3 — Safety: dry-run, conflicts, undo `v0.4`

- **Goal:** make it a tool you can trust.
- **Build:** `--dry-run`; `ConflictStrategy` (skip/overwrite/rename + `--interactive`);
  `Transaction` journal + `gorg undo`.
- **Concepts:** 09 ⭐ (domain errors, `errors.Join`), 10, 14 (JSON journal).
- **Libs:** `encoding/json`. **Done when:** you can preview, apply, and fully revert.

### Phase 4 — Config & CLI polish `v0.5`

- **Goal:** user-defined rules and a real subcommand CLI.
- **Build:** JSON config load/merge (`--config`, `UserConfigDir`); `stats` subcommand;
  `rules` subcommand; consider `Filter`/`GroupBy` generics — keep only if justified.
- **Concepts:** 11 (packages settle), 12 (generics, sparingly), 14 (flag subcommands).
- **Libs:** `encoding/json`, `sort`/`slices`. **Done when:** rules come from config and
  the package layout is clean.

### Phase 5 — Tests everywhere `v1.0`

- **Goal:** prove it with tests, hit the coverage bar.
- **Build:** fake filesystem; table-driven tests for classify / plan / conflict / undo;
  thin integration test of the OS adapter with `t.TempDir()`.
- **Concepts:** 13 ⭐. **Libs:** `testing`. **Done when:** `go test -race ./...` green,
  core packages ≥ 80%, and you'd approve your own code in review.

### Verifiable checkpoint (ties to the roadmap's Week 6)

> `gorg` organizes your real `~/Downloads`, `--dry-run` and `undo` work, tests pass,
> and you've actually used it. That's Part 1 done.

## 11. Post-Part-1 evolution (the "real tool" — NOT now)

Each of these is deliberately deferred so it can teach a _later_ part's concept when
you add it, turning `gorg` into a running example across the whole course:

- **Cobra + Viper** — richer CLI/config → first real third-party deps (revisits ch11).
- **TOML/YAML config** — nicer than JSON (BurntSushi/toml, gopkg.in/yaml).
- **`fsnotify` watch daemon** — auto-organize on file events → **Part 3 concurrency**
  (goroutines, channels, context cancellation, graceful shutdown).
- **bubbletea/lipgloss TUI** — interactive review/confirm screen.
- **Pluggable classifiers** — mime-type (`net/http.DetectContentType`), size, regex,
  rule DSL.
- **OSS release** — `goreleaser`, GitHub Actions, semantic versioning (**Part 4**).

## 12. Decisions locked (for the record)

- Milestone of Part 1 is **`gorg`**, replacing the earlier `gitm` idea.
- **Build at the end of Part 1** as its own module; chapters 03–13 keep their own
  `examples/`/`exercises/` and are _not_ rewritten to build `gorg` incrementally. The
  "grows over time" feel lives in `gorg`'s own Phase 0→5 history instead.
- **Part 1 is 100% standard library.** `flag` not cobra; JSON not TOML. Third-party
  deps, concurrency, and TUI are explicitly post-Part-1.
- A new **Chapter 14 — Building CLIs & JSON** is added to Part 1 so the project's core
  skills (arg parsing + `encoding/json`) are taught before/at the milestone; Chapter 03
  is deepened with strings/bytes/runes. (Roadmap renumber pending.)
