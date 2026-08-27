package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"
)

func main() {
	// shapes := []Shape{
	// 	Rectangle{Width: 4, Height: 5},
	// 	Circle{Radius: 3},
	// }
	// for _, s := range shapes {
	// 	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
	// }

	// temp := Temperature(24.4)
	// fmt.Println(temp)

	// var a fmt.Stringer = Temperature(25.5)
	// var b DisplayTemp = Temperature(25.5)
	// fmt.Println(a.String())
	// fmt.Println(b.String())

	dump(strings.NewReader("Hello, World!\nThis is a test.\n"))
	dump(bytes.NewBufferString("hello"))

	file := File{}
	copyData(file)
	// Go watch copyData(src Reader)
	// and argument file -> File
	// Go check does File satisfy Reader?
	// Yes => compile
	req := HTTPRequest{}
	copyData(req)

	// Go watch copyData(src Reader)
	// and argument body -> Body -> io.ReadCloser -> io.Reader
	// Go check does body satisfy Reader
	rq := http.Request{}
	body := rq.Body
	copyData(body)

	var i any = "so sad"
	// s := i.(int)
	// fmt.Println(s)
	s, ok := i.(string)
	if ok {
		fmt.Println(s)
	}
	n, ok := i.(int)
	if !ok {
		fmt.Println("not an int", n)
	}

	// err := validateAge(-12)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	c := Color{R: 255, G: 128, B: 0}
	fmt.Println(c)

	err := doSomething()
	if err != nil {
		fmt.Println("error is not nil!") // THIS RUNS — surprising!
	}

	fmt.Println(math.Mod(3661, 3600))
}

type MyError struct{}

func (e *MyError) Error() string { return "my error" }

func doSomething() error {
	var err *MyError = nil // a nil pointer
	return err             // ⚠️ returns a non-nil interface!
}

type Color struct {
	R, G, B uint8
}

func (c Color) String() string {
	return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}

type ValidationError struct {
	Field   string
	Message string
}

// Implement the error interface
func (ve *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %s:%s", ve.Field, ve.Message)
}

// Now *ValidationError IS an error, usable anywhere an error is expected
func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{
			Field:   "age",
			Message: "must be non-negative",
		}
	}
	return nil
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

// Implementations
type File struct{}

func (f File) Read(p []byte) (int, error) {
	// ...
	return 0, nil
}

// Implementations
type HTTPRequest struct{}

func (r HTTPRequest) Read(p []byte) (int, error) {
	// ...
	return 0, nil
}

// Consumer of Interface Reader
// Reader is Interface parameter type
func copyData(src Reader) {
	// ...
}

func dump(r io.Reader) {
	buf := make([]byte, 4096)

	for {
		n, err := r.Read(buf)

		if n > 0 {
			fmt.Print(string(buf[:n]))
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}
	}
}

type Temperature float64

func (t Temperature) String() string {
	return fmt.Sprintf("%.1foC", float64(t))
}

type DisplayTemp interface {
	String() string
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
