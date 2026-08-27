package exercises

// Settings uses pointer fields to distinguish three states per option:
// nil means "not set — use the default", while a non-nil pointer means
// "explicitly set" — even when it points at the zero value (0 or false). A plain
// int/bool field could not tell "not set" apart from "set to zero".
type Settings struct {
	Timeout *int  // nil = use default
	Verbose *bool // nil = use default
}

// Resolved is the effective configuration after defaults are applied.
type Resolved struct {
	Timeout int
	Verbose bool
}

// Resolve applies s over the defaults (Timeout=30, Verbose=false): a nil field
// keeps the default; a non-nil field overrides it, even if it points at the
// zero value.
//
// TODO: implement — for each field, keep the default when the pointer is nil,
// otherwise dereference and use the pointed-to value.
func Resolve(s Settings) Resolved {
	// TODO: implement
	return Resolved{}
}
