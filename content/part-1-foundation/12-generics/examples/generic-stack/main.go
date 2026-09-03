// Command generic-stack demonstrates generic TYPES (not just functions): a
// type-safe Stack[T] container with methods, and a two-parameter Pair[K, V].
// Methods restate the receiver's type parameter — func (s *Stack[T]).
//
// Run it from this directory:
//
//	go run .
package main

import "fmt"

// Stack is a generic LIFO container. T can be any type; each Stack[T] instance
// is fixed to one concrete type, checked at compile time.
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop returns the top item and true, or the zero value and false when empty.
// You can't write `return nil` for a generic T — use `var zero T`.
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

// Pair holds two values of possibly different types. K is comparable (so a Pair
// key could be a map key); V is any.
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

func main() {
	intStack := &Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	val, ok := intStack.Pop()
	fmt.Println(val, ok, intStack.Len()) // 2 true 1

	strStack := &Stack[string]{}
	strStack.Push("hello")
	// strStack.Push(42)  // ❌ compile error — Stack[string] only takes strings
	fmt.Println(strStack.Len()) // 1

	p := Pair[string, int]{Key: "age", Value: 30}
	fmt.Println(p.Key, p.Value) // age 30
}
