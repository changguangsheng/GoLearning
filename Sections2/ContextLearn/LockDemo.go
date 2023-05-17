package main

import (
	"sync"
	"time"
)

var (
	x    int64
	lock sync.Mutex
)

func addWithoutLock() {
	for i := 0; i < 2000; i++ {
		x = x + 1
	}
}
func addWithLock() {
	for i := 0; i < 2000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
}
func Add() {
	x = 0
	for i := 0; i < 5; i++ {
		go addWithoutLock()
	}
	time.Sleep(time.Second)
	println("WithoutLock:", x)
	x = 0
	for i := 0; i < 5; i++ {
		go addWithLock()
	}
	time.Sleep(time.Second)
	println("WithLock:", x)
}
func main() {
	Add()
}
