package main

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := ListNode{
		Val:  -1,
		Next: head,
	}

	p1 := &dummy
	p2 := &dummy
	size := 0
	for p1.Next != nil {
		size++
		if size > n {
			p2 = p2.Next
		}
		p1 = p1.Next
	}

	p2.Next = p2.Next.Next

	return dummy.Next
}
