package solution

import (
	"reflect"
	"testing"
)

// 输出函数
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

// TestAddTwoNumbers 单元测试
func TestAddTwoNumbers(t *testing.T) {
	type test struct { // 定义test结构体
		l1   []int
		l2   []int
		want []int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		"exampleTest2": {[]int{0}, []int{0}, []int{0}},
		"exampleTest3": {[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			l1 := makeListNode(ts.l1)
			l2 := makeListNode(ts.l2)
			want := makeListNode(ts.want)
			got := addTwoNumbers(l1, l2)
			retWant := printListNode(want)
			retGot := printListNode(got)

			if !reflect.DeepEqual(retGot, retWant) {
				t.Errorf("excepted:%#v, got:%#v", want, got)
			}
		})
	}
}
