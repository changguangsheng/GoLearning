package main

import "log"

/*
logger会打印每条日志信息的日期、时间，默认输出到系统的标准错误。
Fatal系列函数会在写入日志信息后调用os.Exit(1)。Panic系列函数会在写入日志信息后panic。
*/
func main() {
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")
	log.Fatalln("这是一条会触发fatal的日志。")
}
