# Multi-module go.work, split by content role

The repo is a Go workspace (`go.work`) with a deliberate module boundary: **all
Chapter code lives in the single root Go module**, while **each Milestone project is
its own Go module** with its own `go.mod`. Chapter Exercises are tiny and share one
dependency set, so a single module keeps them frictionless; Milestone projects each
carry their own dependencies (cobra, bubbletea, sqlite drivers…) and their own
lifecycle, so isolating them keeps one project's deps out of another's and out of
the learning code.

This also makes the ADR-0001 exception cheap: a Milestone project that graduates to
its own repo is already a self-contained module and lifts out cleanly. The cost is
having to understand `go.work` early — acceptable, since Go modules and workspaces
are themselves Chapter 11 material.

We start multi-module from day one rather than beginning single-module and
splitting later, because splitting later would rewrite import paths.
