package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	//count int32
	count int32
	wg    sync.WaitGroup
)

func main() {
	wg.Add(2)
	go inCount()
	go inCount()
	wg.Wait()
	fmt.Println(count)
}
func inCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		runtime.Gosched()
		value++
		count = value
	}
}
