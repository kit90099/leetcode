package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	ans := buildTree([]int{4, 2, 5, 1, 6, 3, 7}, []int{4, 5, 2, 6, 7, 3, 1})
	fmt.Println(ans)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder) - 1
	return traverse(inorder, postorder, 0, n, 0, n)
}

func traverse(inorder []int, postorder []int, instart int, inend int, poststart int, postend int) *TreeNode {
	if poststart > postend {
		return nil
	}

	root := postorder[postend]
	rootPos := -1
	for i, num := range inorder {
		if num == root {
			rootPos = i
			break
		}
	}
	leftSize := rootPos - instart

	node := TreeNode{
		Val:   root,
		Left:  traverse(inorder, postorder, instart, rootPos-1, poststart, poststart+leftSize-1),
		Right: traverse(inorder, postorder, rootPos+1, inend, poststart+leftSize, postend-1),
	}

	return &node
}
