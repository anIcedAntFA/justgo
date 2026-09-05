// Command json-roundtrip is the Chapter 14 "encoding/json" example: marshal a
// struct to a file with indentation, then decode it back. Note the struct tags
// controlling the JSON keys, omitempty dropping the zero-value field, and that
// only EXPORTED fields cross the wire. This is exactly how gorg persists its
// rules config and undo journal.
//
// Run it from this directory:
//
//	go run .
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Rule is one classification rule. Struct tags rename the fields to snake_case
// JSON keys; omitempty drops Note when it is the empty string. An unexported
// field here would be invisible to encoding/json.
type Rule struct {
	Category   string   `json:"category"`
	Extensions []string `json:"extensions"`
	Note       string   `json:"note,omitempty"`
}

func main() {
	rules := []Rule{
		{Category: "Images", Extensions: []string{".jpg", ".png", ".gif"}},
		{Category: "Documents", Extensions: []string{".pdf", ".md"}, Note: "text-ish"},
	}

	path := filepath.Join(os.TempDir(), "gorg-rules-example.json")

	// Encode: MarshalIndent gives a human-readable config file.
	data, err := json.MarshalIndent(rules, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "marshal:", err)
		os.Exit(1)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		fmt.Fprintln(os.Stderr, "write:", err)
		os.Exit(1)
	}
	fmt.Printf("wrote %s:\n%s\n\n", path, data)

	// Decode: read it back into a fresh value and confirm the round-trip.
	back, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "read:", err)
		os.Exit(1)
	}
	var loaded []Rule
	if err := json.Unmarshal(back, &loaded); err != nil {
		fmt.Fprintln(os.Stderr, "unmarshal:", err)
		os.Exit(1)
	}
	fmt.Printf("decoded %d rules, first category = %s\n", len(loaded), loaded[0].Category)
}
