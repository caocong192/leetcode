package maximumCount

import (
	"reflect"
	"testing"
)

// TestMaximumCount 单元测试
func TestMaximumCount(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{-2, -1, -1, 1, 2, 3}, 3},
		"exampleTest2": {[]int{-3, -2, -1, 0, 0, 1, 2}, 3},
		"exampleTest3": {[]int{5, 20, 66, 1314}, 4},
		"exampleTest4": {[]int{-1764, -1562, -1226, -1216, -402, -386, -133, 979, 1227, 1992}, 7},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := maximumCount(ts.nums)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
