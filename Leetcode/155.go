package leetcode

import "math"

type MinStack struct {
	stacks    []int
	minstacks []int
}

func Constructor() MinStack {
	return MinStack{
		stacks:    make([]int, 0),
		minstacks: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(val int) {
	this.stacks = append(this.stacks, val)
	top := this.minstacks[len(this.minstacks)-1]
	this.minstacks = append(this.minstacks, min(val, top))
}

func (this *MinStack) Pop() {
	this.stacks = this.stacks[:len(this.stacks)-1]
	this.minstacks = this.minstacks[:len(this.minstacks)-1]
}

func (this *MinStack) Top() int {
	return this.stacks[len(this.stacks)-1]
}

func (this *MinStack) GetMin() int {
	return this.minstacks[len(this.minstacks)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
