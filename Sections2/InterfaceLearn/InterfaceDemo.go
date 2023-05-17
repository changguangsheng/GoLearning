package main

import "fmt"

type Person interface {
	getName() string
}
type Student struct {
	name string
	age  int
}

func (s *Student) getName() string {
	return s.name
}

type Worker struct {
	name   string
	gender string
}

func (w *Worker) getName() string {
	return ""
}

// 将空值 nil 转换为 *Student 类型，再转换为 Person 接口，如果转换失败，说明 Student 并没有实现 Person 接口的所有方法。
var _ Person = (*Student)(nil)

var _ Person = (*Worker)(nil)

func main() {
	var p Person = &Student{
		name: "Tom",
		age:  18,
	}
	fmt.Println(p.getName())
}
