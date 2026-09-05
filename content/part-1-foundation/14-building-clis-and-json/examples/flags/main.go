// Command flags is the Chapter 14 "flag package" example: define typed flags,
// parse them, and read the leftover positional arguments. The stdlib flag
// package is enough for a real CLI — no cobra needed in Part 1.
//
// Run it from this directory:
//
//	go run . --name Go --loud extra args
//	go run . -h
package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	// Define flags: type, name, default, and help text. Each returns a pointer.
	name := flag.String("name", "World", "who to greet")
	loud := flag.Bool("loud", false, "shout the greeting")

	// Parse must be called after all flags are defined and before they are read.
	flag.Parse()

	greeting := fmt.Sprintf("Hello, %s!", *name)
	if *loud {
		greeting = strings.ToUpper(greeting)
	}
	fmt.Println(greeting)

	// flag.Args() holds the positional (non-flag) arguments left over.
	if rest := flag.Args(); len(rest) > 0 {
		fmt.Printf("leftover args: %v\n", rest)
	}
}
