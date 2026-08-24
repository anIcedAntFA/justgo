package main

import (
	"fmt"
)

func main() {
	// u1 := User{
	// 	Name:  "khoi",
	// 	Email: "haha@gmail.com",
	// 	Age:   22,
	// }
	// u2 := User{
	// 	"toan",
	// 	"hehe@gmail.com",
	// 	12,
	// }
	// var u3 User
	// u4 := User{Name: "thien"}
	// u5 := &User{Name: "con"}
	// fmt.Println(u1)
	// fmt.Println(u2)
	// fmt.Println(u3)
	// fmt.Println(u4)
	// fmt.Println(u5)

	// u := User{Name: "trinh", Age: 24}
	// fmt.Println(u.Name)
	// u.Age = 26
	// fmt.Println(u.Age)
	// fmt.Println("original struct", u)

	// p := &u
	// fmt.Println(p.Name)
	// p.Age = 21
	// fmt.Println("original struct", u)

	// u := User{Name: "Khoi", Email: "nk@gmail.com", Age: 0, Pass: "secret"}
	// data, _ := json.Marshal(u)
	// fmt.Println(string(data))

	// rect := Rectangle{
	// 	Width:  10,
	// 	Height: 5,
	// }
	// fmt.Println(rect.Area())
	// fmt.Println(rect.Perimeter())

	rect := Rectangle{Width: 5, Height: 10}
	fmt.Println("Before scaling:", rect)
	rect.Scale(2)
	fmt.Println("After scaling:", rect)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// func (r Rectangle) Scale(factor float64) {
// 	r.Width *= factor // modifies the COPY, not the original
// 	r.Height *= factor
// }

func (r Rectangle) Area() float64 {
	return r.Width + r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age,omitempty"`
	Pass  string `json:"-"`
}
