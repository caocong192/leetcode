package alternatingSubarray

import (
	"reflect"
	"testing"
)

// TestAlternatingSubarray 单元测试
func TestAlternatingSubarray(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{2, 3, 4, 3, 4}, 4},
		"exampleTest2": {[]int{3, 4, 5}, 2},
		"exampleTest3": {[]int{1, 3, 5, 4}, -1},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := alternatingSubarray(ts.nums)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
