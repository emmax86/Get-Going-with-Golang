package main

import "fmt"

func foo(ch chan int) {
	fmt.Println("will block")
	ch <- 3
	ch <- 4
	close(ch)
}

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	go foo(ch)
	for i := range ch {
		fmt.Println(i)
	}

}
