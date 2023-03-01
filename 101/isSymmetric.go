package isSymmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSymmetric 迭代
func isSymmetric(root *TreeNode) bool {
	// 创建p 存入两个树
	var p []*TreeNode
	p = append(p, root)
	p = append(p, root)

	// 一直循环p
	for len(p) > 0 {
		a, b := p[0], p[1]
		p = p[2:]

		// 如果都是nil, 跳过
		if a == nil && b == nil {
			continue
		}
		// 如果其中一个为nil
		if a == nil || b == nil {
			return false
		}
		// 如果值不相同
		if a.Val != b.Val {
			return false
		}
		// 追加, 对称的左右的值
		p = append(p, a.Left)
		p = append(p, b.Right)

		p = append(p, a.Right)
		p = append(p, b.Left)
	}
	return true
}

// 递归
//func isSymmetric(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//
//	leftTree, rightTree := root.Left, root.Right
//	return isAxialSymmetryTree(leftTree, rightTree)
//}
//
//func isAxialSymmetryTree(p *TreeNode, q *TreeNode) bool {
//	if p == nil && q == nil {
//		return true
//	} else if p == nil || q == nil {
//		return false
//	} else if p.Val != q.Val {
//		return false
//	} else {
//		return isAxialSymmetryTree(p.Left, q.Right) && isAxialSymmetryTree(p.Right, q.Left)
//	}
//}
