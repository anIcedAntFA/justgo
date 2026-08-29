package exercises

import (
	"errors"
	"testing"
)

func TestParseErrorAs(t *testing.T) {
	t.Skip("Chapter 09 exercise: implement ParseError and parse, then delete this Skip")

	err := parse("x = = 1")
	if err == nil {
		t.Fatal("parse() = nil, want a wrapped *ParseError")
	}

	// errors.AsType (Go 1.26) pulls the *ParseError out of the wrapping, comma-ok
	// style, so we can read its fields. (errors.As(err, &pe) works the same on older Go.)
	pe, ok := errors.AsType[*ParseError](err)
	if !ok {
		t.Fatalf("errors.AsType failed to extract *ParseError from %v", err)
	}

	if pe.Line != 2 {
		t.Errorf("Line = %d, want 2", pe.Line)
	}
	if pe.Column != 5 {
		t.Errorf("Column = %d, want 5", pe.Column)
	}
}
