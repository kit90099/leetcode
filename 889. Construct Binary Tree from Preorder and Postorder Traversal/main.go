package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	constructFromPrePost([]int{1, 2, 5, 3, 6, 7}, []int{5, 2, 6, 7, 3, 1})
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var postmap map[int]int

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	postmap = make(map[int]int)
	for i, num := range postorder {
		postmap[num] = i
	}
	n := len(preorder) - 1
	return traverse(preorder, postorder, 0, n, 0, n)
}

func traverse(preorder []int, postorder []int, prestart int, preend int, poststart int, postend int) *TreeNode {
	if prestart > preend {
		return nil
	}
	if prestart == preend {
		return &TreeNode{
			Val: preorder[prestart],
		}
	}

	root := preorder[prestart]
	left := preorder[prestart+1]
	leftPostPos := postmap[left]
	leftSize := leftPostPos - poststart + 1

	fmt.Println("test")

	node := &TreeNode{
		Val:   root,
		Left:  traverse(preorder, postorder, prestart+1, prestart+leftSize, poststart, poststart+leftSize-1),
		Right: traverse(preorder, postorder, prestart+leftSize+1, preend, leftPostPos+1, postend-1),
	}

	return node
}
