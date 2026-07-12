# Tooling

Everything is installed through **[mise](https://mise.jdx.dev)** — no Node, no
`node_modules`. See [ADR-0003](./adr/0003-node-free-docs-toolchain.md) for why.

## Install

```sh
just setup        # = mise install && lefthook install
```

`mise install` reads [`mise.toml`](../mise.toml) and installs pinned versions of:

| Tool            | Purpose                       | Config                              |
| --------------- | ----------------------------- | ----------------------------------- |
| `go`            | the language                  | `go.mod`, `go.work`                 |
| `gofumpt`       | stricter `gofmt` formatter    | —                                   |
| `golangci-lint` | Go linter                     | [`.golangci.yml`](../.golangci.yml) |
| `gitleaks`      | secret scanning               | —                                   |
| `dprint`        | format markdown / json / toml | [`dprint.json`](../dprint.json)     |
| `just`          | task runner                   | [`justfile`](../justfile)           |
| `lefthook`      | git hooks                     | [`lefthook.yml`](../lefthook.yml)   |

If you don't use mise, install those tools however you like — the config files above
are all standard.

## Daily commands

```sh
just fmt      # format Go (gofumpt) + docs/config (dprint), in place
just check    # gofumpt --check, go vet, golangci-lint, dprint check, tests, secrets
just test     # go test -race ./...
just lint     # golangci-lint only
just secrets  # gitleaks scan
```

## Git hooks

`lefthook` runs on every commit: it auto-formats staged Go and docs (re-staging the
fixes), lints Go, and blocks commits containing secrets. It deliberately does **not**
run the test suite — that's for `just check` / CI, to keep commits fast.

## Editor

`.vscode/settings.json` is set up for gopls: format-on-save with gofumpt, organize
imports, `golangci-lint` on save, and `-race` test flags.

## Notes

- **dprint plugins**: [`dprint.json`](../dprint.json) pins plugin versions. Refresh
  them with `dprint config update` (it rewrites the URLs and adds checksums).
- **Go toolchain**: `go.mod` requests `go 1.26.3`; mise installs a matching toolchain
  so `go`/`gopls`/CI all agree on the version.
