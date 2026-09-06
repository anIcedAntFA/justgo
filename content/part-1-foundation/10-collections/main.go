package main

import (
	"fmt"
)

func main() {
	// var a [3]int // [0 0 0] — three ints, zero-valued
	// a[0] = 10
	// a[1] = 20
	// fmt.Println(a) // [10 20 0]
	// fmt.Println(len(a))

	// b := [3]string{"a", "b", "c", "4"}
	// fmt.Println(b) // [a b c]

	// c := []int{1, 2, 3, 4, 5}
	// fmt.Println(c)

	// nums := []int{1, 2, 3, 4}
	// fmt.Println(nums)
	// fmt.Printf("%T\n", nums)

	// var empty []int
	// fmt.Println(empty)
	// fmt.Println(empty == nil)

	// s := make([]int, 3)
	// fmt.Println(s)

	// var a []int
	// b := []int{}
	// fmt.Println(a == nil, b == nil)
	// fmt.Println(a, b)

	// v, err := json.Marshal(a)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(v))

	// v, err = json.Marshal(b)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(v))

	// nums := []int{1, 2, 3}
	// nums = append(nums, 4)
	// nums = append(nums, 5, 6)
	// fmt.Println("nums: ", nums)
	// nums2 := append(nums, 7)
	// fmt.Println("nums2: ", nums2)
	// fmt.Println("nums: ", nums)

	// s := make([]int, 3, 10)
	// fmt.Println("s: ", s)
	// fmt.Println("len(s): ", len(s))
	// fmt.Println("cap(s): ", cap(s))
	// s = append(s, 1, 2)
	// fmt.Println("s: ", s)
	// fmt.Println("len(s): ", len(s))
	// fmt.Println("cap(s): ", cap(s))

	// s2 := make([]int, 3, 4)
	// fmt.Println("s2: ", s2)
	// fmt.Println("len(s2): ", len(s2))
	// fmt.Println("cap(s2): ", cap(s2))
	// s2 = append(s2, 1, 2)
	// fmt.Println("s2: ", s2)
	// fmt.Println("len(s2): ", len(s2))
	// fmt.Println("cap(s2): ", cap(s2))

	// original := []int{1, 2, 3, 4}
	// fmt.Println("original: ", original)
	// sub := original[1:3]
	// fmt.Println("sub: ", sub)
	// sub[0] = 0
	// fmt.Println("sub: ", sub)
	// fmt.Println("original: ", original)

	// s := []int{0, 1, 2, 3, 4, 5}
	// fmt.Println(s[2:4]) // 2, 3
	// fmt.Println(s[:3])  // 0, 1, 2
	// fmt.Println(s[4:])  // 4, 5
	// fmt.Println(s[:])   // 0, 1, 2 ,3, 4 ,5

	// original := []int{1, 2, 3, 4}
	// // Fix 1: an explicit independent copy
	// fmt.Println("original1: ", original) // [1, 2, 3, 4]
	// sub := slices.Clone(original[0:2])
	// fmt.Println("original2: ", original) // [1, 2, 3, 4]
	// fmt.Println("sub: ", sub)            // [1, 2]
	// sub = append(sub, 99)
	// fmt.Println("sub after append: ", sub) // [1, 2, 99]
	// fmt.Println("original3: ", original)   // [1, 2, 3, 4]

	// // Fix 2: the three-index full slice expression s[low:high:max] caps the result
	// fmt.Println("original1: ", original) // [1, 2, 3, 4]
	// sub2 := original[0:2:2]
	// sub2 = append(sub2, 99)
	// fmt.Println("sub2 after append: ", sub2) // [1, 2]
	// fmt.Println("original2: ", original)     // [1, 2, 3, 4]

	// s := []int{1, 2, 3}
	// fmt.Println("s before clear: ", s) // [1 2 3]
	// clear(s)
	// fmt.Println("s after clear: ", s) // [0 0 0]

	// s := []int{3, 1, 4, 1, 5, 9, 2, 6}
	// slices.Sort(s)
	// fmt.Println("s: ", s)
	// fmt.Println(slices.Contains(s, 4)) // true
	// fmt.Println(slices.Index(s, 5))
	// fmt.Println(slices.Max(s))
	// fmt.Println(slices.Min(s))
	// slices.Reverse(s)
	// fmt.Println("s after reverse: ", s)

	// clone := slices.Clone(s)
	// fmt.Println("s == clone: ", slices.Equal(s, clone))

	// ages := map[string]int{"khoi": 24, "mai": 20, "dudu": 4, "bear": 3}
	// scores := make(map[string]int)
	// scores["ngockhoi96"] = 8
	// fmt.Println("ages: ", ages)
	// fmt.Println(ages["not exist"]) // 0
	// fmt.Println("scores: ", scores)

	// if age, ok := ages["khoi"]; ok {
	// 	fmt.Println("age:", age)
	// }

	// for key, value := range ages { // ⚠️ order is RANDOM — never rely on it
	// 	fmt.Printf("%s: %d\n", key, value)
	// }

	// for _, key := range slices.Sorted(maps.Keys(ages)) {
	// 	fmt.Printf("%s: %d\n", key, ages[key])
	// }

	type Point struct{ X, Y int }
	m := map[string]Point{"a": {1, 2}}
	// m["a"].X = 10
	fmt.Println("m: ", m)
	p := m["a"]
	p.X = 10
	m["a"] = p
	fmt.Println("m: ", m)

	m2 := map[string]*Point{"a": {1, 2}}
	m2["a"].X = 10
	fmt.Println("m2: ", *m2["a"])
}
