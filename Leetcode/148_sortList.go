package leetcode

import "fmt"

type listNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}
func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	fmt.Println(mid.Val)
	return merge(sort(head, mid), sort(mid, tail))
}
func merge(node1, node2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  -1,
		Next: nil,
	}
	dummy := head
	for node1 != nil && node2 != nil {
		if node1.Val >= node2.Val {
			dummy.Next = node2
			node2 = node2.Next
		} else {
			dummy.Next = node1
			node1 = node1.Next
		}
		dummy = dummy.Next
	}
	if node1 == nil {
		dummy.Next = node2
	}
	if node2 == nil {
		dummy.Next = node1
	}
	return head.Next
}
