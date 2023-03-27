package main

import (
	"errors"
	"fmt"
	"time"
)

func RPCClint(ch chan string, req string) (string, error) {
	//向服务器发送请求
	ch <- req

	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second): //超时
		return "", errors.New("Time out")
	}
}

// 模拟 RPC 服务器端接收客户端请求和回应
func RPCServer(ch chan string) {
	for {
		//接收客户端请求
		data := <-ch
		//打印接收到的数据
		fmt.Println("server received:", data)
		time.Sleep(time.Second * 2)
		ch <- "roger"
	}

}

func main() {
	ch := make(chan string)
	//并发执行服务器逻辑
	go RPCServer(ch)
	recv, err := RPCClint(ch, "hi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client received", recv)
	}
}
