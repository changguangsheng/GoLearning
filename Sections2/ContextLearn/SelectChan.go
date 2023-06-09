package main

import (
	"fmt"
	"time"
)

var stop chan bool

func reqTask(name string) {
	for {
		select {
		case <-stop:
			fmt.Println("Stop ", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}
func main() {
	stop = make(chan bool)
	go reqTask("work1")
	time.Sleep(3 * time.Second)
	stop <- true
	time.Sleep(time.Second * 3)
}
