package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("co len")

	// err := process5()
	// fmt.Println("err:", err)

	// err := process4()
	// fmt.Println("3. error:", err)

	// process3()
	// fmt.Println("3. after process")

	// c := counter()
	// fmt.Println(c(), c(), c())

	// c2 := counter()
	// fmt.Println(c2())

	// funcs := []func(){}
	// for i := 0; i < 3; i++ {
	// 	funcs = append(funcs, func() { fmt.Print(i, " ") })
	// }
	// for _, f := range funcs {
	// 	f()
	// }

	// count := 0
	// f1 := func() {
	// 	fmt.Println(count)
	// }
	// f2 := func() {
	// 	count++
	// }
	// f2()
	// f1()

	// GreetWorld()

	// server := NewServer(WithPort(9000), WithTimeout(100*time.Second))
	// fmt.Println(server)

	// triple := func(n int) int {
	// 	return n * 3
	// }
	// subtract5 := func(n int) int {
	// 	return n - 4
	// }
	// divide3 := func(n int) int {
	// 	return n / 2
	// }
	// fmt.Println("transform(4, triple, subtract5, divide3)", transform(4, triple, subtract5, divide3))

	// Closure carrying state across calls.
	tag := tagger("row-")
	fmt.Println(tag("alpha")) // row-1:alpha
	fmt.Println(tag("beta"))  // row-2:beta
}

func tagger(prefix string) func(string) string {
	count := 0
	return func(label string) string {
		count++
		return fmt.Sprintf("%s%d:%s", prefix, count, label)
	}
}

func transform(n int, fns ...func(int) int) int {
	for _, f := range fns {
		n = f(n)
	}
	return n
}

type ServerConfig struct {
	Port    int
	Timeout time.Duration
	MaxConn int
}

type Option func(*ServerConfig)

func WithPort(port int) Option {
	return func(c *ServerConfig) { c.Port = port }
}

func WithTimeout(t time.Duration) Option {
	return func(c *ServerConfig) { c.Timeout = t }
}

func NewServer(opts ...Option) *ServerConfig {
	config := &ServerConfig{ // defaults
		Port:    8080,
		Timeout: 30 * time.Second,
		MaxConn: 100,
	}
	for _, opt := range opts { // apply overrides
		opt(config)
	}
	return config
}

func Greet(name string) {
	fmt.Printf("Hello, %s\n", name)
}

func GreetWorld() {
	Greet("World")
}

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func process() {
	fmt.Println("1. start")

	panic("something went wrong")

	// fmt.Println("2. end") // never executed — panic stops execution here
}

func process2() {
	defer fmt.Println("cleanup")

	fmt.Println("1. start")

	panic("something went wrong")

	// fmt.Println("2. end") // never executed — panic stops execution here
}

func process3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered", r)
		}
	}()

	fmt.Println("1. start")

	panic("something went wrong")

	// fmt.Println("2. end") // never executed — panic stops execution here
}

func process4() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()

	fmt.Println("1. start")

	panic("something went wrong")

	// fmt.Println("2. end") // never executed — panic stops execution here
	// return nil
}

func process5() (err error) {
	defer func() {
		err = fmt.Errorf("defer changed the error")
	}()

	err = fmt.Errorf("defer changed the error")

	return nil
}
