package main

import (
	"fmt"
	"sync"
	"time"
)

// 全局变量方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易统一
// 2. 如果worker中再启动goroutine，就不太好控制了。

var wg1 sync.WaitGroup
var exit bool

func worker1() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
	wg1.Done()
}

func main() {
	wg1.Add(1)
	go worker1()
	time.Sleep(time.Second * 3) // sleep3秒以免程序过快退出
	exit = true                 // 修改全局变量实现子goroutine的退出
	wg1.Wait()
	fmt.Println("over")
}
