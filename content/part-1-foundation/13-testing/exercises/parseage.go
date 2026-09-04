package exercises

import "errors"

// Sentinel errors for ParseAge. Tests match these with errors.Is (Chapter 9),
// which is why they're exported values rather than ad-hoc strings.
var (
	// ErrNotNumber means s did not parse as an integer at all.
	ErrNotNumber = errors.New("age is not a number")
	// ErrOutOfRange means s parsed but the value is outside 0..150.
	ErrOutOfRange = errors.New("age out of range")
)

// ParseAge parses s into an age in the inclusive range 0..150.
//
// It returns:
//   - (n, nil)            when s is a valid number in range
//   - (0, ErrNotNumber)   when s is not an integer (wrap the strconv error so
//     errors.Is finds ErrNotNumber AND the underlying cause is preserved)
//   - (0, ErrOutOfRange)  when s is a number but < 0 or > 150
//
// TODO:
//  1. Use strconv.Atoi to parse s.
//  2. On a parse error, return 0 and an error that wraps ErrNotNumber, e.g.
//     fmt.Errorf("%w: %q", ErrNotNumber, s).
//  3. If the value is < 0 or > 150, return 0 and ErrOutOfRange.
//  4. Otherwise return the value and nil.
func ParseAge(s string) (int, error) {
	return 0, nil // TODO
}
