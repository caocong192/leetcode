package nextLargerNodes

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) (ans []int) {

	var value []int
	for head != nil {
		ans = append(ans, 0)
		value = append(value, head.Val)

		for i := 0; i < len(value); i++ {
			if ans[i] != 0 {
				continue
			}

			if head.Val > value[i] {
				ans[i] = head.Val
			}
		}

		head = head.Next
	}

	return
}
