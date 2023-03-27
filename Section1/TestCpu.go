package main

import "fmt"

func main() {
	for i := 0; i < 2; i++ {
		fmt.Println(":12312")
		go AsyncFunc(i)
	}
	fmt.Println(":12312dd")
}
func AsyncFunc(index int) {
	fmt.Println("123123123123")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Printf("线程%d, sum为:%d\n", index, sum)
}
