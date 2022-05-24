package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	ans := kthSmallest(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}, 1)
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

var rank int
var res int

func kthSmallest(root *TreeNode, k int) int {
	rank = 0
	res = 0
	traverse(root, k)
	return res
}

func traverse(root *TreeNode, k int) {
	if root == nil {
		return
	}

	traverse(root.Left, k)
	rank++
	if rank == k {
		res = root.Val
		return
	}
	traverse(root.Right, k)
}
