package main

import "fmt"

type Digit int

const (
	Zero Digit = iota
	One
	Two
	Three
	Four
	Five
)

func main() {
	fmt.Printf("typeof Zero: %T, value: %v\n", Zero, Zero)
	var d Digit
	fmt.Printf("typeof d: %T\n", d)
}
