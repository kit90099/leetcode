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

func main() {
	node := genFromSlice([]int{1, 2, 3})
	ans := reverseList(node)
	fmt.Print(ans)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
