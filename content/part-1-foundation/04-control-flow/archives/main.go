package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// age := 18

	// if age >= 18 {
	// 	fmt.Println("adult")
	// } else {
	// 	fmt.Println("not adult")
	// }

	// if err := doSomething(); err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }

	// for i := 0; i < 6; i++ {
	// 	fmt.Println(i)
	// }

	// n := 1
	// for n < 20 {
	// 	n *= 2
	// }
	// fmt.Println(n)

	// names := []string{"Alice", "Bob", "Charlie"}
	// for i := range names {
	// 	fmt.Printf("%d\n", i)
	// }

	// for i, ch := range "Hello, 世界" {
	// 	fmt.Printf("%d, %c\n", i, ch)
	// }

	// ages := map[string]int{"Alice": 30, "Bob": 25}
	// for name, age := range ages {
	// 	fmt.Printf("%s is %d years old\n", name, age)
	// }

	// for i := range 9 {
	// 	fmt.Println(i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// for i := 1; i <= 20; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// day := "Tuesday"
	// switch day {
	// case "Monday":
	// 	fmt.Println("Start of week")
	// case "Tuesday", "Wednesday", "Thursday": // multiple values in one case
	// 	fmt.Println("Midweek")
	// 	fallthrough
	// case "Friday":
	// 	fmt.Println("Almost weekend")
	// default:
	// 	fmt.Println("Weekend")
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Linux")
	// case "darwin":
	// 	fmt.Println("Mac OS")
	// default:
	// 	fmt.Printf("%s\n", os)
	// }

	// age := 25
	// switch {
	// case age < 13:
	// 	fmt.Println("child")
	// case age < 18:
	// 	fmt.Println("teenager")
	// case age < 65:
	// 	fmt.Println("adult")
	// default:
	// 	fmt.Println("senior")
	// }

	// fmt.Println("Start")
	// defer fmt.Println("first defer")
	// defer fmt.Println("second defer")
	// defer fmt.Println("third defer")
	// fmt.Println("End")

	// x := 5
	// defer fmt.Println(x)
	// x = 20
	// fmt.Println(x)

	// x := 10
	// defer func() {
	// 	fmt.Println(x) // reads x when the deferred func runs
	// }()
	// x = 20
	// fmt.Println(x)

	// Timing a function
	// start := time.Now()
	// defer func() {
	// 	fmt.Printf("took %v\n", time.Since(start))
	// }()

	// process()

	// ⚠️ range evaluates the slice's length once, at the start
	nums := []int{1, 2, 3, 4, 5}
	for _, n := range nums {
		if n == 2 {
			nums = append(nums, 4)
		}
	}
	fmt.Println(nums)
}

func process() {
	start := time.Now()

	defer func() {
		fmt.Printf("took %v\n", time.Since(start))
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Processing done")
}

func readFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	// ... read file, process content ...
	// even if something panics, file.Close() still runs
	return nil
}

func doSomething() error {
	return nil
}
