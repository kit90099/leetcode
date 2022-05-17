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
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	return traverse(preorder, inorder, 0, n-1, 0, n-1)
}

func traverse(preorder []int, inorder []int, preStart int, preEnd int, inStart int, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	root := preorder[preStart]

	rootPos := -1
	for i, node := range inorder {
		if node == root {
			rootPos = i
			break
		}
	}

	leftSize := rootPos - inStart

	node := TreeNode{
		Val: root,
	}

	node.Left = traverse(preorder, inorder, preStart+1, preStart+leftSize, inStart, rootPos-1)
	node.Right = traverse(preorder, inorder, preStart+leftSize+1, preEnd, rootPos+1, inEnd)

	return &node
}
