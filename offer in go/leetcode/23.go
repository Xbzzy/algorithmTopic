package leetcode

import "container/heap"

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

type ListNodes []*ListNode1

func (l ListNodes) Len() int {
	return len(l)
}

func (l ListNodes) Less(i, j int) bool {
	return l[i].Val <= l[j].Val
}

func (l ListNodes) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l *ListNodes) Pop() interface{} {
	n := len(*l)
	x := (*l)[n-1]
	*l = (*l)[:n-1]
	return x
}

func (l *ListNodes) Push(x interface{}) {
	*l = append(*l, x.(*ListNode1))
}

var _ heap.Interface = (*ListNodes)(nil)

func mergeKLists(lists []*ListNode1) *ListNode1 {
	head := new(ListNode1)
	pre := head
	hp := &ListNodes{}
	for _, val := range lists {
		if val != nil {
			heap.Push(hp, val)
		}
	}
	for len(*hp) > 0 {
		top := heap.Pop(hp).(*ListNode1)
		pre.Next = top
		pre = pre.Next
		if top.Next != nil {
			heap.Push(hp, top.Next)
		}
	}
	return head.Next
}
