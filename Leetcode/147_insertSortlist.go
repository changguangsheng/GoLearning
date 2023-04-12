package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	nextNode := &ListNode{Next: head}
	dummy := head

	for dummy.Next != nil {

		if dummy.Next.Val >= dummy.Val {
			dummy = dummy.Next
		} else {
			lastNode := nextNode
			nowNode := dummy.Next
			for nowNode.Val >= lastNode.Next.Val {
				lastNode = lastNode.Next
			}

			dummy.Next = nowNode.Next
			nowNode.Next = lastNode.Next
			lastNode.Next = nowNode
		}
	}
	return nextNode.Next
}
