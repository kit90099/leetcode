package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result

	l1p := l1
	l2p := l2

	for l1p != nil && l2p != nil {
		if l1p.Val >= l2p.Val {
			p.Next = l2p
			l2p = l2p.Next
		} else {
			p.Next = l1p
			l1p = l1p.Next
		}
		p = p.Next
	}

	if l1p == nil {
		p.Next = l2p
	}
	if l2p == nil {
		p.Next = l1p
	}

	return result.Next
}

func main() {
	l1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	fmt.Print(mergeTwoLists(&l1, &l2))
}
