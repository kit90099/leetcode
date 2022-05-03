package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import (
	. "com.grpk.leet/utils/LinkedListUtils"
)

func main() {
	ans := reverseKGroup(GenFromSlice([]int{}), 3)
	PrintListNode(ans)
}

/**
* r: decide whether need return
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	newHead := &ListNode{
		Next: head,
	}
	curr := newHead
	r := newHead
	for curr.Next != nil {
		for i := 0; i < k; i++ {
			if r.Next == nil {
				return newHead.Next
			}
			r = r.Next
		}
		next := r.Next

		nextHead := reverse(curr.Next, k-1)
		// curr: 0
		temp := curr.Next    // old head eg 1
		curr.Next = nextHead // update to new head eg 0->2->1
		curr = temp          // new curr: old head
		r = curr
		curr.Next = next // connect new tail(old head) to next part
	}
	return newHead.Next
}

func reverse(head *ListNode, count int) *ListNode {
	if head.Next == nil || count == 0 {
		return head
	}

	newHead := reverse(head.Next, count-1)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
