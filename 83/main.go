package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}

	fmt.Print(*deleteDuplicates(&node))
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head

	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = fast
		}

		fast = fast.Next
	}
	slow.Next = nil
	return dummy.Next
}
