package exercises

// Result holds EITHER a value of type T or an error — a type-safe alternative to
// Go's (T, error) pair, inspired by Rust's Result. It shows a generic type whose
// zero-of-T handling matters.
type Result[T any] struct {
	value T
	err   error
}

// Ok builds a successful Result carrying value.
//
// TODO: return a Result[T] with value set and err nil.
func Ok[T any](value T) Result[T] {
	return Result[T]{} // TODO
}

// Err builds a failed Result carrying err. The value stays at T's zero value.
//
// TODO: return a Result[T] with err set.
func Err[T any](err error) Result[T] {
	return Result[T]{} // TODO
}

// IsOk reports whether the Result holds a value (no error).
//
// TODO: return r.err == nil.
func (r Result[T]) IsOk() bool {
	return false // TODO
}

// Unwrap returns the value. It panics if the Result is an error — call IsOk
// first, or prefer UnwrapOr.
//
// TODO: if r.err != nil, panic(r.err); otherwise return r.value.
func (r Result[T]) Unwrap() T {
	var zero T
	return zero // TODO
}

// UnwrapOr returns the value on success, or fallback when the Result is an error.
//
// TODO: return r.value when IsOk, else fallback.
func (r Result[T]) UnwrapOr(fallback T) T {
	return fallback // TODO
}
