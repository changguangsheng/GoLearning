package main

import (
	"fmt"
	"os"
)

// Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
func main() {
	fmt.Fprintln(os.Stdout, "向标准输出写内容")
	fileObj, err := os.OpenFile("./write.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}
	name := "沙河小王子"
	fmt.Fprintf(fileObj, "往文件中写信息：%s", name)
}
