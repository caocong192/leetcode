package minDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := 10000
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, i int) {
		if node.Left == nil && node.Right == nil {
			if ans > i {
				ans = i
			}
			return
		}

		if node.Left != nil {
			dfs(node.Left, i+1)
		}

		if node.Right != nil {
			dfs(node.Right, i+1)
		}
	}
	dfs(root, 1)
	return ans
}
