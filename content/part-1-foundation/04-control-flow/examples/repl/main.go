// Command repl is a tiny Read-Eval-Print Loop: an infinite for loop reads a line
// from stdin and a switch dispatches on the command. Type "quit" (or Ctrl-D) to
// exit — a demo of the infinite for + switch combo.
//
// Run it from this directory:
//
//	go run .
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("mini-repl — commands: hello, time, quit")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break // EOF (Ctrl-D)
		}

		switch strings.TrimSpace(scanner.Text()) {
		case "quit":
			fmt.Println("bye 👋")
			return
		case "time":
			fmt.Println(time.Now().Format(time.Kitchen))
		case "hello":
			fmt.Println("Hello, Gopher!")
		case "":
			// ignore empty input
		default:
			fmt.Printf("unknown command: %s\n", scanner.Text())
		}
	}
}
