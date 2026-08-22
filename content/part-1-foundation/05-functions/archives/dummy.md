```text
main()
  │
  ▼
process()
  │
  ├── print "start"
  │
  ├── panic()
  │
  ▼
function stops
  │
  ▼
panic propagates to caller
  │
  ▼
program crashes
```

```text
process()
  │
  ├── defer cleanup
  │
  ├── print start
  │
  ├── panic()
  │
  ▼
run deferred functions
  │
  ├── cleanup
  │
  ▼
panic continues propagating
  │
  ▼
program crashes
```

```text
process()
  │
  ├── register defer
  │
  ├── print start
  │
  ├── panic()
  │
  ▼
panic unwinding
  │
  ▼
defer runs
  │
  ▼
recover()
  │
  ▼
panic stopped
  │
  ▼
process() returns
  │
  ▼
main continues
```

```txt
panic
 ↓
defer
 ↓
panic continues
 ↓
crash

panic
 ↓
defer
 ↓
recover()
 ↓
panic stopped
 ↓
continue
```

```txt
main()
  │
  │ c := counter()
  ▼
counter()
  │
  ├── count := 0
  │
  ├── create closure
  │       │
  │       └── captures count
  │
  └── return closure
          │
          ▼
          c

c
│
▼
┌─────────────────────┐
│ closure             │
│                     │
│ function:           │
│   count++           │
│   return count      │
│                     │
│ captured:           │
│   count = 0         │
└─────────────────────┘

main()
  │
  │ c()
  ▼
closure
  │
  ├── count++       // 0 → 1
  │
  └── return 1

c()
  │
  ▼
closure
  │
  ├── count++       // 1 → 2
  │
  └── return 2
```

```txt
main()
  │
  ├── c := counter()
  │       │
  │       ▼
  │    count = 0
  │       │
  │       └── return closure
  │
  ├── c()
  │    └── count: 0 → 1
  │
  ├── c()
  │    └── count: 1 → 2
  │
  │
  ├── c2 := counter()
  │        │
  │        ▼
  │     count = 0   ← NEW count
  │        │
  │        └── return closure
  │
  └── c2()
       └── count: 0 → 1
```
