// Command grades is the Chapter 04 capstone: it combines defer timing, range over
// an integer (Go 1.22+), range over a slice, a tagless switch, and an if with an
// init statement — one small program touching every construct in the chapter.
//
// Run it from this directory:
//
//	go run .
package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"time"
)

func main() {
	// defer for timing — runs last, when main returns.
	start := time.Now()
	defer func() {
		fmt.Printf("\n--- completed in %v ---\n", time.Since(start))
	}()

	// for loop + range over integer (Go 1.22+).
	// math/rand/v2's top-level funcs are auto-seeded, so no Seed call needed.
	fmt.Println("Generating 5 random scores...")
	scores := make([]int, 5)
	for i := range 5 {
		scores[i] = rand.IntN(101) // 0-100
	}

	// range over slice.
	total := 0
	for _, score := range scores {
		total += score
	}
	avg := float64(total) / float64(len(scores))

	// tagless switch for grade.
	var grade string
	switch {
	case avg >= 90:
		grade = "A"
	case avg >= 80:
		grade = "B"
	case avg >= 70:
		grade = "C"
	case avg >= 60:
		grade = "D"
	default:
		grade = "F"
	}

	fmt.Printf("Scores: %v\n", scores)
	fmt.Printf("Average: %.1f → Grade: %s\n", avg, grade)

	// if with init statement.
	if f, err := os.CreateTemp("", "scores-*.txt"); err != nil {
		fmt.Println("Error creating temp file:", err)
	} else {
		defer func() { _ = f.Close() }()
		if _, werr := fmt.Fprintf(f, "Scores: %v\nAverage: %.1f\nGrade: %s\n", scores, avg, grade); werr != nil {
			fmt.Println("Error writing temp file:", werr)
		}
		fmt.Printf("Saved to: %s\n", f.Name())
	}
}
