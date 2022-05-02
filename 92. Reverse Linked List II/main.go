package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func genFromSlice(s []int) *ListNode {
	res := ListNode{}
	nexPtr := &res

	for _, val := range s {
		nexPtr.Next = &ListNode{
			Val: val,
		}
		nexPtr = nexPtr.Next
	}

	return res.Next
}

func printListNode(node *ListNode) {
	curr := node
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

func main() {
	node := genFromSlice([]int{3})
	ans := reverseBetween(node, 1, 1)
	printListNode(ans)
}

// ver1
/* func reverseBetween(head *ListNode, left int, right int) *ListNode {
	curr := head
	lBoundary := head
	for i := 1; i < left; i++ {
		lBoundary = curr
		curr = curr.Next
	}

	newHead, _ := reverse(curr, right-left)

	if left == 1 {
		head = newHead
	} else {
		lBoundary.Next = newHead
	}

	return head
} */

// ver 2
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	newHead := &ListNode{
		Next: head,
	}

	curr := newHead
	for i := 0; i < left-1; i++ {
		curr = curr.Next
	}

	lHead, rBoundary := reverse(curr.Next, right-left)
	curr.Next.Next = rBoundary
	curr.Next = lHead

	return newHead.Next
}

func reverse(head *ListNode, count int) (*ListNode, *ListNode) {
	if head.Next == nil || count == 0 {
		return head, head.Next
	}
	newParent, rBoundary := reverse(head.Next, count-1)
	head.Next.Next = head
	return newParent, rBoundary
}
