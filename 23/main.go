package main

import "container/heap"

func main() {
	l := []*ListNode{}
	l = append(l, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	})
	l = append(l, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	})
	l = append(l, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
		},
	})

	mergeKLists(l)

}

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

type Pq []*ListNode

func (p Pq) Len() int {
	return len(p)
}

func (p Pq) Swap(a, b int) {
	(p)[a], (p)[b] = (p)[b], (p)[a]
}

func (p Pq) Less(a, b int) bool {
	return p[a].Val < p[b].Val
}

func (p *Pq) Pop() interface{} {
	if len(*p) == 0 {
		return nil
	}
	t := (*p)[0]
	*p = (*p)[1:]
	return t
}

func (p *Pq) Push(x interface{}) {
	*p = append(*p, x.(*ListNode))
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := &Pq{}
	heap.Init(pq)
	for _, l := range lists {
		curr := l
		for curr != nil {
			heap.Push(pq, curr)
			curr = curr.Next
		}
	}

	dummy := &ListNode{}
	curr := dummy
	for len(*pq) > 0 {
		n := heap.Pop(pq).(*ListNode)
		curr.Next = n
		curr = curr.Next
	}

	return dummy.Next
}
