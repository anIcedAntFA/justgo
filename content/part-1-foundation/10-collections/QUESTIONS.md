# Chapter 10 — Questions

Recall questions. Try to answer out loud before unfolding. If you can't, re-read the
matching section of [`README.md`](./README.md).

<details>
<summary>1. What's the fundamental difference between an array and a slice in Go, and why do you almost always reach for a slice?</summary>

An **array** has a fixed length that is _part of its type_ (`[3]int` and `[4]int` are
different types) and it is a **value** — assigning or passing it copies every element.
A **slice** is a dynamic, resizable view: a small 3-word header (pointer, length,
capacity) pointing at a separate backing array. You reach for slices because they grow
with `append`, are cheap to pass (you copy the header, not the data), and aren't
locked to one length. Arrays show up mostly as fixed buffers (`[32]byte`) or as the
backing storage behind slices.

</details>

<details>
<summary>2. Why must you write <code>s = append(s, x)</code> and never a bare <code>append(s, x)</code>?</summary>

`append` may need to grow the backing array; when it does, it allocates a **new,
larger array**, copies the elements over, and returns a slice header pointing at the
new array. If you don't reassign, you keep the old header — which still points at the
old, shorter array — and you've thrown away the appended element. Even when it doesn't
regrow, the returned header has the new length, so you always reassign.

</details>

<details>
<summary>3. What are the three fields of a slice header, and what's the difference between length and capacity?</summary>

A slice header holds a **pointer** to the backing array, a **length** (`len` — how
many elements you can index right now, `s[0]..s[len-1]`), and a **capacity** (`cap` —
how many slots exist in the backing array from the slice's start before it must
reallocate). You can append up to `cap` without a new allocation; crossing `cap`
triggers a regrow.

</details>

<details>
<summary>4. How does a slice's capacity grow as you append past it, and how do you avoid the regrowths entirely?</summary>

Growth is **amortized**: Go roughly doubles capacity while the slice is small and
switches to a gentler factor once it gets large, then rounds the result up to an
allocator size class — so the exact `cap` steps are runtime-defined, not a promise you
should hard-code. The takeaway is that `append` is _O(1) amortized_, not that it
follows a fixed sequence. When you know the size ahead of time, `make([]T, 0, n)`
preallocates the capacity so the loop never reallocates.

</details>

<details>
<summary>5. You take <code>sub := original[1:4]</code> and then set <code>sub[0] = 99</code>. What happens to <code>original</code>, and why?</summary>

`original` changes too — it becomes `[1 99 3 4 5]`. Slicing does **not** copy; `sub`
is a second header pointing into the _same_ backing array. `sub[0]` is physically the
same memory cell as `original[1]`, so a write through one is visible through the other.

</details>

<details>
<summary>6. Describe the append-on-subslice corruption trap and the two idiomatic fixes.</summary>

A subslice inherits the parent's spare **capacity**. So `head := parent[0:2]` has
`len 2` but `cap` reaching to the end of `parent`. Appending to `head` sees that spare
capacity and writes **in place** — over `parent`'s later elements — instead of
allocating. Fixes: (1) `slices.Clone(parent[0:2])` copies into a fresh backing array;
(2) the three-index full slice expression `parent[0:2:2]` caps the result so the next
`append` is forced to reallocate. Use one whenever you slice something and intend to
append to the result.

</details>

<details>
<summary>7. What does <code>copy</code> do, how many elements does it move, and how does it compare to <code>slices.Clone</code>?</summary>

`copy(dst, src)` copies `min(len(dst), len(src))` elements from `src` into `dst`'s
existing storage and returns the count — `dst` must already have length. `slices.Clone`
(Go 1.21) is the concise way to get a full, independent copy: it allocates a new slice
the same length and copies everything, so you don't manage the destination yourself.

</details>

<details>
<summary>8. What's the difference between a <code>nil</code> slice and an empty non-nil slice, and when does it matter?</summary>

Both have length 0 and are safe to `append` to and `range` over. The difference is
their state: a `nil` slice (`var s []int`) has no backing array and compares
`== nil`; an empty slice (`[]int{}`) is non-nil with a zero-length backing array. It
matters at boundaries that distinguish "absent" from "present but empty" — most
famously JSON, where a `nil` slice marshals to `null` while `[]int{}` marshals to `[]`.

</details>

<details>
<summary>9. Accessing a missing map key doesn't error — so how do you tell "key is absent" from "key is present with the zero value"?</summary>

A missing key returns the value type's **zero value** (`0`, `""`, `false`, `nil`),
never an error and never "undefined". To distinguish absent from present-but-zero, use
the **comma-ok** form: `v, ok := m[k]` — `ok` is `true` only when the key exists.

</details>

<details>
<summary>10. What happens when you write to a <code>nil</code> map versus read from one? How do you create a writable map?</summary>

Reading from a `nil` map is fine — it returns zero values. **Writing to a `nil` map
panics** (`assignment to entry in nil map`). Create a writable map with `make(map[K]V)`
or a map literal before writing to it.

</details>

<details>
<summary>11. What order does <code>range</code> visit a map in, and how do you iterate deterministically?</summary>

Map iteration order is **randomized** — deliberately, so code can't accidentally
depend on an order that was never guaranteed. Each run may differ. For deterministic
output, collect the keys into a slice and sort them: `for _, k := range
slices.Sorted(maps.Keys(m))`, then index the map by `k`.

</details>

<details>
<summary>12. Why does <code>m["a"].X = 10</code> fail to compile for a <code>map[string]Point</code>, and what's the workaround?</summary>

Map values are **not addressable**, so you can't assign to a field of a struct value
sitting in a map. Two options: replace the whole value (`p := m["a"]; p.X = 10; m["a"]
= p`), or make it a **map of pointers** (`map[string]*Point`), where `m["a"].X = 10`
works because you're modifying through the pointer.

</details>

<details>
<summary>13. What does the <code>clear</code> builtin (Go 1.21) do to a slice versus a map?</summary>

On a **map**, `clear(m)` deletes every key, leaving `len(m) == 0`. On a **slice**,
`clear(s)` sets every element to its zero value while leaving the length unchanged.
One builtin, two type-appropriate meanings.

</details>

<details>
<summary>14. Go has no built-in <code>.map()</code>/<code>.filter()</code>. What do you use instead, and what does the <code>slices</code> package give you (Go 1.21+)?</summary>

You write an explicit `for range` loop (preallocating the result when you know the
size), or a generic helper (Chapter 12). The `slices` package covers the common
operations: `Sort`, `SortFunc`, `Contains`, `Index`, `Max`, `Min`, `Reverse`,
`Clone`, `Equal`, plus editing helpers like `Delete`, `Insert`, `Compact`, and
`BinarySearch`. Note you **cannot** compare slices or maps with `==` (only against
`nil`) — use `slices.Equal` / `maps.Equal` for content comparison.

</details>

<details>
<summary>15. What did Go 1.22 change about the loop variable in a <code>for range</code>, and what bug did it fix?</summary>

Before 1.22 the loop variable was created **once** and updated each iteration, so a
goroutine or closure that captured it saw the _last_ value — the classic
"loop-variable capture" bug. In Go 1.22 **each iteration creates a new variable**, so
captured loop variables behave as you'd expect. (It's gated on the `go 1.22`+ directive
in `go.mod`.) Go 1.22 also added ranging over an integer: `for i := range n`.

</details>

<details>
<summary>16. What is an iterator in Go 1.23 (range-over-func), and how is it like a JS generator?</summary>

A function whose type is `iter.Seq[T]` — i.e. `func(yield func(T) bool)` — can be used
directly in `for v := range f()`. The function _pushes_ values by calling `yield`;
`yield` returns `false` when the consumer breaks, letting the producer stop early. This
gives lazy, custom sequences without materializing a whole collection, much like a JS
generator's `yield`. The `maps` and `slices` packages expose iterators too
(`maps.Keys`, `slices.Collect`, `slices.Sorted`).

</details>
