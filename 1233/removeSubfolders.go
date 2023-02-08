package removeSubfolders

import "strings"

//func removeSubfolders(folder []string) (want []string) {
//	for _, str := range folder {
//		index := 0
//		for ; index < len(want); index++ {
//			// 如果相等, 则跳过
//			// 如果匹配前缀, 并且是子目录, (后面一位字符是 / ) 则跳过
//			if str == want[index] || strings.HasPrefix(str, want[index]) && string(str[len(want[index])]) == "/" {
//				goto end
//			} else if strings.HasPrefix(want[index], str) { // 反向查找, 如果str 是want中字符的前缀, 替换字符串
//				want = append(want[:index], want[index+1:]...)
//				index--
//				//want = append(want, str)
//				//break
//			}
//		}
//		// 如果上面循环没有匹配到, 则把str 加入到 want 中
//		want = append(want, str)
//
//	end:
//	}
//	return
//}

// 字典树
func removeSubfolders(folder []string) (want []string) {
	tree := createDictTree()

	// 插入字典树数据
	for i, str := range folder {
		tree.insert(i, str)
	}

	// 查找字典树中 val不为-1 的值
	for _, i := range tree.search() {
		want = append(want, folder[i])
	}
	return
}

// dictTree 字典树
type dictTree struct {
	children map[string]*dictTree
	val      int
}

func createDictTree() *dictTree {
	return &dictTree{
		map[string]*dictTree{},
		-1,
	}
}

func (t *dictTree) insert(val int, s string) {
	node := t

	strList := strings.Split(s, "/")
	for _, v := range strList[1:] {
		if _, ok := node.children[v]; !ok {
			node.children[v] = createDictTree()
		}
		node = node.children[v]
	}
	node.val = val
}

// search dfs查找字典树中第一个 val 不为-1的值
func (t *dictTree) search() (want []int) {
	var dfs func(*dictTree)
	dfs = func(root *dictTree) {
		if root.val != -1 {
			want = append(want, root.val)
			return
		}

		for _, children := range root.children {
			dfs(children)
		}
	}
	dfs(t)
	return
}
