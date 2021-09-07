package removeNthFromEnd

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToListNode(nums []int) *ListNode {
	var head ListNode
	var index = &head
	for i := 0; i < len(nums); i++ {
		var node = ListNode{
			nums[i],
			nil,
		}
		index.Next = &node
		index = index.Next
	}
	return head.Next
}

func ListNodeToSlice(node *ListNode) []int {
	var s = []int{}
	for ; node != nil; node = node.Next {
		s = append(s, node.Val)
	}
	return s
}

func GetListNodeLen(node *ListNode) (n int) {
	for ; node != nil; node = node.Next {
		n += 1
	}
	return
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := GetListNodeLen(head)

	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}
