package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHello() {
	fmt.Printf("Hello! My name is %s. I am %d years old.\n", p.Name, p.Age)
}

func main() {
	p := Person{Name: "Steve", Age: 20}
	p.SayHello()
}
