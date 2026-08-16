// Command countdown counts down from 10 to 1, one second apart, then a deferred
// call prints the launch message when main returns — a tiny demo of defer + for.
//
// Run it from this directory:
//
//	go run .
package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("🚀 Launch!")

	for i := 10; i >= 1; i-- {
		fmt.Printf("%d...\n", i)
		time.Sleep(time.Second)
	}
}
