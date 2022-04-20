package main

import (
	"com.grpk.test/src/queue"
)

func main() {
	arr1 := ListNode{
		1,
		&ListNode{
			4,
			&ListNode{
				5,
				nil,
			},
		},
	}

	arr2 := ListNode{
		1,
		&ListNode{
			3,
			&ListNode{
				4,
				nil,
			},
		},
	}

	arr3 := ListNode{
		2,
		&ListNode{
			6,
			nil,
		},
	}

	arrList := []*ListNode{&arr1, &arr2, &arr3}

	queue := queue.New()

	for _, value := range arrList {

		if value == nil {
			next
		}

		temp := value
		for true {
			queue.AddElement(temp.Val)

			if temp.Next == nil {
				break
			}

			temp = temp.Next
		}
	}

	result := &ListNode{
		0,
		nil,
	}
	next := result

	for true {
		min := queue.PopMin()

		if min == -1 {
			break
		}

		next.Next = &ListNode{
			min,
			nil,
		}

		next = next.Next
	}

	return result.Next

}

type ListNode struct {
	Val  int
	Next *ListNode
}
