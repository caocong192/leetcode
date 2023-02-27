package movesToMakeZigzag

import (
	"reflect"
	"testing"
)

// TestMovesToMakeZigzag 单元测试
func TestMovesToMakeZigzag(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{1, 2, 3}, 2},
		"exampleTest2": {[]int{9, 6, 1, 6, 2}, 4},
	}
	for name, ts := range tests {
		t.Run(name, func(s *testing.T) { // 使用t.Run()执行子测试
			got := movesToMakeZigzag(ts.nums)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
