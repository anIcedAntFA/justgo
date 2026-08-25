// Command io-writer demonstrates the power of the small io.Writer interface: one
// function writes to ANY destination — the terminal, an in-memory buffer, a
// compressed stream — because they all satisfy io.Writer. io.Copy does the same
// across io.Reader → io.Writer.
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
	"os"
	"strings"
)

// writeReport writes a formatted report to any io.Writer. The caller chooses the
// destination; this function neither knows nor cares what it is.
func writeReport(w io.Writer, lines []string) error {
	for i, line := range lines {
		if _, err := fmt.Fprintf(w, "%d. %s\n", i+1, line); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lines := []string{"first item", "second item", "third item"}

	// Destination 1: the terminal.
	fmt.Println("--- to stdout ---")
	if err := writeReport(os.Stdout, lines); err != nil {
		log.Fatal(err)
	}

	// Destination 2: an in-memory buffer.
	var buf bytes.Buffer
	if err := writeReport(&buf, lines); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("--- captured in memory (%d bytes) ---\n%s", buf.Len(), buf.String())

	// io.Copy connects any io.Reader to any io.Writer. Always check its returns.
	fmt.Println("--- io.Copy: string reader -> stdout ---")
	src := strings.NewReader("copied straight through io.Copy\n")
	if _, err := io.Copy(os.Stdout, src); err != nil {
		log.Fatal(err)
	}
}
