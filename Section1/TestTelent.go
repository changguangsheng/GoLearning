package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 服务逻辑, 传入地址和退出的通道
func server(address string, exitChan chan int) {
	// 根据给定地址进行侦听
	l, err := net.Listen("tcp", address)
	// 如果侦听发生错误, 打印错误并退出
	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}
	// 打印侦听地址, 表示侦听成功
	fmt.Println("listen: " + address)
	// 延迟关闭侦听器
	defer l.Close()
	// 侦听循环
	for {
		// 新连接没有到来时, Accept 是阻塞的
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		// 根据连接开启会话, 这个过程需要并行执行
		go HandleSession(conn, exitChan)
	}
}

func HandleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("Session started:")

	reader := bufio.NewReader(conn)

	for {
		// 读取字符串, 直到碰到回车返回
		str, err := reader.ReadString('\n')
		if err == nil {
			// 去掉字符串尾部的回车
			str = strings.TrimSpace(str)
			// 处理 Telnet 指令
			if !processTelnetCommande(str, exitChan) {
				conn.Close()
				break
			}
			// Echo 逻辑, 发什么数据, 原样返回
			conn.Write([]byte(str + "\r\n"))

		} else {
			//发生错误
			fmt.Println("Session closed")
			conn.Close()
			break
		}

	}
}

func processTelnetCommande(str string, exitChan chan int) bool {
	//@close 指令表示终止本次会话
	if strings.HasPrefix(str, "@close") {
		fmt.Println("Session closed")
		//告诉外部需要断开连接
		return false
		//@shutdown 指令表示终止服务进程
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("Server shutdown")
		//往通道中写入 0, 阻塞等待接收方处理
		exitChan <- 0
		return false
	}
	// 打印输入的字符串
	fmt.Println(str)
	return true
}

func main() {
	//创建一个程序结束码的通道
	exitChan := make(chan int)
	//将服务器并发运行
	go server("127.0.0.1:7001", exitChan)
	//通道阻塞, 等待接收返回值
	code := <-exitChan
	// 标记程序返回值并退出
	os.Exit(code)
}
