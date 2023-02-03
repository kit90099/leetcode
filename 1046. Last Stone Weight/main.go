package main

import "fmt"

func main() {
	fmt.Println(lastStoneWeight([]int{7, 6, 7, 6, 9}))
}

func lastStoneWeight(stones []int) int {

	pq := HeapConstructor()
	for _, s := range stones {
		pq.insert(s)
	}

	for pq.size > 1 {
		a := pq.popMax()
		b := pq.popMax()
		pq.insert(a - b)
	}

	return pq.popMax()
}

type Heap struct {
	data []int
	size int
}

func HeapConstructor() *Heap {
	h := Heap{
		data: []int{-1},
	}
	return &h
}

func (this *Heap) swap(a, b int) {
	this.data[a], this.data[b] = this.data[b], this.data[a]
}

func (this *Heap) swim(idx int) {
	for idx > 1 && this.data[parent(idx)] < this.data[idx] {
		this.swap(idx, parent(idx))
		idx = parent(idx)
	}
}

func (this *Heap) sink(idx int) {
	for left(idx) <= this.size {
		max := left(idx)

		if right(idx) <= this.size && this.data[max] < this.data[right(idx)] {
			max = right(idx)
		}

		if this.data[idx] > this.data[max] {
			break
		}

		this.swap(max, idx)
		idx = max
	}
}

func (this *Heap) insert(num int) {
	this.data = append(this.data, num)
	this.swim(len(this.data) - 1)
	this.size++
}

func (this *Heap) popMax() int {
	if this.size == 0 {
		return -1
	}
	res := this.data[1]
	this.swap(1, this.size)
	this.data = this.data[:len(this.data)-1]
	this.size--
	this.sink(1)
	return res
}

func parent(idx int) int {
	return idx / 2
}

func left(idx int) int {
	return idx * 2
}

func right(idx int) int {
	return idx*2 + 1
}
