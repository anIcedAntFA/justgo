# Dummy

## Mental model

- File define Read method.
- `io.Reader` required a method has signature `Read(...)`.
- Due to File has method Read => File satisfy `io.Reader`.

```txt
IMPLEMENTATION
concrete type
     │
     │ defines
     ▼
Read(...)
     │
     │ matches requirement
     ▼
io.Reader
  INTERFACE
```

- Interface **does not** own Read implementation.

## Distinguish between

1. Interface: describe contract/behavior.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

> Bất kỳ thứ gì muốn được coi là Reader phải có method này.

2. Implementation: a concrete type provide behavior. File là một implementation của `io.Reader`.

```go
type File struct {
    // ...
}

func (f *File) Read(p []byte) (int, error) {
    // đọc từ file
}
```

3. Consumer: function that use Interface. ReadAll is consumer of `io.Reader`

```go
func ReadAll(r io.Reader) {
  ...
}
```

> Tao ko cần biết mày là File, HTTP, body hay Buffer, tao chỉ cần capability Read là được.

4. Concrete value / argument: thứ caller thực sự truyền vào Consumer.

```go
file := ...
readAll(file)
```

```txt
file
  │
  └── concrete value

readAll(r io.Reader)
          ▲
          │
       interface
```

```txt
Reader requires:
    Read([]byte) (int, error)

HTTPRequest has:
    Read([]byte) (int, error)

        ↓

    MATCH

        ↓

HTTPRequest satisfies Reader
```

## Abstractions

```go
func copyData(src os.File) {}
```

=> consumer is coupling with `os.File`. chỉ nhận File.

```go
func copyData(src io.Reader) {}
```

=> consumer is coupling with behavior `Read`.

```txt
request.Body
    │
    ├── Read()
    └── Close()
    │
    ▼
io.ReadCloser
    │
    └── also satisfies io.Reader

  ||
  
request.Body
     │
     │ implements/satisfies
     ▼
io.Reader
```

```txt
CALLER
  main
   │
┌─────────────┴─────────────┐
│                           │
▼                           ▼
os.Stdout                    &bytes.Buffer
│                           │
│ satisfies                │ satisfies
▼                           ▼
io.Writer                    io.Writer
│                           │
└─────────────┬─────────────┘
   │
   ▼
writeReport()
CONSUMER
   │
   │ uses
   ▼
io.Writer
   │
   ▼
fmt.Fprintf()
   │
   ▼
Write()
```

```txt
src
 │
 │ concrete value
 ▼
*strings.Reader
 │
 │ satisfies
 ▼
io.Reader
 │
 │
 ▼
┌──────────────────────┐
│       io.Copy        │
│                      │
│ src.Read()           │
│      ↓               │
│    bytes             │
│      ↓               │
│ dst.Write()          │
└──────────────────────┘
          │
          │
          ▼
      os.Stdout
          │
          │ satisfies
          ▼
      io.Writer
```

```txt
io.Reader
┌─────────────────────┐
│ "Tôi cần Read()"    │
└─────────────────────┘

File          🔌 Read
HTTP Body     🔌 Read
Buffer        🔌 Read
StringReader  🔌 Read
GzipReader    🔌 Read

func process(r io.Reader)
```

```txt
io.Writer
┌─────────────────────┐
│ "Tôi cần Write()"   │
└─────────────────────┘

File          🔌 Write
Stdout        🔌 Write
Buffer        🔌 Write
HTTP Response 🔌 Write
GzipWriter    🔌 Write

             io.Copy
          ┌───────────┐
          │           │
Reader 🔌─┤           ├─🔌 Writer
          │           │
          └───────────┘
```

```txt
INTERFACE
io.Reader
  │
defines contract
  │
  ▼
Read(...)
  ▲
  │
must provide
  │
┌─────────┼─────────┐
│         │         │
▼         ▼         ▼
File       Buffer    HTTP Body
│         │         │
└─────────┼─────────┘
  │
implementations
  │
  ▼
CONSUMER
io.Copy(...)
  │
  │ receives
  ▼
src io.Reader
```

```txt
io.Writer
    │
defines contract
    │
    ▼
 Write(...)
    ▲
    │
┌────────┼─────────┐
│        │         │
▼        ▼         ▼
File     Stdout    Buffer
│        │         │
└────────┼─────────┘
    │
    ▼
 CONSUMER
io.Copy(...)
    │
    ▼
dst io.Writer
```

### Rules

1. Implementation provide behavior.
2. Interface describes behavior.
3. Consumer depend only on behavior.
4. Caller choose the implementation.
