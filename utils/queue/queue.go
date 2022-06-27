package queue

type Queue[E any] interface {
	Push(e E)
	Pop() *E
	Peek() *E
	PeekEnd() *E
	Len() int
}
