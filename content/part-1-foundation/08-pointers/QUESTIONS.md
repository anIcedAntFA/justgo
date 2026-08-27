# Chapter 08 — Questions

Recall questions. Try to answer out loud before unfolding. If you can't, re-read the
matching section of [`README.md`](./README.md).

<details>
<summary>1. What does a pointer hold, and what do <code>&</code> and <code>*</code> each do?</summary>

A pointer holds the **memory address** of a value, not the value itself. `&x` is the
**address-of** operator — it gives you a pointer to `x`. `*p` is the **dereference**
operator — it gives you the value stored at the address `p` holds (and you can assign
through it: `*p = 100`). Note `*` does double duty: in a **type** like `*int` it means
"pointer to int"; as an **operator** on a value it means "dereference".

</details>

<details>
<summary>2. Go passes arguments by value. So how can a function modify the caller's variable?</summary>

By passing a **pointer**. Every argument is copied, so `func inc(n int)` only ever
mutates its local copy. `func inc(n *int)` receives a copy of the _address_; dereferencing
it (`*n++`) reaches the original. The pointer itself is copied — but it points at the same
underlying value.

</details>

<details>
<summary>3. How do Go's value semantics differ from JavaScript's for structs/objects?</summary>

In JavaScript, objects are **always references** — `obj2 = obj1` makes both names point
at the same object. In Go, structs are **values**: `s2 := s1` **copies** the whole struct,
so the two are independent. To get JS-style shared-reference behavior in Go you take a
pointer explicitly (`p2 := p1` shares). Primitives copy in both languages.

</details>

<details>
<summary>4. Why does <code>u.Name</code> work when <code>u</code> is a <code>*User</code>? Don't you need <code>(*u).Name</code>?</summary>

Go **automatically dereferences** a pointer in a selector expression. Writing `u.Name` on
a `*User` is shorthand the compiler expands to `(*u).Name`. The same auto-dereference
applies to **method calls** through a pointer. You practically never write `(*u).Name`
yourself.

</details>

<details>
<summary>5. What happens if you dereference a nil pointer, and how do you prevent it?</summary>

It **panics** at run time: `invalid memory address or nil pointer dereference`. A pointer's
zero value is `nil` (points at nothing). Go has **no optional chaining** (`?.`), so you
guard explicitly: `if p != nil { … }` before touching any field or method that dereferences.
This is one of the most common Go runtime panics.

</details>

<details>
<summary>6. Give two distinct reasons to reach for a pointer.</summary>

1. **Mutation** — you need the function/method to modify the caller's value (a pointer
   receiver, or a `*T` parameter).
2. **Efficiency** — the value is a large struct and you want to avoid copying it on every
   call; a pointer copies just the address (8 bytes on a 64-bit machine).

A third, softer reason: **representing absence** — a `*int` field where `nil` means "not
set" versus a real pointer to a value (even `0`).

</details>

<details>
<summary>7. When should you prefer a value over a pointer?</summary>

When the type is **small and naturally a value** (a couple of primitive fields, like a
`Point{X, Y int}` or a `time.Time`), and especially when you want **immutability** — a value
parameter is a safe copy the callee can't mutate. Don't use a pointer for a small value
you're only reading: it adds indirection and nil-risk for no benefit. The common default
for structs _with methods_, though, is pointers.

</details>

<details>
<summary>8. Is it safe to return a pointer to a local variable? Why doesn't the memory vanish?</summary>

Yes — safe and common in Go. The compiler's **escape analysis** notices the address
outlives the function and allocates the variable on the **heap** instead of the stack, so
the pointer stays valid. The garbage collector frees it once no references remain. (In C
this same code dangles.) You can see the compiler's decisions with
`go build -gcflags=-m`.

</details>

<details>
<summary>9. Why does Go forbid pointer arithmetic, and what operations <em>are</em> allowed?</summary>

Forbidding arithmetic (`p++`, `p + 5`) is a **safety** decision — it eliminates buffer
overflows and use-after-free-via-arithmetic. Go pointers may only: be **taken** (`&x`), be
**dereferenced** (`*p`), be **compared** (`==`, `!=`, including against `nil`), and be
`nil`. The `unsafe` package is the low-level escape hatch for systems/interop code — you'll
rarely need it.

</details>

<details>
<summary>10. What does <code>new(T)</code> do, and why do you usually write <code>&T{}</code> instead?</summary>

`new(T)` allocates **zeroed** storage for a `T` and returns a `*T`. It's equivalent to
`&T{}` for a struct, but it can't initialize fields. `&T{...}` is the idiomatic form because
it lets you set fields as you allocate: `&User{Name: "Alice"}`. Reach for `new` mainly to
get a pointer to a zeroed primitive (`new(int)`); reach for `&T{}` for almost everything else.

</details>
