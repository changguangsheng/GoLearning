package main

import (
	"fmt"
	"time"
)

func main() {
	// 声明一个退出用的通道
	exit := make(chan int)

	fmt.Println("start")
	// 过 1 秒后, 调用匿名函数
	//调用 time.AfterFunc() 函数，传入等待的时间和一个回调。回调使用一个匿名函数，在时间到达后，匿名函数会在另外一个
	//goroutine 中被调用。
	time.AfterFunc(time.Second*10, func() {
		// 1 秒后, 打印结果
		fmt.Println("one second after")
		// 通知 main()的 goroutine 已经结束
		exit <- 0
	})
	//等待结束
	<-exit
}
