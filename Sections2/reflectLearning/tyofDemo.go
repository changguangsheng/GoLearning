package main

import (
	"fmt"
	"reflect"
)

// Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。
type myInt int64

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	//fmt.Printf("type:%v\n", v)
	fmt.Printf("type:%v kind:%v\n", v.Name(), v.Kind())
}

func main() {
	var a float32 = 3.14
	reflectType(a)
	var b int64 = 100
	reflectType(b)

	var c *float32 // 指针
	var d myInt    // 自定义类型
	var e rune     // 类型别名
	reflectType(c)
	reflectType(d)
	reflectType(e)
	type person struct {
		name string
		age  int
	}
	type book struct {
		title string
	}
	var f = person{
		name: "小王子",
		age:  18,
	}
	var g = book{title: "learn go"}
	reflectType(f)
	reflectType(g)

}
