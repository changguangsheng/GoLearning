package main

import (
	"fmt"
	"time"
)

func running() {
	var times int
	//构建一个无限循环
	for {
		times++
		fmt.Println("tick", times)

		//延迟1s
		time.Sleep(time.Second)
	}
}

func main() {
	go running()

	var input string
	fmt.Scanln(&input)
}