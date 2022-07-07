package leetcode

func canJump(nums []int) bool {
	l := len(nums)
	if l == 1 {
		return true
	}
	var jumpMost int
	for i := 0; i < l; i++ {
		if i <= jumpMost {
			if jumpMost < nums[i]+i {
				jumpMost = nums[i] + i
			}
			if jumpMost >= l-1 {
				return true
			}
		}
	}
	return false
}
