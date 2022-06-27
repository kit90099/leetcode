package main

import (
	"fmt"

	"com.grpk.utils/queue"
)

func main() {
	var t queue.Queue[int] = queue.NewMonotonicQueue[int]()
	t.Push(1)
	t.Push(2)
	t.Push(3)
	t.Push(4)
	print(*t.Pop())
	print(*t.Pop())
	print(*t.Pop())
	print(*t.Pop())
}

func print(i int) {
	fmt.Println(i)
}
