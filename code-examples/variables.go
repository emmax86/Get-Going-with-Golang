package main

import "fmt"

func main() {
	// Declare x and initialize it now
	var x int = 13

	// Declare y and initialize it later
	var y int
	y = 12

	// Declare z and infer its type
	z := x * y
	fmt.Println(z)
}
