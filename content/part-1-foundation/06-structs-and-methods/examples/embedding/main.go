// Command embedding shows composition through struct embedding: promoted fields
// and methods, embedding more than one type, and an outer method shadowing an
// embedded one.
//
// Run it from this directory:
//
//	go run .
package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Eat() {
	fmt.Printf("%s is eating\n", a.Name)
}

func (a Animal) Describe() string {
	return "some animal named " + a.Name
}

// Dog embeds Animal (composition, not inheritance) — Animal's Name and Eat are
// promoted onto Dog. Dog also SHADOWS Describe with its own version.
type Dog struct {
	Animal
	Breed string
}

func (d Dog) Bark() {
	fmt.Printf("%s barks!\n", d.Name) // Name promoted from Animal
}

func (d Dog) Describe() string {
	return d.Name + " the " + d.Breed // shadows Animal.Describe
}

// Logger and Validator are embedded together into Service to show multiple
// embedding — Service gains both capabilities.
type Logger struct{ prefix string }

func (l Logger) Log(msg string) { fmt.Printf("%s: %s\n", l.prefix, msg) }

type Validator struct{}

func (v Validator) Validate() bool { return true }

type Service struct {
	Logger
	Validator
	name string
}

func main() {
	d := Dog{
		Animal: Animal{Name: "Rex"},
		Breed:  "Labrador",
	}

	d.Eat()                          // promoted from Animal
	d.Bark()                         // Dog's own method
	fmt.Println(d.Name)              // promoted field
	fmt.Println(d.Describe())        // Dog's Describe wins (shadowing)
	fmt.Println(d.Animal.Describe()) // reach the embedded one explicitly

	s := Service{
		Logger: Logger{prefix: "SVC"},
		name:   "payment",
	}
	s.Log("starting " + s.name)         // from Logger
	fmt.Println("valid:", s.Validate()) // from Validator
}
