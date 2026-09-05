# Go Mastery — Learning Path

> **From JavaScript/TypeScript frontend to Go backend. Practical, modern, project-driven.**

## TL;DR

A 4-6 month roadmap to learn Go from zero to production-ready, designed for an experienced JS/TS frontend developer. Covers language fundamentals → web development → advanced patterns → capstone project. Each phase ends with a real, usable project. JS/TS comparisons are woven in where they genuinely help — not forced.

Go version: **1.26** (latest stable, February 2026).

## ELI5 🧒

JavaScript is like a Swiss Army knife — it does everything, but nothing perfectly. Go is like a Japanese chef's knife — it does fewer things, but each one with precision and speed. This path teaches you to wield that knife: simple tools, sharp results, no magic.

---

## 👤 Target Audience

- Middle/Senior Frontend Developer
- Proficient in JavaScript/TypeScript, React ecosystem
- Wants to expand into backend development and CLI tooling
- Prefers practical, modern, no-fluff learning

## 🎯 Success Criteria

Upon completion, you will be able to:

1. ✅ Read and write idiomatic Go — not "JavaScript written in Go syntax"
2. ✅ Build production-ready REST APIs with proper error handling, middleware, and database integration
3. ✅ Understand and use Go's concurrency model (goroutines, channels, select) for real problems
4. ✅ Design and build CLI tools that you actually use daily
5. ✅ Architect a multi-component backend service (API gateway with proxy, auth, rate limiting)
6. ✅ Know Go conventions, tooling, and ecosystem well enough to contribute to existing Go projects

---

## 🧠 Why Go? (For a JS/TS Dev)

**What Go gives you that JS doesn't:**

- True concurrency (goroutines are not async/await — they run in parallel)
- Single binary deployment (no node_modules, no runtime, no Docker gymnastics)
- Compile-time type safety that's actually simple (not TypeScript's type gymnastics)
- Predictable performance (no GC pauses you can't reason about, no JIT warmup)
- A standard library that covers HTTP servers, JSON, crypto, testing — without npm

**What Go takes away:**

- No classes, no inheritance, no decorators, no generics magic (until recently)
- No try/catch — errors are values you handle explicitly, every time
- No ternary operator, no optional chaining, no nullish coalescing
- No implicit type coercion, no `any` escape hatch (well, there's `interface{}` but it's discouraged)
- Formatter is not optional — `gofmt` is law

**The mental shift:**

```
JS/TS mindset:  "How do I make this elegant/clever?"
Go mindset:     "How do I make this obvious/boring?"
```

Go's philosophy is radical simplicity. Code should be readable by someone who's never seen it before. If you find yourself being clever in Go, you're doing it wrong.

---

## 📐 Overall Structure

```
Part 1: Foundation ────── Go basics, language mechanics, mental model shift
Part 2: Web & APIs ────── HTTP, REST, middleware, database, real-world patterns  
Part 3: Advanced ──────── Concurrency deep-dive, networking, system design
Part 4: Production ────── Capstone project, deployment, production patterns
```

---

## 📚 Part 1 — Foundation (Week 1-6)

> **Goal:** Think in Go. Write idiomatic Go. Understand why Go is the way it is.

### Chapters

| #  | Chapter              | Key Topics                                                                                                                                    |
| -- | -------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| 01 | History & Philosophy | Why Google created Go, the problems it solves, design principles, Go vs JS/TS mental model                                                    |
| 02 | Setup & Tooling      | Installation, Go workspace, modules, go CLI, editor setup, gofmt/goimports/golangci-lint                                                      |
| 03 | Types & Variables    | Basic types, zero values, type inference, constants, type conversions (no coercion!)                                                          |
| 04 | Control Flow         | if/else, for (the only loop), switch (much more powerful than JS), defer                                                                      |
| 05 | Functions            | Multiple returns, named returns, variadic, first-class functions, closures                                                                    |
| 06 | Structs & Methods    | Structs as data containers, methods with receivers, pointer vs value receivers, composition over inheritance                                  |
| 07 | Interfaces           | Implicit satisfaction, empty interface, type assertions, type switches, io.Reader/Writer pattern                                              |
| 08 | Pointers             | Why they exist, when to use them, pointer receivers, nil pointers, no pointer arithmetic                                                      |
| 09 | Error Handling       | Errors as values, error wrapping (fmt.Errorf + %w), errors.Is/As, custom error types, sentinel errors                                         |
| 10 | Collections          | Arrays, slices (capacity, append, copy, gotchas), maps, range, iterators (Go 1.23+)                                                           |
| 11 | Packages & Modules   | Package organization, visibility (exported vs unexported), go.mod, dependency management, internal packages                                   |
| 12 | Generics             | Type parameters, constraints, when to use vs when not to, comparison with TS generics                                                         |
| 13 | Testing              | Built-in testing, table-driven tests, test coverage, benchmarks, testify basics                                                               |
| 14 | Building CLIs & JSON | `flag` package, manual subcommand dispatch, `os.Args`, exit codes, `encoding/json` (Marshal/Unmarshal), struct tags, JSON round-trip to files |

### Milestone Project: `gorg` — File Organizer CLI

A fast, safe file organizer built from the Go **standard library only**. Point it at a messy directory (Downloads, screenshots, a project dump) and it sorts files into category folders by configurable rules — safely, with preview and undo. You use this daily on your Arch Linux setup. Full spec: [`gorg/PLAN.md`](./part-1-foundation/gorg/PLAN.md).

```
$ gorg <dir>               → organize a directory (default command)
$ gorg <dir> --dry-run     → preview operations; change nothing
$ gorg <dir> --recursive   → descend into subdirectories
$ gorg <dir> --config <f>  → use a custom JSON rules file
$ gorg stats <dir>         → report counts/sizes per category
$ gorg undo                → revert the most recent run
$ gorg rules               → print the effective ruleset

Storage: ~/.config/gorg/ (rules + undo journal, JSON)
```

**Go concepts exercised:** structs & methods, interfaces (`FileSystem`, `Classifier`, `ConflictStrategy`), pointers, error handling (wrapping, sentinel errors, `errors.Join`), collections, packages, generics (sparingly), testing (fake filesystem), CLI args (`flag`) + `encoding/json`.

**Structure:** built at the end of Part 1 as its own Go module — chapters 03–14 keep their own `examples/`/`exercises/`; the "grows over time" feel lives in `gorg`'s own Phase 0→5 git history.

**Bonus (post-Part 1):** cobra/viper, TOML config, `fsnotify` watch daemon (Part 3 concurrency), bubbletea/lipgloss TUI.

### Verifiable Checkpoints

```
Week 2:  Can explain Go's type system, zero values, and why there's no "undefined"
         → Verify: write a program without looking at docs

Week 4:  Can model a domain with structs + interfaces, handle all errors idiomatically
         → Verify: code review your own code — would a Go dev approve?

Week 6:  gorg CLI works — organizes ~/Downloads, --dry-run and undo work, has tests
         → Verify: `go test ./...` passes, you actually tidied a real folder with it
```

---

## 📚 Part 2 — Web & APIs (Week 7-12)

> **Goal:** Build real HTTP services. Understand Go's net/http at the foundation level before reaching for frameworks.

### Chapters

| #  | Chapter                     | Key Topics                                                                         |
| -- | --------------------------- | ---------------------------------------------------------------------------------- |
| 15 | HTTP Fundamentals in Go     | net/http, HandlerFunc, ServeMux (new Go 1.22+ routing), request/response lifecycle |
| 16 | JSON & Serialization        | encoding/json, struct tags, custom marshalers, streaming JSON, validation          |
| 17 | Middleware Pattern          | Handler chaining, logging, auth, CORS, recovery, request ID — all from scratch     |
| 18 | Project Structure           | Flat vs layered, domain-driven layout, internal/, cmd/, when to split packages     |
| 19 | Database Fundamentals       | database/sql, connection pooling, prepared statements, transactions, migrations    |
| 20 | Repository Pattern          | Clean data access, interface-based repos, testability, sqlc introduction           |
| 21 | Configuration & Environment | env vars, config files, 12-factor app, Viper vs stdlib approaches                  |
| 22 | Logging & Observability     | slog (structured logging, Go 1.21+), log levels, request tracing                   |
| 23 | Intro to Frameworks         | Echo vs Chi vs Gin vs stdlib — tradeoffs, when to use which, hands-on with Echo    |
| 24 | Authentication Basics       | JWT basics, middleware auth, bcrypt, session vs token for Go services              |
| 25 | API Design                  | RESTful conventions, pagination, filtering, error responses, API versioning        |

### Milestone Project: `dropshare` — File/Content Sharing Service

A mini transfer.sh — upload anything, get a link, anyone can download. Auto-expire. You use this to share files quickly in your team.

```
POST   /upload          → upload file/text → returns short link
GET    /:slug           → view/download content  
GET    /:slug/meta      → metadata (size, type, created, expires)
DELETE /:slug           → delete early (with delete token)

Features:
- Auto expiration (1h, 24h, 7d, custom)
- Max file size limit
- Content-type detection (render markdown, highlight code, display image)
- Streaming upload/download (no full file in RAM)
- Optional: password protect
- Optional: one-time view (self-destruct after first access)
```

**Go concepts exercised:** net/http server, middleware chain, file I/O streaming (io.Reader/Writer), database (SQLite or Postgres), background goroutines (cleanup expired content), JSON API, configuration, structured logging.

### Verifiable Checkpoints

```
Week 8:  Can build a REST API with net/http, proper middleware, JSON responses
         → Verify: API serves correctly with curl, no framework used

Week 10: Can connect to database, write clean repository layer, handle transactions
         → Verify: integration tests pass against real DB

Week 12: dropshare service works end-to-end, with tests and proper error handling
         → Verify: upload a file from curl, download from browser, expired files auto-cleaned
```

---

## 📚 Part 3 — Advanced (Week 13-20)

> **Goal:** Master Go's concurrency model. Build networking tools. Understand system-level patterns.

### Chapters

| #  | Chapter                      | Key Topics                                                                                               |
| -- | ---------------------------- | -------------------------------------------------------------------------------------------------------- |
| 26 | Goroutines Deep Dive         | Goroutine lifecycle, stack growth, scheduling (GMP model), goroutine vs OS thread vs JS event loop       |
| 27 | Channels                     | Unbuffered vs buffered, directional channels, channel patterns, when channels vs mutexes                 |
| 28 | Select & Multiplexing        | select statement, timeouts, cancellation, fan-in/fan-out                                                 |
| 29 | sync Package                 | Mutex, RWMutex, WaitGroup, Once, Pool, atomic operations                                                 |
| 30 | Context                      | context.Context everywhere, cancellation propagation, timeouts, values, why Go passes context explicitly |
| 31 | Concurrency Patterns         | Worker pools, pipelines, rate limiters, semaphores, errgroup                                             |
| 32 | Networking Fundamentals      | net package, TCP/UDP, connection handling, timeouts, keep-alive                                          |
| 33 | WebSocket                    | gorilla/websocket or nhooyr/websocket, upgrade handshake, concurrent read/write, connection management   |
| 34 | Reverse Proxy                | httputil.ReverseProxy, load balancing algorithms, health checks, header manipulation                     |
| 35 | Performance & Profiling      | pprof, benchmarks, escape analysis, memory allocation, race detector                                     |
| 36 | Code Generation & Reflection | go generate, reflect package (when/why), build tags                                                      |

### Milestone Projects (pick order by interest)

**Project A: `goproxy` — Reverse Proxy / Load Balancer**

Understand what sits between browser and server. As a FE dev, this demystifies Nginx/Traefik.

```
Features:
- Reverse proxy with multiple backends
- Round-robin / least-connections load balancing
- Health check goroutines (periodic backend pinging)
- Request/response header manipulation
- Access logging with slog
- Graceful shutdown
- Config file (YAML) for backend definitions
```

**Go concepts exercised:** httputil.ReverseProxy, goroutines for health checks, sync.RWMutex for backend list, context cancellation, net/http transport customization.

**Project B: `gochat` — Real-time Chat Server**

See Go's concurrency model shine — one goroutine per connection, broadcasting via channels.

```
Features:
- WebSocket server with rooms/channels
- 1 goroutine per connection (read) + 1 (write)
- Hub pattern: central broadcaster managing all connections
- Join/leave rooms, direct messages
- Message history (in-memory ring buffer or DB)
- Connection heartbeat/ping-pong
- Graceful disconnect handling
```

**Go concepts exercised:** goroutines (hundreds concurrent), channels for message passing, sync.Map or mutex for connection registry, select for multiplexing, context for connection lifecycle.

### Verifiable Checkpoints

```
Week 15: Can explain goroutines vs threads vs async/await, write race-free concurrent code
         → Verify: `go test -race ./...` passes on concurrent code

Week 17: goproxy handles multiple backends, health checks run concurrently, no races
         → Verify: kill a backend, proxy auto-removes it, requests still served

Week 20: gochat handles 100+ concurrent connections, messages broadcast correctly
         → Verify: load test with multiple WebSocket clients, no message loss
```

---

## 📚 Part 4 — Production (Week 21-26)

> **Goal:** Tie everything together. Build something production-grade. Learn deployment and operational patterns.

### Chapters

| #  | Chapter                 | Key Topics                                                                                         |
| -- | ----------------------- | -------------------------------------------------------------------------------------------------- |
| 37 | Production Architecture | Service design, clean architecture in Go, dependency injection without frameworks                  |
| 38 | Advanced Middleware     | Rate limiting (token bucket, sliding window), circuit breaker, request validation, compression     |
| 39 | gRPC & Protocol Buffers | Why gRPC, protobuf, code generation, gRPC vs REST tradeoffs, gRPC-Gateway                          |
| 40 | Docker & Deployment     | Multi-stage Docker builds (Go binary = tiny image), health checks, graceful shutdown, 12-factor    |
| 41 | CI/CD & Release         | GitHub Actions for Go, goreleaser, linting in CI, semantic versioning                              |
| 42 | Security Hardening      | Input validation, SQL injection prevention, rate limiting, CORS, security headers, OWASP for Go    |
| 43 | Monitoring & Metrics    | Prometheus metrics, health endpoints, structured logging in production, distributed tracing basics |

### Capstone Project: `gogate` — API Gateway

Combine everything: proxy + auth + rate limiting + logging + monitoring. A mini Kong/Traefik you built yourself.

```
Architecture:
┌──────────────────────────────────────────────────┐
│                    gogate                         │
│                                                  │
│  ┌─────────┐  ┌──────────┐  ┌───────────────┐   │
│  │  Rate    │→ │   Auth   │→ │  Reverse      │   │
│  │  Limiter │  │  (JWT)   │  │  Proxy        │   │
│  └─────────┘  └──────────┘  └───────┬───────┘   │
│                                     │            │
│  ┌─────────┐  ┌──────────┐         │            │
│  │ Logger  │  │ Metrics  │         │            │
│  │ (slog)  │  │(Prometh.)│         │            │
│  └─────────┘  └──────────┘         │            │
└─────────────────────────────────────┼────────────┘
              ┌───────────────────────┼──────┐
              ▼           ▼           ▼      ▼
          Service A   Service B   Service C  ...

Features (MVP):
- YAML-based route configuration
- Reverse proxy to multiple backend services
- JWT authentication middleware
- Rate limiting per client/route (token bucket)  
- Request/response logging (structured, slog)
- Prometheus metrics endpoint
- Health check aggregation
- Graceful shutdown
- Docker multi-stage build

Extended:
- Circuit breaker for failing backends
- Request/response transformation
- API key management
- Admin API for runtime config changes
- WebSocket proxying
- Load balancing strategies (round-robin, weighted)
```

**Why this capstone works:**

| Aspect       | What it proves                                        |
| ------------ | ----------------------------------------------------- |
| Networking   | Deep HTTP understanding, proxy mechanics              |
| Concurrency  | Concurrent request handling, background health checks |
| Architecture | Clean separation, middleware chain, DI                |
| Production   | Docker, metrics, logging, graceful shutdown           |
| Real utility | You can actually put this in front of your services   |

### Verifiable Checkpoints

```
Week 22: Core gateway works — routes requests to backends, auth middleware blocks unauthorized
         → Verify: curl through gateway to backend, JWT required

Week 24: Rate limiting, metrics, logging all functional
         → Verify: exceed rate limit → 429, check Prometheus metrics

Week 26: Dockerized, CI pipeline, load tested
         → Verify: docker-compose up, locust/k6 load test, no crashes under load
```

---

## 📖 Chapter Format

Every chapter follows this structure:

```
## TL;DR
2-3 lines. What this chapter is about and why it matters.

## Core Concept  
Main explanation with code examples.
JS/TS comparison inline where it genuinely helps understanding.

## Code Examples
Working, runnable Go code. Not snippets — complete programs you can `go run`.

## Common Mistakes
What Go beginners (especially from JS) get wrong here.

## Exercises
2-3 small coding challenges to reinforce the concept.

## Key Takeaways
Bullet points of what to remember.
```

No emojis overload. No filler. Every section earns its place.

The `## Exercises` and recall questions don't stay inline: **coding exercises live
as test-driven code under `exercises/`**, and **recall questions live in
`QUESTIONS.md`** (see the File Structure below). The README keeps the teaching prose;
the doing and the quizzing get their own files.

---

## 🗂️ File Structure

Each **Chapter** is a folder `NN-slug/` holding its own theory, code, and questions.
Each **Milestone project** is a folder inside its Part with its own `go.mod`. See
[`CONTEXT.md`](../CONTEXT.md) for the exact meaning of Part / Chapter / Milestone
project / Exercise / Question, and [`docs/adr/`](../docs/adr/) for why the repo is
shaped this way.

```
justgo/
├── README.md                       ← repo landing → points here
├── CONTEXT.md                      ← glossary (Part, Chapter, project, …)
├── go.work                         ← ties root module + project modules
├── go.mod                          ← root module: all Chapter code
├── docs/
│   ├── adr/                        ← architecture decision records
│   └── tooling.md                  ← how to install & run the toolchain
├── content/
│   ├── ROADMAP.md                  ← 📍 You are here
│   └── part-1-foundation/
│       ├── README.md               ← Part index
│       ├── 01-history-and-philosophy/
│       │   ├── README.md           ← the chapter (theory)
│       │   └── QUESTIONS.md        ← recall questions
│       ├── 02-setup-and-tooling/
│       │   ├── README.md
│       │   └── QUESTIONS.md
│       ├── 03-types-and-variables/
│       │   ├── README.md
│       │   ├── QUESTIONS.md
│       │   ├── examples/           ← runnable demos (go run .)
│       │   └── exercises/          ← test-driven exercises (go test)
│       └── …                       ← 04–14 as you reach them
│       └── gorg/                   ← Milestone project (own go.mod)
└── (parts 2–4 scaffolded when you get there)
```

**Per-Chapter folder convention** (the mould — copy chapter 03):

- `README.md` — the theory. GitHub renders it as the folder's index page.
- `QUESTIONS.md` — interview-style recall questions with foldable answers.
- `examples/` — runnable `package main` demos referenced by the README (`go run .`).
- `exercises/` — a normal package with stubs + `_test.go`; make the tests pass.

Chapters with no code (e.g. 01 History) skip `examples/` and `exercises/`.

---

## ⏱️ Timeline

| Part       | Chapters                     | Est. Time                | Pace                                 |
| ---------- | ---------------------------- | ------------------------ | ------------------------------------ |
| Foundation | 01–14 + gorg project         | 6 weeks                  | ~2-3 chapters/week                   |
| Web & APIs | 15–25 + dropshare project    | 6 weeks                  | ~2 chapters/week                     |
| Advanced   | 26–36 + goproxy + gochat     | 8 weeks                  | ~1-2 chapters/week (harder material) |
| Production | 37–43 + gogate capstone      | 6 weeks                  | ~1 chapter/week + heavy project work |
| **Total**  | **43 chapters + 4 projects** | **~26 weeks (6 months)** | **1-3h/day**                         |

> Chapters are not equal length. Part 1 chapters are shorter (building blocks). Part 3-4 chapters are denser. Adjust pace accordingly — deep understanding of 1 concept beats skimming through 5.

---

## 🔗 Recommended Resources (Modern, Updated)

**Official:**

- [Go Documentation](https://go.dev/doc/) — always the source of truth
- [Go by Example](https://gobyexample.com/) — quick reference for syntax
- [Effective Go](https://go.dev/doc/effective_go) — official style guide
- [Go Blog](https://go.dev/blog/) — feature deep dives from the Go team
- [Go 1.26 Release Notes](https://go.dev/doc/go1.26) — latest language changes

**Learning:**

- [roadmap.sh/golang](https://roadmap.sh/golang) — visual roadmap, good for tracking progress
- [Let's Go](https://lets-go.alexedwards.net/) — Alex Edwards' book, best practical web dev in Go
- [Let's Go Further](https://lets-go-further.alexedwards.net/) — advanced API patterns
- [Learning Go](https://www.oreilly.com/library/view/learning-go-2nd/9781098139285/) — Jon Bodner, 2nd edition, idiomatic Go

**Concurrency:**

- [Concurrency in Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/) — Katherine Cox-Buday, the concurrency bible

**Community:**

- [r/golang](https://reddit.com/r/golang) — active community
- [Gophers Slack](https://gophers.slack.com/) — official Slack workspace
- [Go Weekly Newsletter](https://golangweekly.com/) — stay current

---

## 🧭 Navigation

| Direction | Link                                                                                          |
| --------- | --------------------------------------------------------------------------------------------- |
| **Start** | [Chapter 01: History & Philosophy →](./part-1-foundation/01-history-and-philosophy/README.md) |
