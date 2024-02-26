package rangeSumBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func rangeSumBST(root *TreeNode, low int, high int) int {
//
//	var dfs func(root *TreeNode, val *int) int
//	dfs = func(root *TreeNode, val *int) int {
//
//		if root == nil {
//			return 0
//		}
//
//		if root.Val >= low && root.Val <= high {
//			*val += root.Val
//		}
//
//		if root.Left != nil {
//			dfs(root.Left, val)
//		}
//
//		if root.Right != nil {
//			dfs(root.Right, val)
//		}
//
//		return *val
//	}
//
//	ans := 0
//	return dfs(root, &ans)
//
//}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}

	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}

	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}
