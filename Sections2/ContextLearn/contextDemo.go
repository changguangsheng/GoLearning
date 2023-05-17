package main

import (
	"context"
	"fmt"
	"time"
)

func reqTask1(ctx context.Context, name string) {
	for {
		select {
		//在子协程中，使用 select 调用 <-ctx.Done() 判断是否需要退出。
		case <-ctx.Done():
			fmt.Println("stop ", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}
func main() {
	//context.Backgroud() 创建根 Context，通常在 main 函数、初始化和测试代码中创建，作为顶层 Context。
	//context.WithCancel(parent) 创建可取消的子 Context，同时返回函数 cancel。
	ctx, cancel := context.WithCancel(context.Background())
	go reqTask1(ctx, "worker2")
	go reqTask1(ctx, "worker3")
	time.Sleep(3 * time.Second)
	//主协程中，调用 cancel() 函数通知子协程退出。
	cancel()
	time.Sleep(3 * time.Second)
}
