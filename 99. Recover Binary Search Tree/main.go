package main

import "math"

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

func recoverTree(root *TreeNode) {
	var prev *TreeNode = &TreeNode{Val: math.MinInt}
	var first *TreeNode = nil
	var second *TreeNode = nil
	var inorderRecover func(*TreeNode)
	inorderRecover = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorderRecover(node.Left)

		if node.Val < prev.Val {
			if first == nil {
				first = prev
			}
			second = node
		}
		prev = node
		inorderRecover(node.Right)
	}
	inorderRecover(root)
	first.Val, second.Val = second.Val, first.Val
}

/* func recoverTree(root *TreeNode) {
	prev = &TreeNode{
		Val: math.MinInt,
	}
	first = nil
	second = nil
	inorderRecover(root)

	first.Val, second.Val = second.Val, first.Val
} */

/* func inorderRecover(root *TreeNode) {
	if root == nil {
		return
	}

	inorderRecover(root.Left)

	if root.Val < prev.Val {
		if first == nil {
			first = prev
		}
		second = root
	}
	prev = root

	inorderRecover(root.Right)
} */
