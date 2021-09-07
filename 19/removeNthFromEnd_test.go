package removeNthFromEnd

import (
	"reflect"
	"testing"
)

// TestRemoveNthFromEnd 单元测试
func TestRemoveNthFromEnd(t *testing.T) {
	type test struct { // 定义test结构体
		nums   []int
		n      int
		want []int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		"exampleTest2": {[]int{1}, 1, []int{}},
		"exampleTest3": {[]int{1, 2}, 1, []int{1}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := removeNthFromEnd(SliceToListNode(ts.nums), ts.n)
			if !reflect.DeepEqual(ListNodeToSlice(got), ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, ListNodeToSlice(got))
			}
		})
	}
}
