package leetcode

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{3, 1, 5, 7, 2, 4, 9, 6, 10, 8}
	fmt.Println(a)
	quickSort(a, 0, 9)
	fmt.Println(a)
}
