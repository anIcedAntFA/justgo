package exercises

import (
	"fmt"
	"io"
)

// WriteReport writes a numbered report to w — one line per item, 1-indexed:
//
//	[]string{"alpha", "beta"} → "1. alpha\n2. beta\n"
//
// Accepting io.Writer is the whole point: the same function writes to a file, a
// bytes.Buffer, or os.Stdout without knowing which.
//
// TODO: implement with fmt.Fprintf (add the import); return the first write
// error, or nil on success.
func WriteReport(w io.Writer, data []string) error {
	for i, s := range data {
		if _, err := fmt.Fprintf(w, "%d. %s\n", i+1, s); err != nil {
			return err
		}
	}
	return nil
}

// WriteReport -> Consumer, w -> dep, data -> data need writing
// w -> write ability
// `io.Writer` -> Contract
//
// for mỗi item
//     lấy index
//     tạo string "<index + 1>. <item>\n"
//     ghi string đó vào w
//     nếu lỗi → return lỗi ngay
// return nil
//
// func Fprintf(w io.Writer, format string, a ...any) (n int, err error) {}
// -> format text and write to Writer
//
// Printf -> write to terminal
// Fprint -> write `io.Writer` -> let caller decide Destination
//
// terminal
//
// data
//  │
//  ▼
// WriteReport
//  │
//  │ Fprintf
//  ▼
// os.Stdout
//  │
//  ▼
// terminal
//
// WriteReport does not know terminal exist!
//
// Dependency Injection
//
// look signature
// func WriteReport(w io.Writer, data []string) error
// WriteReport does not create writer:
// BAD for this use case
// func WriteReport(data []string) error {
//     file, _ := os.Create(...)
// }
//
// it was injected a writer:
// caller
//   │
//   │ chooses implementation
//   ▼
// os.Stdout / file / buffer
//   │
//   │ inject
//   ▼
// WriteReport(w io.Writer, ...)
//
// Caller decide dependency, Consumer only need Contract
//
// 2 Consumers here:
// func WriteReport(w io.Writer, data []string) error
// func Fprintf(w io.Writer, format string, a ...any) (n int, err error)
//
// WriteReport
//    │
//    │ uses
//    ▼
// io.Writer
//    ▲
//    │
//    │ implementation
//    │
// os.Stdout
//
// WriteReport
//     │
//     │ calls
//     ▼
// fmt.Fprintf
//     │
//     │ consumes
//     ▼
// io.Writer
//
// => WriteReport receives Interface and give that Interface to other API that known Interface.
// -> Composibility
//
// Mental execution:
//
// Caller
//   │
//   │ &bytes.Buffer
//   ▼
// WriteReport
//   │
//   │ w = io.Writer
//   │
//   │ for
//   ▼
// fmt.Fprintf(w, ...)
//   │
//   │ calls Write(...)
//   ▼
// bytes.Buffer
//   │
//   ▼
// memory
//
// Both WriteReport and Fprintf does not know bytes.Buffer
// Both only need Write()
//
// 1. What is consumer? -> WriteReport
// 2. What capabilities do consumers need? -> Write
// 3. Which contract provides that capability? -> io.Writer
// 4. Algorithm ->
// loop
//   format index + item
// 	 write
//   error -> return
// success -> nil
// 5. API format and write -> fmt.Fprintf
//
// CALLER
//   │
//   │ chooses concrete implementation
//   │
// ┌──────────┼───────────┐
// ▼          ▼           ▼
// os.Stdout     File     bytes.Buffer
// │          │           │
// └──────────┼───────────┘
//   │
//   │ satisfies
//   ▼
// io.Writer
// CONTRACT
//   │
//   │ injected into
//   ▼
// WriteReport
// CONSUMER
//   │
//   │ calls
//   ▼
// fmt.Fprintf
//   │
//   │ uses Writer.Write()
//   ▼
// actual destination
