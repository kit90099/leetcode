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

func minDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}

	slice := []*TreeNode{root}
	depth := 1

	for len(slice) != 0 {
		temp := make([]*TreeNode, 0)
		for _, node := range slice {
			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				temp = append(temp, node.Left)
			}

			if node.Right != nil {
				temp = append(temp, node.Right)
			}

		}
		depth++
		slice = temp
	}

	return depth
}
