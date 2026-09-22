package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	for i := range 10 {
		fmt.Printf("i: %02d, random value: %2d\n", (i + 1), randInt(1, 10))
	}
}

func randInt(min int, max int) int {
	return rand.IntN(max-min+1) + min
}
