package main

import (
	"fmt"
	"sync"
)

var (
	count      int
	countGuard sync.Mutex
	//countsGuard sync.RWMutex
)

func GetCount() int {
	countGuard.Lock()

	defer countGuard.Unlock()

	return count
}

func SetCount(c int) {
	countGuard.Lock()

	count = c
	countGuard.Unlock()
}
func main() {
	SetCount(1)
	fmt.Println(GetCount())
}
