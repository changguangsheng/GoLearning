package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	//第一种
	//flag.Type(flag名, 默认值, 帮助信息)*Type
	//name := flag.String("name", "张三", "姓名")
	//age := flag.Int("age", 18, "年龄")
	//married := flag.Bool("married", false, "婚否")
	//delay := flag.Duration("d", 0, "时间间隔")
	//fmt.Println(*name, *age, *married, *delay)

	//第二种
	//flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")
	fmt.Println(name, age, married)
}
