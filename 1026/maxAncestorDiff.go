package maxAncestorDiff

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	var dfs func(root *TreeNode, m, n int) int
	dfs = func(root *TreeNode, m, n int) int {
		if root == nil {
			return 0
		}
		diff := max(abs(root.Val-m), abs(root.Val-n))
		m, n = max(root.Val, m), min(root.Val, n)
		diff = max(diff, dfs(root.Left, m, n))
		diff = max(diff, dfs(root.Right, m, n))
		return diff
	}

	return dfs(root, root.Val, root.Val)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
