package maxOperations

import (
	"reflect"
	"testing"
)

// TestMaxOperations 单元测试
func TestMaxOperations(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{3, 2, 1, 4, 5}, 2},
		"exampleTest2": {[]int{3, 2, 6, 1, 4}, 1},
		"exampleTest3": {[]int{2, 2, 3, 2, 4, 2, 3, 3, 1, 3}, 1},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := maxOperations(ts.nums)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
