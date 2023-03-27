package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type value struct {
	memAccess sync.Mutex
	value     int
}

func main() {
	runtime.GOMAXPROCS(3)
	var wg sync.WaitGroup
	sum := func(v1, v2 *value) {
		defer wg.Done()
		v1.memAccess.Lock()
		time.Sleep(2 * time.Second)
		v2.memAccess.Lock()
		fmt.Printf("sum = %d\n", v1.value+v2.value)
		v2.memAccess.Unlock()
		v1.memAccess.Unlock()
	}

	product := func(v1, v2 *value) {
		defer wg.Done()
		v2.memAccess.Lock()
		time.Sleep(2 * time.Second)
		v1.memAccess.Lock()
		fmt.Printf("product = %d\n", v1.value*v2.value)
		v1.memAccess.Unlock()
		v2.memAccess.Unlock()
	}

	var v1, v2 value
	v1.value = 1
	v2.value = 1
	wg.Add(2)
	go sum(&v1, &v2)
	go product(&v1, &v2)
	wg.Wait()
}
