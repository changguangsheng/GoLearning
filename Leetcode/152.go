package leetcode

func maxProduct1(nums []int) int {
	length := len(nums)
	ans := 1
	negNumber := make([]int, 0)
	for i := 0; i < length; i++ {
		if nums[i] == 0 {
			return max(maxProduct1(nums[:i]), maxProduct1(nums[i+1:]))
		}
		ans *= nums[i]
		if nums[i] < 0 {
			negNumber = append(negNumber, i)
		}
	}

	if ans > 0 {
		return ans
	}
	left := 1
	for i := 0; i <= negNumber[0]; i++ {
		left *= nums[i]
	}
	right := 1
	for i := length - 1; i >= negNumber[len(negNumber)]; i-- {
		right *= nums[i]
	}
	ans = ans / (left * right / max(left, right))
	return ans
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func min(num1, num2 int) int {
	if num1 > num2 {
		return num2
	}
	return num1
}

// 以当前的nums[i]作为子序列的末尾数字， 之前的最大 和最小
func maxProduct(nums []int) int {
	ans, maxn, minn := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxn, minn
		maxn = max(nums[i]*mx, max(nums[i]*mn, nums[i]))
		minn = min(nums[i]*mn, min(nums[i]*mx, nums[i]))
		ans = max(maxn, minn)
	}
	return ans
}
