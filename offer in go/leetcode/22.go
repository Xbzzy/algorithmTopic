package leetcode

func generateParenthesis(n int) []string {
	res := []string{}
	if n == 0 {
		return res
	}
	dfs("", n, n, &res)
	return res
}

func dfs(curStr string, leftNum int, rightNum int, res *[]string) {
	if leftNum == 0 && rightNum == 0 {
		*res = append(*res, curStr)
		return
	}
	if leftNum > rightNum {
		return
	}
	if leftNum > 0 {
		dfs(curStr+"(", leftNum-1, rightNum, res)
	}
	if rightNum > 0 {
		dfs(curStr+")", leftNum, rightNum-1, res)
	}
}
