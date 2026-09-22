package main

import "fmt"

func main() {
	// 1. ARRAY

	a := [3]int{1, 2, 3}

	// print the type of a:
	fmt.Printf("type of a: %T\n", a)

	// try to append to array (would not compile, so commented out)
	// a = append(a, 4)
	//

	printSlice(a[:])

	// 2. SLICE
	b := []int{1, 2, 3}
	fmt.Printf("type of b: %T\n", b)

	b = append(b, 4)
	printSlice(b)

	// 3. SLICE via make

	c := make([]int, 3, 3) // even if we set length, capacity - the result will be SLICE
	fmt.Printf("type of c: %T\n", c)
	c = append(c, []int{1, 2, 3, 4, 5}...)
	printSlice(c)
}

func printSlice(a []int) {
	for i, v := range a {
		fmt.Printf("idx: %d, value: %d, type: %T\n", i, v, v)
	}
}
