package leetcode

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func maxVal(a, b int) int {
	if math.Abs(float64(a)) > math.Abs(float64(b)) {
		return a
	} else {
		return b
	}
}

func dfs1(node *TreeNode, up, low int, max *int) {
	if node == nil {
		return
	}
	*max = maxVal(maxVal(node.Val-up, node.Val-low), *max)
	up = maxVal(node.Val, up)
	low = int(math.Min(float64(node.Val), float64(low)))
	dfs1(node.Left, up, low, max)
	dfs1(node.Right, up, low, max)
}

func maxAncestorDiff(root *TreeNode) int {
	var max int
	dfs1(root, root.Val, root.Val, &max)
	return int(math.Abs(float64(max)))
}
