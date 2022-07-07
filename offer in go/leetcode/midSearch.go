package leetcode

func search(data []int, target int) int {
	left, right := 0, len(data)-1
	for left <= right {
		mid := (right-left)/2 + left
		num := data[mid]
		if num == target {
			return mid
		} else if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
