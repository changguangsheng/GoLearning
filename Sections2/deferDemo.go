package main

import "fmt"

func main() {
	if true {
		defer fmt.Println("1")
	} else {
		defer fmt.Println("2")
	}
	defer fmt.Println("3")
}
