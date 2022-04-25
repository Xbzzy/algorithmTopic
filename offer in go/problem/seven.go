package problem

import "fmt"

type stack struct {
	data   []int
	eleNum int
	cap    int
}

func (s *stack) initStack(cap int) {
	s.data = make([]int, cap)
	s.cap = cap
}

func (s *stack) push(value int) bool {
	if s.eleNum < s.cap {
		s.data[s.eleNum] = value
		s.eleNum++
		return true
	} else {
		return false
	}
}

func (s *stack) pop() int {
	s.eleNum--
	return s.data[s.eleNum]
}

type queue struct {
	stack1 stack
	stack2 stack
}

func (q *queue) initQueue(cap int) {
	q.stack1.initStack(cap)
	q.stack2.initStack(cap)
}

func (q *queue) append(value int) {
	q.stack1.push(value)
}

func (q *queue) delete() {
	if q.stack2.eleNum == 0 {
		for q.stack1.eleNum > 0 {
			q.stack2.push(q.stack1.pop())
		}
		q.stack2.pop()
	} else {
		q.stack2.pop()
	}
}

func TestSeven() {
	tmpQueue := new(queue)
	tmpQueue.initQueue(10)
	for i := 0; i < 10; i++ {
		tmpQueue.append(i)
	}
	fmt.Println(tmpQueue.stack1.data)
	for j := 0; j < 5; j++ {
		tmpQueue.delete()
	}
	fmt.Println(tmpQueue.stack2.eleNum)
	fmt.Println(tmpQueue.stack2.data)
}
