package exercises

// Exercise 4: Recovery Middleware (Preview of Part 2).
//
// safeExecute runs fn and turns a panic into a returned error, so one bad call can't
// crash the program. This is the seed of the HTTP recovery middleware in Chapter 16.
//
// Behaviour:
//   - fn panics        → return an error whose message is
//     "recovered from panic: <value>"
//   - fn returns an err → return that err unchanged
//   - fn returns nil    → return nil
//
// TODO: use a deferred closure with recover() and a named return value so the
// closure can set the returned error after a panic.
func safeExecute(fn func() error) (err error) {
	// TODO: implement
	return fn()
}
