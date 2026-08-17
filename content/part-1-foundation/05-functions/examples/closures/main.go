// Command closures shows state captured and kept alive by closures: an
// independent counter per constructor call, and a memoized Fibonacci whose
// cache persists across calls because the returned function closes over it.
//
// Run it from this directory:
//
//	go run .
package main

import "fmt"

// counter returns a function that increments and returns its own private count.
// Each call to counter() creates a fresh count, closed over by the returned func.
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// memoizedFib returns a Fibonacci function backed by a cache it closes over, so
// repeated calls reuse earlier results instead of recomputing them.
func memoizedFib() func(int) int {
	cache := map[int]int{}

	var fib func(int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		if v, ok := cache[n]; ok {
			return v
		}
		result := fib(n-1) + fib(n-2)
		cache[n] = result
		return result
	}
	return fib
}

func main() {
	c := counter()
	fmt.Println("counter c:", c(), c(), c()) // 1 2 3

	c2 := counter()
	fmt.Println("counter c2:", c2()) // 1 — independent of c

	fib := memoizedFib()
	fmt.Println("fib(10):", fib(10)) // 55
	fmt.Println("fib(40):", fib(40)) // 102334155 — fast, thanks to the cache
}
