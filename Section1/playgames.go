package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 模拟一个选手在打网球
func player(name string, count chan int) {
	//函数退出时调用Done， 来通知main函数工作已经完成
	defer wg.Done()

	for {
		//等待球被击打过来
		ball, ok := <-count
		if !ok {
			//如果通道被关闭，我们就赢了
			fmt.Printf("player %s won\n", name)
			return

		}
		//选随机数，然后用证书判断我们是否丢球
		n := rand.Intn(100)

		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(count)
			return
		}

		fmt.Printf("player %s Hit %d\n", name, ball)
		ball++
		count <- ball
	}
}
func main() {
	//创建一个无缓冲的通道
	court := make(chan int)
	//计数加2，表示要等待两个goroutine
	wg.Add(2)
	//启动两个选手
	go player("Nadal", court)
	go player("Djokovic", court)
	//发球
	court <- 1
	//等待游戏结束
	wg.Wait()
}
