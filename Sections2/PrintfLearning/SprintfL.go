package main

import (
	"fmt"
	"reflect"
)

// Sprint系列函数会把传入的数据生成并返回一个字符串。
func main() {
	s1 := fmt.Sprint("沙河小王子")
	name := "沙河小王子"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("沙河小王子")
	fmt.Println(s1, s2, s3)
	fmt.Println(reflect.TypeOf(s1).Name())
}
