package canJump

import (
	"reflect"
	"testing"
)

// TestCanJump 单元测试
func TestCanJump(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		want bool
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{2, 3, 1, 1, 4}, true},
		"exampleTest2": {[]int{3, 2, 1, 0, 4}, false},
		"exampleTest3": {[]int{0}, true},
		"exampleTest4": {[]int{1, 1, 1, 0}, true},
		"exampleTest5": {[]int{5, 4, 0, 2, 0, 1, 0, 1, 0}, false},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := canJump(ts.nums)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
