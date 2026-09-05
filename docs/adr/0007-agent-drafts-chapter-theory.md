# Agent drafts chapter theory when the owner has no draft

The earlier rule was that **all** `README.md` teaching prose is owner-authored
(composed with claude.ai and pasted in), and the agent must not "fabricate or rewrite
chapter teaching content unless explicitly asked." In practice that made
`scaffold-chapter`, when no `.docs/NN-slug.md` draft existed, emit a **skeleton of
`TODO`s** — headings with no lesson (Chapter 14 shipped exactly this way). An empty
shell is low value: the owner still has to write everything from scratch, and the
scaffold proved little beyond the folder layout.

**Decision.** Chapter theory now follows an **"agent drafts → owner refines"** model:

- If `.docs/NN-slug.md` exists, derive the README from it (verified) — unchanged.
- If not, the **agent writes a complete first draft** as an expert Go engineer: real
  prose, JS/TS comparisons, Common Mistakes, milestone ties — **not `TODO`s**. Claims
  are verified against primary Go sources (pkg.go.dev, spec, go.dev) while writing,
  and a cited notes file is left in the git-ignored `.docs/`.
- The **owner owns the final voice**. Once a chapter has been owner-edited, treat it
  as owner-owned: draft only what's missing, never silently overwrite refined prose.

**Why this trade-off.** We give up the guarantee that every published word is the
owner's own first draft. We gain a usable, fact-checked starting point for every
chapter instead of an empty checklist, while preserving the thing that actually
mattered in the old rule — the owner has the final say on voice and can rewrite
freely. The `.docs/` draft stays the preferred path when it exists; agent-drafting is
the fallback, not the default aspiration.

**Consequences.** `CLAUDE.md` (Content authorship) and the `scaffold-chapter` skill
are updated to match. Chapter 14 is the first chapter authored under this rule.
Reversing it means going back to `TODO` skeletons for undrafted chapters — cheap
mechanically, but it re-opens the "empty shell" problem this ADR closes.
