package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		//从3循环到0
		for i := 3; i >= 0; i-- {
			//发送3到0之间的数值
			ch <- i
			//每次发送完时等待
			time.Sleep(time.Second)
		}
	}()

	for data := range ch {
		fmt.Println(data)

		if data == 0 {
			break
		}
	}
}
