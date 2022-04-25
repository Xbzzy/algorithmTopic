package problem

import (
	"fmt"
)

func binaryTen(data int) int {
	count := 0
	for data != 0 {
		count++
		data = (data - 1) & data
	}
	return count
}

func TestTen() {
	for i := 0; i < 30; i += 5 {
		fmt.Println(binaryTen(i))
	}
}
