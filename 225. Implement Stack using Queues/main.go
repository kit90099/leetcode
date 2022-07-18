package main

type Node struct {
	val  *int
	next *Node
}

type Queue struct {
	head *Node
	end  *Node
	size int
}

func NewQueue() *Queue {
	return &Queue{
		size: 0,
	}
}

func (q *Queue) Push(val int) {
	node := Node{
		val: &val,
	}
	if q.size != 0 {
		q.end.next, q.end = &node, &node
	} else {
		q.head, q.end = &node, &node
	}
	q.size++
}

func (q *Queue) Pop() (x *int) {
	if q.size == 0 {
		x = nil
	} else if q.size == 1 {
		q.size = 0
		x = q.head.val
		q.head, q.end = nil, nil
	} else {
		q.size--
		x = q.head.val
		q.head = q.head.next
	}
	return
}

func (q *Queue) Peek() *int {
	if q.size == 0 {
		return nil
	} else {
		return q.head.val
	}
}

type MyStack struct {
	queue Queue
	end   int
}

func Constructor() MyStack {
	return MyStack{
		queue: *NewQueue(),
	}
}

func (this *MyStack) Push(x int) {
	this.queue.Push(x)
	this.end = x
}

func (this *MyStack) Pop() int {
	for this.queue.size != 0 && *(this.queue.Peek()) != this.end {
		n := this.queue.Pop()
		this.queue.Push(*n)
	}
	this.end = *this.queue.end.val
	return *this.queue.Pop()
}

func (this *MyStack) Top() int {
	return this.end
}

func (this *MyStack) Empty() bool {
	return this.queue.size == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
