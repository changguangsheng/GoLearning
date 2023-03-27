package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个打点器, 每 500 毫秒触发一次
	ticker := time.NewTicker(time.Millisecond * 500)
	//创建一个计时器，2s后触发
	stopper := time.NewTimer(time.Second * 2)

	var i int

	for {
		select {
		case <-stopper.C:
			fmt.Println("stop")
			goto StopHere
		case <-ticker.C:
			i++
			fmt.Println("tick", i)
		}
	}
StopHere:
	fmt.Println("dne")
}
