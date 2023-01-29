package minimumSize

import (
	"reflect"
	"testing"
)

// TestMinimumSize 单元测试
func TestMinimumSize(t *testing.T) {
	type test struct { // 定义test结构体
		nums          []int
		maxOperations int
		want          int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{9}, 2, 3},
		"exampleTest2": {[]int{2, 4, 8, 2}, 4, 2},
		"exampleTest3": {[]int{7, 20}, 2, 7},
		"exampleTest4": {[]int{7, 17}, 2, 7},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := minimumSize(ts.nums, ts.maxOperations)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
