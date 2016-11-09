package main

import "fmt"

func main() {
	groceries := make([]string, 3)
	groceries[0] = "Lettuce"
	groceries[1] = "Apples"
	groceries[2] = "Beef"
	fmt.Printf("Size: %d, Slice: %#v\n", len(groceries), groceries)
	groceries = append(groceries, "Milk")
	fmt.Printf("Size: %d, Slice: %#v\n", len(groceries), groceries)
	fmt.Printf("Size: %d, Slice: %#v\n", len(groceries[1:3]), groceries[1:3])
}
