package solution

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	curNode := &ListNode{}
	head := curNode
	flag1 := false
	flag2 := false

	for {
		if flag1 && flag2 {
			if carry == 1 {
				newNode := &ListNode{
					1,
					nil,
				}
				curNode.Next = newNode
			}
			return head.Next
		}

		newNode := &ListNode{}
		if l1.Val+l2.Val+carry > 9 {
			newNode.Val = l1.Val + l2.Val + carry - 10
			carry = 1
		} else {
			newNode.Val = l1.Val + l2.Val + carry
			carry = 0
		}
		curNode.Next = newNode
		curNode = newNode

		if l1.Next == nil {
			flag1 = true
		}

		if l2.Next == nil {
			flag2 = true
		}

		if !flag1 {
			l1 = l1.Next
		} else {
			l1.Val = 0
		}

		if !flag2 {
			l2 = l2.Next
		} else {
			l2.Val = 0
		}
	}
}
