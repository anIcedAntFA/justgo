// Command embedding demonstrates interface embedding: io.ReadWriteCloser is just
// io.Reader + io.Writer + io.Closer composed together. A type satisfies it by
// having Read, Write, AND Close — and it can then be passed anywhere any of the
// smaller interfaces is expected.
//
// Run it from this directory:
//
//	go run .
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

// memFile is an in-memory value that satisfies io.ReadWriteCloser: it embeds a
// *bytes.Buffer (which already provides Read and Write) and adds Close.
type memFile struct {
	*bytes.Buffer
	closed bool
}

func (m *memFile) Close() error {
	m.closed = true
	return nil
}

// process accepts the compound interface, then hands the value to functions that
// each want only a smaller piece of it — Writer, then Reader.
func process(rwc io.ReadWriteCloser) error {
	if _, err := io.WriteString(rwc, "written through io.Writer\n"); err != nil {
		return err
	}
	data, err := io.ReadAll(rwc) // rwc is also an io.Reader
	if err != nil {
		return err
	}
	fmt.Printf("read back through io.Reader: %q\n", string(data))
	return rwc.Close()
}

func main() {
	// A *memFile satisfies io.ReadWriteCloser via embedding.
	var f io.ReadWriteCloser = &memFile{Buffer: new(bytes.Buffer)}
	if err := process(f); err != nil {
		log.Fatal(err)
	}
	fmt.Println("done")
}
