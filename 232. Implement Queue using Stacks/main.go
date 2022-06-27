package main

type MyQueue struct {
	left  *MyStack
	right *MyStack
}

type MyStack struct {
	top  *StackNode
	size int
}

type StackNode struct {
	val  int
	next *StackNode
}

func newStatck() *MyStack {
	return &MyStack{
		size: 0,
	}
}

func (s *MyStack) push(i int) {
	node := &StackNode{
		val:  i,
		next: s.top,
	}
	s.top = node
	s.size++
}

func (s *MyStack) pop() *int {
	if s.top == nil {
		return nil
	} else {
		res := &s.top.val
		s.top = s.top.next
		s.size--
		return res
	}
}

func (s *MyStack) peek() *int {
	if s.top == nil {
		return nil
	} else {
		return &s.top.val
	}
}

func Constructor() MyQueue {
	return MyQueue{
		left:  newStatck(),
		right: newStatck(),
	}
}

func (this *MyQueue) Push(x int) {
	if this.left.size == 0 {
		this.right.push(x)
	} else {
		for this.left.size != 0 {
			i := this.left.pop()
			this.right.push(*i)
		}
		this.left.push(x)
	}
}

func (this *MyQueue) Pop() int {
	for this.right.size != 0 {
		i := this.right.pop()
		this.left.push(*i)
	}
	return *this.left.pop()
}

func (this *MyQueue) Peek() int {
	for this.right.size != 0 {
		i := this.right.pop()
		this.left.push(*i)
	}
	return *this.left.peek()
}

func (this *MyQueue) Empty() bool {
	return this.left.size+this.right.size == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
