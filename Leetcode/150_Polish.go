package Leetcode

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, token := range tokens {
		num, err := translamb(token)
		if err == false {
			num, stack = operate(stack, token)
		}
		stack = push(stack, num)
		// fmt.Println(num, stack)
	}
	return stack[0]
}
func operate(stack []int, token string) (int, []int) {
	t2, stack := pull(stack)
	t1, stack := pull(stack)
	switch token {
	case "+":
		return t1 + t2, stack
	case "-":
		return t1 - t2, stack
	case "*":
		return t1 * t2, stack
	case "/":
		return int(t1 / t2), stack
	default:
		return 0, stack
	}
}

func push(stack []int, num int) []int {
	return append(stack, num)
}

func pull(stack []int) (int, []int) {
	length := len(stack)
	// fmt.Println("length", length)
	num := stack[length-1]
	stack = stack[:length-1]
	return num, stack
}

// 将字符串转换成int
func translamb(token string) (int, bool) {
	if token == "+" || token == "-" || token == "*" || token == "/" {
		return 0, false
	}
	length := len(token)
	ans := 0
	op := 1
	for i := 0; i < length; i++ {
		if token[i] == '-' {
			op = -1
			continue
		}
		ans = ans*10 + int(token[i]-'0')
	}
	return ans * op, true
}
