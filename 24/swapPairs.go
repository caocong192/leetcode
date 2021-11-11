package swapPairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// newHead 指向第二个节点
	newHead := head.Next
	// 递归后面的节点
	head.Next = swapPairs(newHead.Next)
	// 第二个节点指向第一个节点
	newHead.Next = head
	return newHead
}
