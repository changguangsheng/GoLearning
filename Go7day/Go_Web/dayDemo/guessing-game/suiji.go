package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("请输入一个数")
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\r')
		if err != nil {
			fmt.Println("请重新输入")
			continue
		}
		input = strings.TrimSuffix(input, "\r")
		if input == "-1" {
			return
		}
		guessNum, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("请重新输入")
			continue
		}
		flag := rand.Intn(100)
		if flag%2 == 0 {
			flag = -1
		} else {
			flag = 1
		}
		randomFloat := (rand.Float64() * 0.5)
		output1 := guessNum + randomFloat*float64(flag)

		flag = rand.Intn(100)
		if flag%2 == 0 {
			flag = -1
		} else {
			flag = 1
		}
		randomFloat = (rand.Float64() * (0.5 - randomFloat))
		output2 := guessNum + randomFloat*float64(flag)

		output3 := guessNum*3. - output2 - output1
		fmt.Println(output1, output2, output3)
		fmt.Println("请输入一个数（退出请输入-1）")
	}
}
