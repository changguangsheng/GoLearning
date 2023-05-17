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
	MaxNum := 100
	rand.Seed(time.Now().UnixNano())
	TruthNum := rand.Intn(MaxNum)
	
	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input, Please try again", err)
			return
		}
		input = strings.TrimSuffix(input, "\n")
		guessNum, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("An error occured while reading input, Please try again", err)
			return
		}
		//fmt.Scan(&guessNum)
		println("You guess is ", guessNum)
		switch {
		case guessNum > TruthNum:
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		case guessNum < TruthNum:
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		default:
			fmt.Println("Correct, you Legend!")
			return
		}
	}
}
