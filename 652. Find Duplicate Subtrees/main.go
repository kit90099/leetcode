package main

import (
	"fmt"
)

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

var arrMap map[string]int
var arr []*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	arrMap = make(map[string]int)
	arr = make([]*TreeNode, 0)

	traverse(root)
	return arr
}

func traverse(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	left := traverse(root.Left)
	right := traverse(root.Right)

	subTree := fmt.Sprintf("%s,%s,%d", left, right, root.Val)
	freq, exist := arrMap[subTree]
	if exist && freq == 1 {
		arr = append(arr, root)
	}

	arrMap[subTree] = arrMap[subTree] + 1
	return subTree
}
