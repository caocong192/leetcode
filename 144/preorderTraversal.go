package preorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func preorderTraversal(root *TreeNode) (ans []int) {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			ans = append(ans, node.Val)
		}
		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
	}
	if root == nil {
		return
	}
	dfs(root)
	return
}
