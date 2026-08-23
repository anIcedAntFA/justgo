# Chapter 06 — Questions

Recall questions. Try to answer out loud before unfolding. If you can't, re-read the
matching section of [`README.md`](./README.md).

<details>
<summary>1. Go has no classes. What two things replace them, and why is keeping them separate considered a feature?</summary>

**Structs** hold the data (a typed collection of fields) and **methods** attach the
behavior (functions with a receiver). They're deliberately decoupled: a struct
definition and its methods are written separately, so data and behavior aren't welded
together the way they are inside a class. That separation makes types easier to reason
about and lets you attach methods to a type from elsewhere in the same package.

</details>

<details>
<summary>2. Show the five ways to create a <code>User</code> value. Which literal form should you avoid, and why?</summary>

```go
u1 := User{Name: "Alice", Email: "a@x.com", Age: 30} // named fields (preferred)
u2 := User{"Bob", "b@x.com", 25}                      // positional (avoid)
var u3 User                                           // zero value
u4 := User{Name: "Charlie"}                           // partial (rest zero-valued)
u5 := &User{Name: "Dave"}                             // pointer to a new struct
```

Avoid the **positional** form (`u2`): it breaks silently if you reorder or add fields.
Named fields are explicit and refactor-safe.

</details>

<details>
<summary>3. What is a struct tag? What do <code>json:"age,omitempty"</code> and <code>json:"-"</code> do?</summary>

A struct tag is a metadata string attached to a field, read at runtime (via reflection)
by libraries — JSON, database mappers, validators. `json:"age,omitempty"` marshals the
field under the key `age` and **omits it entirely when it holds its zero value**.
`json:"-"` means **never include this field** in JSON. Tags are Go's answer to what
decorators do in TypeScript; Chapter 15 covers them in depth.

</details>

<details>
<summary>4. What is a "receiver", and how does it differ from JavaScript's <code>this</code>?</summary>

The receiver is the type a method is attached to, named in parentheses before the method
name: `func (r Rectangle) Area() float64`. Unlike JS's implicit `this`, the receiver is
**explicit and named** (`r`), so there's no `this`-binding confusion and no
arrow-vs-regular-function surprises. Methods are also defined **outside** the type, not
inside a class body.

</details>

<details>
<summary>5. The core question: a value receiver vs a pointer receiver — what does each get, and which one can modify the original?</summary>

A **value receiver** (`func (r T)`) gets a **copy** of the struct; mutations touch only
the copy, so the original is unchanged — use it for read-only methods on small types. A
**pointer receiver** (`func (r *T)`) gets a **pointer** to the struct; mutations affect
the **original** — use it to modify state or to avoid copying a large struct. The
practical rule: when in doubt use a pointer receiver, and be consistent across all of a
type's methods.

</details>

<details>
<summary>6. You call <code>c.Increment()</code> where <code>Increment</code> has a pointer receiver but <code>c</code> is a value, not a pointer. Why does it compile? When does it <em>not</em>?</summary>

Because `c` is **addressable** (it's stored in a variable), Go automatically rewrites the
call as `(&c).Increment()` — you don't write the `&` yourself. It fails when the value
isn't addressable: calling a pointer-receiver method on a **literal** like
`Counter{}.Increment()` won't compile, because there's no address to take.

</details>

<details>
<summary>7. Go has no <code>new MyClass()</code> constructor. What's the convention, and what can a constructor do that a keyword can't?</summary>

The convention is a plain function named `NewTypeName` that returns the value (often a
pointer): `func NewServer(host string, port int) *Server`. Because it's an ordinary
function it can **validate inputs and return an error** (`(*Server, error)`), **set
defaults**, and control exactly how the struct is built — none of which a keyword could
do. (The builtin `new(T)` exists but only allocates zeroed memory; `new(Server)` is just
`&Server{}`.)

</details>

<details>
<summary>8. Go embeds one struct in another. Why is that HAS-A composition and not IS-A inheritance? What still gets "promoted"?</summary>

Embedding a `Shape` inside a `Circle` means a `Circle` **has** a `Shape`, not that a
`Circle` **is** a `Shape` — you can't assign a `Circle` to a `Shape` variable
(`var s Shape = Circle{}` won't compile). What you do get is **promotion**: the embedded
type's fields and methods become callable directly on the outer type (`c.Describe()`
works even though `Describe` is defined on `Shape`). For polymorphism — using one type
where another is expected — you need interfaces (Chapter 07).

</details>

<details>
<summary>9. If both the outer type and an embedded type define <code>Describe()</code>, which runs on the outer value? How do you still reach the embedded one?</summary>

The **outer type's** method wins — it shadows the embedded one by name resolution (there
is no virtual dispatch). `d.Describe()` calls `Derived`'s method; you reach the embedded
version explicitly through the embedded field name: `d.Base.Describe()`.

</details>

<details>
<summary>10. When are two structs equal with <code>==</code>, and which structs can't be compared at all?</summary>

Two structs are `==` when all their corresponding fields are equal — value equality, not
reference equality (so two separately-built `Point{1,2}` are equal, and comparable
structs work as **map keys**). A struct is **not comparable** if any field is a slice,
map, or function; comparing such a struct with `==` is a compile error. This is
impossible in JavaScript, where `{x:1} === {x:1}` is always `false`.

</details>
