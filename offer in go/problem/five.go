package problem

import "fmt"

type listNode struct {
	value int
	next  *listNode
}

func printFive(head *listNode) {
	if head != nil {
		if head.next != nil {
			printFive(head.next)
		}
		fmt.Println(head.value)
	}
}

func TestFive() {
	list := new(listNode)
	head := list
	for i := 1; i < 5; i++ {
		head.next = new(listNode)
		head = head.next
		head.value = i
	}
	printFive(list)
}
