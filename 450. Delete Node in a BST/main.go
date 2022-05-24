package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	/* ans := deleteNode(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}, 3) */

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
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if (root.Left == nil) || (root.Right == nil) {
			if root.Left == nil {
				return root.Right
			} else {
				return root.Left
			}
		} else {
			right := getRight(root.Left)

			deleteNode(root, right.Val)
			right.Left = root.Left
			right.Right = root.Right

			return right
		}
	} else {
		if key < root.Val {
			root.Left = deleteNode(root.Left, key)
		} else {
			root.Right = deleteNode(root.Right, key)
		}
	}

	return root
}

func getRight(root *TreeNode) *TreeNode {
	for root.Right != nil {
		root = root.Right
	}
	return root
}
