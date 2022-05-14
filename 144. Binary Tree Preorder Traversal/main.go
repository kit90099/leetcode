package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	arr := []int{}
	traverse(root, &arr)
	return arr
}

func traverse(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}

	(*arr) = append(*arr, node.Val)
	traverse(node.Left, arr)
	traverse(node.Right, arr)
}
