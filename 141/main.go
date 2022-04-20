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

func hasCycle(head *ListNode) bool {
	p1, p2 := head, head

	for p1 != nil && p1.Next != nil {
		p1 = p1.Next.Next
		p2 = p2.Next

		if p1 == p2 {
			return true
		}
	}

	return false
}
