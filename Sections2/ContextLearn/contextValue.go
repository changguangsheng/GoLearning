package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

type Options struct {
	Interval time.Duration
}

func reqTask2(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			//在子协程中，使用 ctx.Value("options") 获取到传递的值，读取/修改该值。
			op := ctx.Value("options").(*Options)
			fmt.Println(reflect.TypeOf(op))
			time.Sleep(time.Second * op.Interval)
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//context.WithValue() 创建了一个基于 ctx 的子 Context，并携带了值 options。
	vctx := context.WithValue(ctx, "options", &Options{1})

	go reqTask2(vctx, "worker1")
	go reqTask2(vctx, "worker2")

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}
