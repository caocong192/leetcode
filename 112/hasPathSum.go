package hasPathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	var dfs func(root *TreeNode, val int) bool
	dfs = func(root *TreeNode, val int) bool {

		if val+root.Val == targetSum && root.Left == nil && root.Right == nil {
			return true
		}
		if root.Left != nil {
			if dfs(root.Left, val+root.Val) {
				return true
			}
		}
		if root.Right != nil {
			if dfs(root.Right, val+root.Val) {
				return true
			}
		}

		return false
	}
	return dfs(root, 0)
}
