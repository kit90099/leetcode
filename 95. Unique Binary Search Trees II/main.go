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
func generateTrees(n int) []*TreeNode {
	memo = make([][][]*TreeNode, n)
	for i := 0; i < n; i++ {
		memo[i] = make([][]*TreeNode, n)
	}

	return getTrees(1, n)
}

var memo [][][]*TreeNode

func getTrees(lo int, hi int) []*TreeNode {
	if lo > hi {
		return []*TreeNode{nil}
	}

	if len(memo[lo-1][hi-1]) != 0 {
		return memo[lo-1][hi-1]
	}

	res := []*TreeNode{}
	for i := lo; i <= hi; i++ {

		left := getTrees(lo, i-1)
		right := getTrees(i+1, hi)

		for _, l := range left {
			for _, r := range right {
				res = append(res, &TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				})
			}
		}
	}

	memo[lo-1][hi-1] = res
	return res
}
