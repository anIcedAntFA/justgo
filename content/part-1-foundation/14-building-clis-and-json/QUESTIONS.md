# Chapter 14 — Questions

Recall questions. Try to answer out loud before unfolding. If you can't, re-read the
matching section of [`README.md`](./README.md).

<details>
<summary>1. What do <code>os.Args[0]</code> and <code>os.Args[1:]</code> hold, and what does <code>flag.Args()</code> give you <em>after</em> <code>flag.Parse()</code>?</summary>

`os.Args[0]` is the program's path/name; `os.Args[1:]` are the raw arguments as
passed. After `flag.Parse()`, `flag.Args()` returns the **non-flag** (positional)
arguments left over — e.g. the `<dir>` in `gorg <dir> --dry-run`.

</details>

<details>
<summary>2. Why does <code>encoding/json</code> ignore a struct field named <code>size</code> (lowercase) but marshal one named <code>Size</code>?</summary>

`encoding/json` uses reflection and can only see **exported** (capitalized) fields.
An unexported field like `size` is invisible — it won't be marshaled or unmarshaled.
Capitalize it (`Size`) and use a struct tag to control the JSON key:
`Size int64 \`json:"size"\``.

</details>

<details>
<summary>3. What does the <code>omitempty</code> option in <code>json:"path,omitempty"</code> do, and what counts as "empty"?</summary>

`omitempty` drops the field from the JSON output when its value is the type's **zero
value** — `0`, `""`, `false`, `nil`, or an empty slice/map. Useful for optional config
keys so the file stays clean.

</details>

<details>
<summary>4. How do you implement git-style subcommands (<code>gorg stats</code>, <code>gorg undo</code>) with only the <code>flag</code> package?</summary>

Switch on `os.Args[1]` (the subcommand name), then give each subcommand its own
`flag.NewFlagSet("stats", flag.ExitOnError)` and call `fs.Parse(os.Args[2:])`. Each
subcommand gets its own flags, isolated from the others — no third-party router needed.

</details>

<details>
<summary>5. When would you reach for <code>json.NewEncoder(w).Encode(v)</code> instead of <code>json.Marshal(v)</code>?</summary>

Use an `Encoder`/`Decoder` when the source or destination is a **stream** — a file, an
`http` body, a network connection — so you don't buffer the whole payload in memory.
`Marshal`/`Unmarshal` are for when you already have (or want) the bytes in memory. For
a human-readable config file, `MarshalIndent` (or `Encoder.SetIndent`) pretty-prints.

</details>

<details>
<summary>6. Why is <code>os.Exit(1)</code> inside a helper function risky, and what's the idiomatic alternative?</summary>

`os.Exit` terminates immediately and **skips every pending `defer`** — file flushes,
cleanup, unlocks all get lost. The idiom is to **return an error up to `main`** and
call `os.Exit` (or `log.Fatal`) in exactly one place, after deferred cleanup has run.

</details>
