package queue

func Test() {

}

type PriorityQueue struct {
	arr *[]int
}

func New() *PriorityQueue {
	return &PriorityQueue{
		&([]int{-1}),
	}
}

func (queue *PriorityQueue) AddElement(element int) {
	*queue.arr = append(*(queue.arr), element)
	queue.shiftUp(len(*queue.arr) - 1)
}

func (queue *PriorityQueue) shiftUp(key int) {
	for key > 1 && (*queue.arr)[key] < (*queue.arr)[queue.parent(key)] {
		queue.switchPlace(key, queue.parent(key))
		key = queue.parent(key)
	}
}

func (queue *PriorityQueue) shiftDown(key int) {
	for queue.right(key) <= len(*queue.arr) {
		var smallerKey int
		left := queue.left(key)
		right := queue.right(key)

		lenarr := len(*queue.arr)

		if left >= lenarr {
			break
		}

		if right >= lenarr {
			right = left
		}

		if (*queue.arr)[left] > (*queue.arr)[right] {
			smallerKey = right
		} else {
			smallerKey = left
		}

		if (*queue.arr)[key] > (*queue.arr)[smallerKey] {
			queue.switchPlace(key, smallerKey)
			key = smallerKey
		} else {
			break
		}
	}

}

func (queue *PriorityQueue) parent(key int) int {
	return key / 2
}

func (queue *PriorityQueue) left(key int) int {
	return key * 2
}

func (queue *PriorityQueue) right(key int) int {
	return key*2 + 1
}

func (queue *PriorityQueue) switchPlace(k1 int, k2 int) {
	temp := (*queue.arr)[k1]
	(*queue.arr)[k1] = (*queue.arr)[k2]
	(*queue.arr)[k2] = temp
}

func (queue *PriorityQueue) PopMin() int {
	if len(*queue.arr) >= 2 {

		result := (*queue.arr)[1]

		if queue.left(1) > len(*queue.arr)-1 && queue.right(1) > len(*queue.arr)-1 {
			if len(*queue.arr) == 2 {
				*queue.arr = (*queue.arr)[0 : len(*queue.arr)-1]
				return result
			}
			return -1
		}

		queue.switchPlace(1, len(*queue.arr)-1)
		*queue.arr = (*queue.arr)[0 : len(*queue.arr)-1]

		queue.shiftDown(1)

		return result
	}
	return -1
}
