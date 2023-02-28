package inorderTraversal

Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}

	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r.Left != nil {
			dfs(r.Left)
		}
		ans = append(ans, r.Val)
		if r.Right != nil {
			dfs(r.Right)
		}
	}
	dfs(root)
	return
}
