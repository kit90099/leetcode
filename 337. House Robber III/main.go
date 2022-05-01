package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	robResult := dp(root)
	return max(robResult[0], robResult[1])
}

func dp(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}

	left := dp(node.Left)
	right := dp(node.Right)

	return []int{
		node.Val + left[1] + right[1],
		max(left[0], left[1]) + max(right[0], right[1]),
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
