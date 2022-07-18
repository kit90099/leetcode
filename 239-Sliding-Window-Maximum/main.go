package main

import "fmt"

func main() {
	ans := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	fmt.Println(ans)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type MonotonicQueue struct {
	head *Node
	tail *Node
	size int
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{
		head: &Node{},
		tail: &Node{},
	}
}

func (mq *MonotonicQueue) Push(v int) {
	if mq.size == 0 {
		node := &Node{
			Val: v,
		}
		mq.tail.Left, mq.head.Right = node, node
		mq.size = 1
	} else {
		for mq.size > 0 && mq.tail.Left.Val < v {
			// pop last
			if mq.size == 1 {
				mq.head.Right = mq.tail
				mq.tail.Left = mq.head
				mq.size = 0
				break
			}
			mq.tail.Left, mq.tail.Left.Left.Right = mq.tail.Left.Left, mq.tail
			mq.size--
		}
		//add to tail
		node := &Node{
			Val:   v,
			Left:  mq.tail.Left,
			Right: mq.tail,
		}
		mq.tail.Left, mq.tail.Left.Right = node, node
		mq.size++
	}
}

func (mq *MonotonicQueue) Pop(v int) {
	if mq.head.Right.Val == v {
		mq.head.Right, mq.head.Right.Left = mq.head.Right.Right, nil
		mq.size--
	}
}

func (mq *MonotonicQueue) Max() int {
	return mq.head.Right.Val
}

func maxSlidingWindow(nums []int, k int) []int {
	mq := NewMonotonicQueue()

	// init k-1 element into queue
	for i := 0; i < k-1; i++ {
		mq.Push(nums[i])
	}
	res := make([]int, 0)
	for i := k - 1; i < len(nums); i++ {
		mq.Push(nums[i])
		res = append(res, mq.Max())
		mq.Pop(nums[i-k+1])
	}

	return res
}
