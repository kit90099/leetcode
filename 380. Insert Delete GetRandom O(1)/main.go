package main

import (
	"math"
	"math/rand"
)

type RandomizedSet struct {
	arr    []int
	posMap map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		arr:    make([]int, 0),
		posMap: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, exist := this.posMap[val]
	if exist {
		return false
	} else {
		this.arr = append(this.arr, val)
		this.posMap[val] = len(this.arr) - 1
		return true
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	pos, exist := this.posMap[val]
	if !exist {
		return false
	} else {
		// swap
		temp := this.arr[pos]
		valToBeSwapped := this.arr[len(this.arr)-1]
		this.posMap[valToBeSwapped] = pos
		this.arr[pos] = valToBeSwapped
		this.arr = this.arr[:len(this.arr)-1]
		delete(this.posMap, temp)
		return true
	}
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[int(math.Floor(rand.Float64()*float64(len(this.arr))))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
