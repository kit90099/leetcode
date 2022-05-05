package main

import (
	"fmt"

	. "com.grpk.leet/utils/LinkedListUtils"
)

func main() {
	ans := isPalindrome(GenFromSlice([]int{1, 2, 2, 1}))
	fmt.Println(ans)
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	reversedP2 := reverse(slow)
	p1, p2 := head, reversedP2
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}

		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func reverse(node *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := node

	for curr != nil {
		next := curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
}
