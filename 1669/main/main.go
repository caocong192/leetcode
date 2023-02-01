package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	head := list1
	var aList, bList *ListNode

	for list1.Next != nil {

		if list1.Val == a {
			break
		}
		aList = list1
		list1 = list1.Next
	}

	for list1.Next != nil {
		if list1.Val == b {
			break
		}
		bList = list1
		list1 = list1.Next
	}

	aList.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = bList.Next

	return head
}

func generateListNode(l []int) *ListNode {
	var head *ListNode
	var ln = &ListNode{
		Val: 0,
		Next: nil,
	}
	head = ln.Next
	for _, v := range l {
		var newListNode = &ListNode{
			Val:  v,
			Next: nil,
		}
		ln.Next = newListNode
		ln = ln.Next
	}
	return head.Next
}

func printListNode(ln *ListNode) {
	for ln.Next != nil {
		fmt.Printf("%d ", ln.Val)
	}
}

func main() {
	var list1 = []int{0, 1, 2, 3, 4, 5}
	var list2 = []int{1000000, 1000001, 1000002}
	res := mergeInBetween(generateListNode(list1), 3, 4, generateListNode(list2))
	printListNode(res)
}
