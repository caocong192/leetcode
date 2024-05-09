package removeDuplicates

import (
	"reflect"
	"testing"
)

// TestRemoveDuplicates 单元测试
func TestMinimumRefill(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{1, 1, 1, 2, 2, 3}, 5},
		"exampleTest2": {[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := removeDuplicates(ts.nums)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
