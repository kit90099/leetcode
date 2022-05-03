package LinkedListUtils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenFromSlice(s []int) *ListNode {
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

func PrintListNode(node *ListNode) {
	curr := node
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}
