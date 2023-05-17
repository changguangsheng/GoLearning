package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func doTask(n int) {
	time.Sleep(time.Duration(n))
	fmt.Printf("Task %d Done\n", n)
	wg.Done()
}
func main() {

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go doTask(i + 1)
	}
	// 如何优雅的实现结束子goroutine
	wg.Wait()
	fmt.Println("All Task Done!")
}
