package main

import "fmt"

func main() {
	var (
		name    string
		age     int
		married bool
	)
	//fmt.Scanf不同于fmt.Scan简单的以空格作为输入数据的分隔符，fmt.Scanf为输入数据指定了具体的输入内容格式，只有按照格式输入数据才会被扫描并存入对应变量。
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name：%s  age：%d   married：%t \n", name, age, married)
}
