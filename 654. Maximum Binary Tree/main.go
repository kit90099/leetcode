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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return traverse(nums)
}

func traverse(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}

	max, maxPos := -1, -1
	for i, num := range nums {
		if num > max {
			max = num
			maxPos = i
		}
	}

	node := TreeNode{
		Val: max,
	}

	node.Left = traverse(nums[:maxPos])
	node.Right = traverse(nums[maxPos+1:])

	return &node
}
