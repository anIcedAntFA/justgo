package exercises

// Limiter returns a closure that allows up to max successful calls — each of the
// first max calls returns nil — and returns a non-nil error on every call after
// that. The remaining budget is state captured by the closure.
//
//	allow := Limiter(2)
//	allow() // nil
//	allow() // nil
//	allow() // error: limit exceeded
//
// TODO: capture a counter in the returned closure and compare against max.
func Limiter(max int) func() error {
	return nil // TODO: replace
}
