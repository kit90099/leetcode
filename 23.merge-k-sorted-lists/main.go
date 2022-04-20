package main

import "fmt"

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

func main() {
	/* list1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}

	list2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	list3 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
		},
	} */

	list := []*ListNode{}

	ans := mergeKLists(list)
	fmt.Println(ans)
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}

	pq := []int{}

	for _, list := range lists {
		next := list

		if list == nil {
			continue
		}

		for true {
			insert(&pq, (*next).Val)

			if next.Next == nil {
				break
			}

			next = next.Next
		}
	}

	res := ListNode{
		Val: -1,
	}
	var ptr *ListNode = &res
	for len(pq) != 0 {
		next := ListNode{
			Val: popMin(&pq),
		}
		ptr.Next = &next
		ptr = &next
	}

	return res.Next
}

func parent(node int) int {
	return (node - 1) / 2
}

func left(root int) int {
	return root*2 + 1
}

func right(root int) int {
	return root*2 + 2
}

func insert(list *[]int, node int) {
	temp := append((*list), node)

	swim(&temp, len(temp)-1)

	*list = temp
}

func swim(list *[]int, index int) {
	if index <= 0 {
		return
	}

	target := (*list)[index]
	parentIndex := parent(index)

	if target < (*list)[parentIndex] {
		swap(list, index, parentIndex)
		swim(list, parentIndex)
	}
}

func sink(list *[]int, index int) {
	size := len(*list)

	for index < size {
		idxLeft := left(index)
		idxRight := right(index)

		leftCanSwap := false
		rightCanSwap := false

		if idxLeft < size && ((*list)[index] >= (*list)[idxLeft]) {
			leftCanSwap = true
		}

		if idxRight < size && ((*list)[index] >= (*list)[idxRight]) {
			rightCanSwap = true
		}

		var idxToSwap int

		if !leftCanSwap && !rightCanSwap {
			break
		}

		if leftCanSwap && rightCanSwap {
			if (*list)[idxLeft] >= (*list)[idxRight] {
				idxToSwap = idxRight
			} else {
				idxToSwap = idxLeft
			}
		} else {
			if leftCanSwap {
				idxToSwap = idxLeft
			}

			if rightCanSwap {
				idxToSwap = idxRight
			}
		}

		swap(list, index, idxToSwap)

		index = idxToSwap
	}
}

func popMin(list *[]int) int {
	root := (*list)[0]

	swap(list, 0, len(*list)-1)
	*list = (*list)[:len(*list)-1]
	sink(list, 0)

	return root
}

func swap(list *[]int, index1 int, index2 int) {
	node1 := (*list)[index1]
	(*list)[index1] = (*list)[index2]
	(*list)[index2] = node1
}
