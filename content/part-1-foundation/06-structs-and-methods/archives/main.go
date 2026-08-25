package main

import "fmt"

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

	// rect := Rectangle{Width: 5, Height: 10}
	// fmt.Println("Before scaling:", rect)
	// rect.Scale(2)
	// fmt.Println("After scaling:", rect)

	// c := Counter{count: 1}
	// fmt.Println("Initial count:", c.count)
	// c.Increment() // auto does (&c).Increment()
	// fmt.Println(c.count)

	// server, err := NewServer("localhost", 8080)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// d := Dog{
	// 	Animal: Animal{Name: "de"},
	// 	Breed:  "blabla",
	// }
	// d.Eat()
	// d.Bark()
	// fmt.Println(d.Name)
	// fmt.Println(d.Breed)

	// d := Derived{}
	// fmt.Println(d.Describe())
	// fmt.Println(d.Base.Describe())

	// config := struct {
	// 	Host string
	// 	Port int
	// }{
	// 	Host: "localhost",
	// 	Port: 8080,
	// }
	// fmt.Println(config.Host)

	p1 := Point{1, 2}
	p2 := Point{1, 2}
	p3 := Point{2, 3}
	fmt.Println(p1 == p2)
	fmt.Println(p1 == p3)

	seen := map[Point]bool{}
	seen[Point{1, 2}] = true
	seen[Point{3, 4}] = true
	fmt.Println(seen[Point{1, 2}])
	fmt.Println(seen[Point{3, 4}])
}

type Point struct {
	X, Y int
}

// type Base struct{}
// func (b Base) Describe() string {
// 	return "I am base struct"
// }
// type Derived struct {
// 	Base
// }
// func (d Derived) Describe() string {
// 	return "I am derived struct"
// }

// type Animal struct {
// 	Name string
// }
// func (a Animal) Eat() {
// 	fmt.Printf("%s is eating\n", a.Name)
// }
// type Dog struct {
// 	Animal
// 	Breed string
// }
// func (d Dog) Bark() {
// 	fmt.Printf("%s barks!\n", d.Name)
// }

// type Server struct {
// 	host    string
// 	port    int
// 	timeout time.Duration
// }

// func NewServer(host string, port int) (*Server, error) {
// 	if port < 1 || port > 65535 {
// 		return nil, fmt.Errorf("invalid port: %d", port)
// 	}

// 	return &Server{
// 		host:    host,
// 		port:    port,
// 		timeout: 60 * time.Second,
// 	}, nil
// }

// type Counter struct{ count int }
// func (c *Counter) Increment() {
// 	c.count++
// }

// type Rectangle struct {
// 	Width  float64
// 	Height float64
// }
// func (r *Rectangle) Scale(factor float64) {
// 	r.Width *= factor
// 	r.Height *= factor
// }
// // func (r Rectangle) Scale(factor float64) {
// // 	r.Width *= factor // modifies the COPY, not the original
// // 	r.Height *= factor
// // }
// func (r Rectangle) Area() float64 {
// 	return r.Width + r.Height
// }
// func (r Rectangle) Perimeter() float64 {
// 	return 2 * (r.Width + r.Height)
// }

// type User struct {
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// 	Age   int    `json:"age,omitempty"`
// 	Pass  string `json:"-"`
// }
