package exercises

import "io"

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
	return nil // TODO: replace
}
