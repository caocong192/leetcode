package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	curNode := &ListNode{}
	head := curNode
	flag1 := false // 标识 l1 是否已经遍历完了
	flag2 := false // 标识 l2 是否已经遍历完了

	for {

		// 如果都遍历完了, 则返回结果
		if flag1 && flag2 {
			// 如果进位是1, 则先进位
			if carry == 1 {
				newNode := &ListNode{
					1,
					nil,
				}
				curNode.Next = newNode
			}
			return head.Next
		}

		// 相加逻辑
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

		// 判断l1是否还有下一个链
		if l1.Next == nil {
			// 如果没有, 则设置标识位flag1 为true
			flag1 = true
		}
		// 判断l2是否还有下一个链
		if l2.Next == nil {
			// 如果没有, 则设置标识位flag2 为true
			flag2 = true
		}

		// 如果l1还有链
		if !flag1 {
			// 往下一个链走
			l1 = l1.Next
		} else {
			// 如果没有链了, 则置当前值为0
			l1.Val = 0
		}

		// 如果l2还有链
		if !flag2 {
			// 往下一个链走
			l2 = l2.Next
		} else {
			// 如果没有链了, 则置当前值为0
			l2.Val = 0
		}
	}
}

// 构造函数, 构造 ListNode 对象
func makeListNode(nodeArr []int) *ListNode {
	var head *ListNode
	curNode := &ListNode{}
	head = curNode

	for i := 0; i < len(nodeArr); i++ {
		newNode := &ListNode{
			nodeArr[i],
			nil,
		}
		curNode.Next = newNode
		curNode = newNode
	}
	return head.Next
}

func printListNode(node *ListNode) []int {
	var ret []int
	for {
		ret = append(ret, node.Val)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
	return ret
}

func main() {
	arr1 := []int{9, 9, 9, 9, 9, 9, 9}
	arr2 := []int{9, 9, 9, 9}

	l1 := makeListNode(arr1)
	l2 := makeListNode(arr2)
	ret1 := printListNode(l1)
	fmt.Println(ret1)
	ret2 := printListNode(l2)
	fmt.Println(ret2)

	want := addTwoNumbers(l1, l2)

	ret := printListNode(want)
	fmt.Println(ret)
}
