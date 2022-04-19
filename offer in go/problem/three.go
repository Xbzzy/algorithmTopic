package problem

import "fmt"

func SearchInArrayThree(num, rows, columns int, data [][]int) bool {
	result := false
	if data != nil && rows > 0 && columns > 0 {
		row := 0
		column := columns - 1
		for row < rows && column >= 0 {
			if data[row][column] == num {
				result = true
				break
			} else if data[row][column] > num {
				column--
			} else {
				row++
			}
		}
	}
	return result
}

func TestThree() {
	data := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15}}
	fmt.Println(SearchInArrayThree(15, 4, 4, data))
}
