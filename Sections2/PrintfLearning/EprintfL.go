package main

import (
	"errors"
	"fmt"
)

//函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。

func main() {
	err := fmt.Errorf("这是一个错误")
	e := errors.New("meta error w")
	w := fmt.Errorf("wrap了一个错误%w", e)
	fmt.Println("cuwu:", w, err)
}
