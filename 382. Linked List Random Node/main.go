package main

import "math/rand"

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

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	s := Solution{
		head: head,
	}
	return s
}

func (this *Solution) GetRandom() int {
	i := 0
	res := 0

	ptr := this.head
	for ptr != nil {
		i++

		if rand.Intn(i) == 0 {
			res = ptr.Val
		}
		ptr = ptr.Next
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
