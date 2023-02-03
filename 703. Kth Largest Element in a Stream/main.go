package main

import "container/heap"

type KthLargest struct {
	pq *Pq
	k  int
}

func Constructor(k int, nums []int) KthLargest {
	pq := make(Pq, 0)
	for _, num := range nums {
		heap.Push(&pq, num)
		if len(pq) > k {
			heap.Pop(&pq)
		}
	}

	return KthLargest{
		pq: &pq,
		k:  k,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.pq, val)
	if len(*this.pq) > this.k {
		heap.Pop(this.pq)
	}

	return (*this.pq)[len(*this.pq)-1]
}

type Pq []int

func (pq *Pq) Len() int {
	return len(*pq)
}

func (pq *Pq) Less(i, j int) bool {
	return (*pq)[i] < (*pq)[j]
}

func (pq *Pq) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *Pq) Push(num interface{}) {
	*pq = append(*pq, num.(int))
}

func (pq *Pq) Pop() interface{} {
	res := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return res
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
