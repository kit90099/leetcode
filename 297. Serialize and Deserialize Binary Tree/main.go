package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
	}

	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(root)
	ans := deser.deserialize(data)

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

type Codec struct {
	splitStr string
	nilStr   string
}

func Constructor() Codec {
	return Codec{
		splitStr: ",",
		nilStr:   "#",
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return this.nilStr
	}

	res := strconv.Itoa(root.Val)
	left := this.serialize(root.Left)
	right := this.serialize(root.Right)

	res = res + this.splitStr + left + this.splitStr + right
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	splitted := strings.Split(data, this.splitStr)

	node, _ := de(splitted, 0, this.nilStr)
	return node
}

func de(arr []string, rangeStart int, nilStr string) (*TreeNode, int) {
	if rangeStart > len(arr) {
		return nil, rangeStart
	}

	if arr[rangeStart] == nilStr {
		return nil, rangeStart
	}

	root, _ := strconv.Atoi(arr[rangeStart])

	node := &TreeNode{
		Val: root,
	}

	var end int
	node.Left, end = de(arr, rangeStart+1, nilStr)
	node.Right, end = de(arr, end+1, nilStr)

	return node, end
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
