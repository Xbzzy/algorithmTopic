package leetcode

func maxArea(height []int) int {
	var result int
	dp := make([][]int, len(height))
	for i := range dp {
		dp[i] = make([]int, len(height))
	}
	for i := 1; i < len(height); i++ {
		for j := 0; j < i; j++ {
			dp[j][i] = max(dp[j][i-1], min(height[j], height[i])*(i-j))
			result = max(result, dp[j][i])
		}
	}
	return result
}

func maxArea1(height []int) int {
	l, r, area, ans := 0, len(height)-1, 0, 0
	for l < r {
		if height[l] < height[r] {
			area = height[l] * (r - l)
		} else {
			area = height[r] * (r - l)
		}
		if ans < area {
			ans = area
		}
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}
