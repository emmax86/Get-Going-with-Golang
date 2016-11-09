package main

import "fmt"

func Yell(s string) string {
	return s + "!"
}

func Swap(x int, y int) (int, int) {
	return y, x
}

func Say(s string) {
	fmt.Println(s)
}

func main() {
	Say(Yell("Steve"))
	fmt.Println(Swap(37, 24))
}
