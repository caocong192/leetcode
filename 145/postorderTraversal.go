package postorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 后续遍历
func postorderTraversal(root *TreeNode) (ans []int) {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {

		if node.Right != nil {
			dfs(node.Right)
		}

		if node.Left != nil {
			dfs(node.Left)
		}

		if node != nil {
			ans = append(ans, node.Val)
		}
	}
	if root == nil {
		return
	}
	dfs(root)
	return
}
