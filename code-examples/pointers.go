package main

import (
	"fmt"
)

func main() {
	x := 4
	y := &x
	*y = 1
	z := new(int)
	*z = 10
	fmt.Printf("X: %d \n&X: %p \nY: %p \n*Y: %d \nZ: %p \n*Z: %d", x, &x, y, *y, z, *z)
}
