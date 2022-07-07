package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	cur := head
	for cur.Next != nil {
		next := cur.Next
		secNext := next.Next
		next.Next = cur
		cur.Next = secNext
		head = next
	}
	return head
}
