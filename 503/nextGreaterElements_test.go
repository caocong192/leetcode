package nextGreaterElements

import (
	"reflect"
	"testing"
)

// TestNextGreaterElements 单元测试
func TestNextGreaterElements(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		want []int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{1, 2, 1}, []int{2, -1, 2}},
		"exampleTest2": {[]int{1, 2, 3, 4, 3}, []int{2, 3, 4, -1, 4}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := nextGreaterElements(ts.nums)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
