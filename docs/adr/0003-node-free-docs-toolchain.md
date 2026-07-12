# Node-free docs toolchain with dprint

Markdown and config files are formatted with **dprint** (a single Rust binary
managed by mise), not with the `oxfmt` + `markdownlint-cli2` pair used in the
owner's dotfiles repo. Both of those tools run on Node.js and would drag
`package.json` + `node_modules` + pnpm into what is otherwise a pure Go repo. Since
this is a Go learning environment, keeping the toolchain node-free — everything
installable through `mise` alongside go, gofumpt, golangci-lint, and gitleaks — is
worth the deliberate divergence from the dotfiles convention.

We format only (dprint) and skip a dedicated markdown linter for now; `.editorconfig`
plus dprint formatting covers the common issues. A Rust markdown linter (e.g. mado)
can be added later if rule enforcement is missed — it is intentionally not there yet.
