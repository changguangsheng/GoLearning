package main

import "fmt"

func main() {
	ch := make(chan int, 4)

	fmt.Println(len(ch))

	ch <- 2
	ch <- 1
	ch <- 3
	fmt.Println(len(ch))
	close(ch)
}
