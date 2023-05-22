package sufficientSubset

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if calculateSum(root, 0, limit) {
		return root
	}
	return nil
}

func calculateSum(root *TreeNode, sum, limit int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val+sum >= limit
	}

	haveSufficientLeft := calculateSum(root.Left, sum+root.Val, limit)
	haveSufficientRight := calculateSum(root.Right, sum+root.Val, limit)
	if !haveSufficientLeft {
		root.Left = nil
	}

	if !haveSufficientRight {
		root.Right = nil
	}

	return haveSufficientLeft || haveSufficientRight
}
