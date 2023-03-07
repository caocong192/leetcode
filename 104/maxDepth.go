package maxDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) (ans int) {
	if root == nil {
		return 0
	}
	var dfs func(root *TreeNode, n int)
	dfs = func(root *TreeNode, n int) {
		if n > ans {
			ans = n
		}
		if root.Left != nil {
			dfs(root.Left, n+1)
		}
		if root.Right != nil {
			dfs(root.Right, n+1)
		}
	}
	dfs(root, 1)
	return ans
}
