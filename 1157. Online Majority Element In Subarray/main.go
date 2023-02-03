package main

import (
	"fmt"
	"math"
)

/* type MajorityChecker struct {
}

func Constructor(arr []int) MajorityChecker {

}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {

} */

type SegmentTree struct {
	root *TreeNode
	size int
}

func (this *SegmentTree) query(left int, right int) int {
	return this._query(1, this.size, this.root, left, right)
}

func (this *SegmentTree) _query(start int, end int, node *TreeNode, left int, right int) int {
	if node == nil {
		return 0
	}
	// fully cover
	if left <= start && end <= right {
		return node.maxAppearance
	}

	mid := left + (right-left)/2
	maxAppearance := math.MinInt
	if left <= mid {
		maxAppearance = max(maxAppearance, this._query(start, mid, node.left, left, right))
	}
	if mid < right {
		maxAppearance = max(maxAppearance, this._query(mid+1, right, node.right, left, right))
	}
	return maxAppearance
}

func (this *SegmentTree) update(val int) {
	this._update(this.root, 1, this.size, val)
}

func (this *SegmentTree) _update(node *TreeNode, start int, end int, val int) {
	if start == end && start == val {
		node.appearance += 1
		node.maxAppearance = node.appearance
		return
	}

	createNode(node)
	mid := start + (end-start)/2
	if val <= mid {
		this._update(node.left, start, mid, val)
	} else {
		this._update(node.right, mid+1, end, val)
	}
	this.pushUp(node)
}

func (this *SegmentTree) pushUp(node *TreeNode) {
	maxAppearance := math.MinInt
	if node.left != nil {
		maxAppearance = max(maxAppearance, node.left.maxAppearance)
	}
	if node.right != nil {
		maxAppearance = max(maxAppearance, node.right.maxAppearance)
	}
	node.maxAppearance = maxAppearance
}

func createNode(root *TreeNode) {
	if root.left == nil {
		root.left = &TreeNode{}
	}
	if root.right == nil {
		root.right = &TreeNode{}
	}
}

type TreeNode struct {
	left          *TreeNode
	right         *TreeNode
	maxAppearance int
	appearance    int
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/**
 * Your MajorityChecker object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,threshold);
 */

func main() {
	s := SegmentTree{
		size: 10,
		root: &TreeNode{},
	}
	s.update(1)
	s.update(1)
	s.update(2)
	s.update(3)
	s.update(4)
	s.update(5)
	s.update(6)
	s.update(7)
	s.update(8)
	s.update(9)
	s.update(10)
	ans := s.query(2, 10)

	fmt.Println(ans)
}
