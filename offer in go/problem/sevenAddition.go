package problem

import "fmt"

type queue1 struct {
	data      []int
	headIndex int
	tailIndex int
	cap       int
	eleNum    int
}

func (q *queue1) initQueue(cap int) {
	q.cap = cap
	q.data = make([]int, cap)
	q.tailIndex = -1
	q.headIndex = -1
}

func (q *queue1) append(value int) bool {
	if q.eleNum == 0 {
		q.headIndex++
		q.tailIndex++
		q.data[0] = value
		return true
	}
	if q.eleNum != q.cap {
		q.eleNum++
		q.tailIndex = (q.tailIndex + q.cap - 1) % q.cap
		q.data[q.tailIndex] = value
		return true
	} else {
		return false
	}
}

func (q *queue1) delete() {
	if q.eleNum != 0 {
		q.eleNum--
		q.headIndex = (q.headIndex + q.cap - 1) % q.cap
	}
}

func TestSevenAdd() {
	tmpQueue := new(queue1)
	tmpQueue.initQueue(10)
	tmpQueue.append(1)
	fmt.Println(tmpQueue.data[0])
}
