package getIntersectionNode

type ListNode struct {
	Val  int
	Next *ListNode
}

//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	if headA == nil || headB == nil {
//		return nil
//	}
//
//	for
//}

//
//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	if headA == nil || headB == nil {
//		return nil
//	}
//
//	var pA, pB = headA, headB
//	for pA != pB {
//		if pA == nil {
//			pA = headB
//		} else {
//			pA = pA.Next
//		}
//
//		if pB == nil {
//			pB = headA
//		} else {
//			pB = pB.Next
//		}
//	}
//	return pA
//}
