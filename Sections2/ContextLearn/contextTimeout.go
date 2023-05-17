package main

import (
	"context"
	"fmt"
	"time"
)

func reqTask4(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(time.Second)
		}
	}
}
func main() {
	//使用 context.WithTimeout 创建具有超时通知机制的 Context 对象。
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	go reqTask4(ctx, "worker1")
	go reqTask4(ctx, "worker2")

	time.Sleep(3 * time.Second)
	fmt.Println("before cancel")
	cancel()
	time.Sleep(3 * time.Second)
}
