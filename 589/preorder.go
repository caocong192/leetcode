package preorder

type Node struct {
	Val      int
	Children []*Node
}

//func preorder(root *Node) []int {
//	var ans []int
//	var dfs func(*Node)
//	dfs = func(node *Node) {
//		if node == nil {
//			return
//		}
//		ans = append(ans, node.Val)
//		for _, ch := range node.Children {
//			dfs(ch)
//		}
//	}
//	dfs(root)
//	return ans
//}

func preorder(root *Node) (ans []int) {
	// 2. 迭代法
	// 模拟栈 stack FILO特性
	stack := []*Node{root}
	for len(stack) > 0 {
		// 出栈
		node := stack[len(stack)-1]
		ans = append(ans, node.Val)
		stack = stack[:len(stack)-1]

		// 从右往左入栈, 则每次出栈为最左边的树
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return
}
