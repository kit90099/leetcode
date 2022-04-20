package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node := TreeNode{
		Left: &TreeNode{
			Val: 0,
		},
	}

	fmt.Println(diameterOfBinaryTree(&node))
}

var maxDiameter = 0

func diameterOfBinaryTree(root *TreeNode) int {
	maxDepth(root)
	return maxDiameter
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	temp := l + r

	maxDiameter = max(&maxDiameter, &temp)

	return max(&l, &r) + 1
}

func max(left *int, right *int) int {
	if *left > *right {
		return *left
	} else {
		return *right
	}
}
