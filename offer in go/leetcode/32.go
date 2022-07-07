package leetcode

func jump(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	dp[0] = 1000
	dp[l-1] = 0
	for i := 0; i < l; i++ {
		for j := nums[i]; j > 0; j-- {
			if dp[i] > 1+dp[j] {
				dp[i] = 1 + dp[j]
			}
		}
	}
	return dp[0]
}
