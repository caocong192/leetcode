package removeNodes

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	dummy := &ListNode{}
	for head != nil {
		p := head
		head = head.Next
		p.Next = dummy.Next
		dummy.Next = p
	}
	return dummy.Next
}

func removeNodes(head *ListNode) *ListNode {
	head = reverse(head)
	for p := head; p.Next != nil; {
		if p.Val > p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return reverse(head)
}
