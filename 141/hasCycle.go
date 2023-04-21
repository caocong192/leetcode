package hasCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	} else if head.Next == nil {
		return false
	}

	var head2 = head.Next

	for head != nil && head2 != nil {
		if head == head2 {
			return true
		}
		head = head.Next

		if head2.Next != nil {
			head2 = head2.Next.Next
		} else {
			return false
		}

	}
	return false
}

// 哈希表超出时间限制
//var hash = map[int]int{}

//func hasCycle(head *ListNode) bool {
//	for head != nil {
//		hash[head.Val]++
//
//		if hash[head.Val] == 2 {
//			return true
//		}
//		head = head.Next
//	}
//	return false
//}
