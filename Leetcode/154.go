package leetcode

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		middlen := left + (right-left)/2
		if nums[middlen] < nums[right] {
			right = middlen
		} else if nums[middlen] > nums[right] {
			left = middlen + 1
		} else {
			right -= 1
		}
	}
	return nums[right]
}
