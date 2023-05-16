package isBalanced

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return getHeight(root) >= 0
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 自下而上 判断每个子叶和节点是否是平衡二叉树
	var leftHeight, rightHeight int
	if leftHeight = getHeight(root.Left); leftHeight == -1 {
		return -1
	}
	if rightHeight = getHeight(root.Right); rightHeight == -1 {
		return -1
	}
	// 差值大于1 则返回-1
	if abs(leftHeight, rightHeight) > 1 {
		return -1
	}

	return max(leftHeight, rightHeight) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a-b < 0 {
		return (a - b) * -1
	}
	return a - b
}
