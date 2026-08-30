package main

import (
	"fmt"
)

func main() {
	// x := 20
	// p := &x

	// fmt.Println(x)  // value
	// fmt.Println(p)  // memory address
	// fmt.Println(*p) // * means value at this address - dereferencing

	// x1 := 28
	// p1 := &x1        // &: address of -> p is a *int pointing at x
	// fmt.Println(x1)  // value
	// fmt.Println(p1)  // memor address
	// fmt.Println(*p1) // *: value at dereference p1
	// *p1 = 24
	// fmt.Println(*p1)
	// fmt.Println(p1)
	// fmt.Println(x1)

	// maihu := &Person{Name: "mai", Age: 28}
	// fmt.Println(maihu)
	// fmt.Println((*maihu).Name)
	// // auto dereference
	// fmt.Println(maihu.Name)
	// maihu.Name = "mailun"
	// fmt.Println(maihu.Name)

	// x := 10
	// increment(&x)
	// fmt.Println(x)

	// acc := Account{Balance: 100}
	// deposit(&acc, 100)
	// fmt.Println(acc.Balance)
}

type Account struct {
	Balance float64
}

func deposit(acc *Account, amount float64) {
	fmt.Println(acc)
	fmt.Println(&acc)
	(*acc).Balance += amount
}

func increment(n *int) {
	// get value at address that n points to, then increment
	*n++
}

type Person struct {
	Name string
	Age  int
}
