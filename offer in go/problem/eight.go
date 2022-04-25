package problem

import "fmt"

func MinEight(data []int) int {
	length := len(data)
	head := data[0]
	tail := data[length-1]
	mid := head
	for data[head] >= data[tail] {
		if tail-head == 1 {
			mid = tail
			break
		}
		mid = (head + tail) / 2
		if data[head] == data[tail] && data[mid] == data[head] {
			result := data[head]
			for i := head + 1; i < tail; i++ {
				if result > data[i] {
					result = data[i]
				}
			}
			return result
		}
		if data[mid] >= data[head] {
			head = mid
		} else if data[mid] <= data[tail] {
			tail = mid
		}
	}
	return data[mid]
}

func TestEight() {
	data := []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}
	fmt.Println(MinEight(data))
}
