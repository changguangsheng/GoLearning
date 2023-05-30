package main

import "fmt"

func reverseWords(s string) string {
	i := 0
	length := len(s)
	for ; i < length; i++ {
		if s[i] != ' ' {
			break
		}
	}
	j := length - 1
	for ; j >= i; j-- {
		if s[j] != ' ' {
			break
		}
	}
	s = s[i : j+1]
	strList := make([]string, 0)
	lastIndex := 0
	for i = 0; i < len(s); i++ {
		if s[i] == ' ' {
			strList = append(strList, s[lastIndex:i])
			for ; i < len(s) && s[i+1] == ' '; {
				i++
			}
			lastIndex = i + 1
		}
	}
	strList = append(strList, s[lastIndex:i])
	fmt.Println(s)
	return s
}

func main() {
	s := "123456789"
	s = s[3:5]
	fmt.Println(s)
}
