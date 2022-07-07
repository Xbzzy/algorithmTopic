package leetcode

func partition(a []int, left, right int) int {
	key := a[left]
	for left < right {
		for left < right && a[right] >= key {
			right--
		}
		a[left], a[right] = a[right], a[left]
		for left < right && a[left] <= key {
			left++
		}
		a[left], a[right] = a[right], a[left]
	}
	return left
}

func quickSort(a []int, left, right int) {
	if left < right {
		tmp := partition(a, left, right)
		quickSort(a, left, tmp-1)
		quickSort(a, tmp+1, right)
	}
}
