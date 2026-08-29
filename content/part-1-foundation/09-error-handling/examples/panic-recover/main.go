// Command panic-recover shows the ONE legitimate everyday use of recover: stopping
// a panic in one unit of work from crashing the whole program. It is not try/catch —
// expected failures should be returned as errors, not panicked.
package main

import "fmt"

// safeRun runs fn and converts any panic into a returned error, using a named
// return so the deferred closure can set it. This is the seed of HTTP recovery
// middleware (Chapter 16).
func safeRun(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()
	fn()
	return nil
}

func main() {
	// A panicking task is contained — the program keeps running.
	err := safeRun(func() {
		panic("something exploded")
	})
	fmt.Println("task 1:", err) // task 1: recovered from panic: something exploded

	// A well-behaved task returns nil.
	err = safeRun(func() {
		fmt.Println("task 2: doing work")
	})
	fmt.Println("task 2:", err) // task 2: <nil>

	fmt.Println("main still running") // reached — the panic did not crash us
}
