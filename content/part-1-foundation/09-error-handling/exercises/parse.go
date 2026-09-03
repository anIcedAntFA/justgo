package exercises

import (
	"fmt"
)

// Exercise 3: Custom Error Type with errors.As.
//
// Define a ParseError that carries structured data (line, column, message). Give it
// Error() and Unwrap(), then a parser that returns a ParseError wrapped in extra
// context. The test uses errors.As to pull the ParseError back out and read its
// fields — the payload a plain sentinel can't carry.

// ParseError describes a failure at a specific position in some input. The fields
// are given; your job is to implement the methods and the parser below.
type ParseError struct {
	Line    int
	Column  int
	Message string
	Err     error // the wrapped underlying error, if any
}

// Error implements the error interface. Format it however you like, but include the
// line and column, e.g. "parse error at line 2, column 5: unexpected token".
//
// TODO: implement.
func (e *ParseError) Error() string {
	return fmt.Sprintf("parse error at line %v, column %v: %v", e.Line, e.Column, e.Message)
}

// Unwrap returns the wrapped underlying error so errors.Is/As can traverse it.
//
// TODO: return the wrapped error field.
func (e *ParseError) Unwrap() error {
	return e.Err
}

// parse always fails at line 2, column 5. It returns a *ParseError wrapped in
// additional context with fmt.Errorf("%w"), so the caller must use errors.As (not a
// type assertion) to recover it.
//
// TODO: build a *ParseError{Line: 2, Column: 5, Message: "unexpected token"} and
// return it wrapped, e.g. fmt.Errorf("parsing %q: %w", input, pe).
func parse(input string) error {
	pe := &ParseError{
		Line:    2,
		Column:  5,
		Message: "unexpected token",
	}
	return fmt.Errorf("parsing %q: %w", input, pe)
}

// parse("x = = 1")
//         │
//         ▼
// create ParseError
//         │
//         │ Line = 2
//         │ Column = 5
//         │ Message = "unexpected token"
//         ▼
//       pe
//         │
//         │ %w
//         ▼
// fmt.Errorf("parsing %q: %w", input, pe)
//         │
//         ▼
//       return
//         │
//         ▼
//        err
//         │
//         │ errors.AsType[*ParseError]
//         ▼
//    traverse chain
//         │
//         │ Unwrap()
//         ▼
//    *ParseError
//         │
//         ├── Line   = 2
//         ├── Column = 5
//         └── Message
