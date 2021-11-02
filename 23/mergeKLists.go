package mergeKLists

type ListNode struct {
	Val  int
	Next *ListNode
}

// 合并K个ListNode, 调用merge方法
func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

// 分治合并: 两两合并ListNode
func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) >> 1
	return mergeTwoListNode(merge(lists, l, mid), merge(lists, mid+1, r))
}

// 合并两个 ListNode
func mergeTwoListNode(node1 *ListNode, node2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	for node1 != nil && node2 != nil {
		if node1.Val > node2.Val {
			tail.Next = node2
			node2 = node2.Next
		} else {
			tail.Next = node1
			node1 = node1.Next
		}
		tail = tail.Next
	}

	if node1 == nil {
		tail.Next = node2
	} else {
		tail.Next = node1
	}
	return head.Next
}
