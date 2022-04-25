package problem

import "fmt"

func FiboNine(num uint) int {
	var i uint
	result := []int{0, 1}
	if num < 2 {
		return result[num]
	}
	fiboOne := 1
	fiboTwo := 0
	fiboNum := 0
	for i = 2; i <= num; i++ {
		fiboNum = fiboOne + fiboTwo
		fiboTwo = fiboOne
		fiboOne = fiboNum
	}
	return fiboNum
}

func TestNine() {
	var i uint
	for i = 0; i < 10; i++ {
		fmt.Println(FiboNine(i))
	}
}
