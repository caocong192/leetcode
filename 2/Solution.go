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
	var carry int
	var node ListNode
	for {
		var new_node ListNode
		if l1.Next != nil || l2.Next != nil || carry != 0 {
			if l1.Val+l2.Val >= 10 {
				new_node.Val = (l1.Val + l2.Val) - 10
				carry = 1
			}
			l1.Next

		}
	}

}
