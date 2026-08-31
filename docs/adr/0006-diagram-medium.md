# Diagram medium: ASCII for memory, Mermaid for flow

Some Chapters teach concepts that are **spatial** (how bytes and headers sit in
memory — a slice header pointing at a backing array, a struct's fields, a pointer
aliasing another value) or **temporal** (how control or data moves through a system —
a request lifecycle, a state machine, a scheduler). Prose alone teaches these poorly;
a picture earns its place. This ADR fixes _which_ picture, so the choice is consistent
across all 42 Chapters instead of decided ad-hoc each time.

## Decision

- **ASCII line-art for memory and structural layout** — anything where the _shape in
  memory_ is the lesson: slice headers and backing arrays, struct fields, byte
  buffers, pointer aliasing, capacity vs length. ASCII lives inline in a fenced block
  right beside the code it explains, diffs cleanly in git, and renders identically
  everywhere with no toolchain.
- **Mermaid for flow and state** — control flow with branches, data flow across
  components, state machines, sequence/lifecycle diagrams. GitHub renders ```mermaid
  fences natively, so a flow with branches or cycles stays readable without ASCII
  gymnastics.

ASCII is the **default**; reach for Mermaid only when a flow genuinely branches or
cycles in a way ASCII can't show cleanly. A Chapter whose concept is neither spatial
nor temporal (e.g. Types & Variables) carries **no diagram** — the trigger is the
concept, not a quota.

## Why not the alternatives

- **Mermaid for everything** — Mermaid draws memory/byte layout badly; you cannot
  express "a 3-word header pointing into slot 1 of a 5-slot array" as clearly as a few
  ASCII boxes. Forcing it would make the most important diagrams the worst.
- **ASCII for everything** — a branching control flow or a multi-component data flow
  becomes a tangle of hand-drawn arrows that rots on every edit; Mermaid expresses
  those as source that GitHub lays out for you.
- **A rendered-image pipeline (SVG/PNG from a tool)** — would drag a rendering
  toolchain (and likely Node) into the repo, contradicting
  [ADR-0003](./0003-node-free-docs-toolchain.md). Mermaid keeps us node-free: it is
  plain text in the Markdown, rendered by GitHub, formatted by `dprint` like any other
  fence.

## Consequence

This convention is a natural extension of the node-free stance in ADR-0003. The
`scaffold-chapter` skill points here rather than restating the rule, so the reasoning
lives in one place.
