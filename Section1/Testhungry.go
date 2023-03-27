package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(3)
	var wg sync.WaitGroup

	const runtime = time.Second
	var shareLock sync.Mutex

	greedyWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			shareLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			shareLock.Unlock()
			count++
		}
		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	}
	politeWorker := func() {
		defer wg.Done()
		var count int
		for bigin := time.Now(); time.Since(bigin) <= runtime; {
			shareLock.Lock()
			time.Sleep(time.Nanosecond)
			shareLock.Unlock()

			shareLock.Lock()
			time.Sleep(time.Nanosecond)
			shareLock.Unlock()

			shareLock.Lock()
			time.Sleep(time.Nanosecond)
			shareLock.Unlock()
			count++
		}
		fmt.Printf("Polite worker was able to execute %v work loops\n", count)
	}
	wg.Add(2)
	go greedyWorker()
	go politeWorker()

	wg.Wait()

}
