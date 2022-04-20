package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	return max(&l, &r) + 1
}

func max(left *int, right *int) int {
	if *left > *right {
		return *left
	} else {
		return *right
	}
}
