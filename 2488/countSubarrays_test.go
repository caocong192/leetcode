package countSubarrays

import (
	"reflect"
	"testing"
)

// TestCountSubarrays 单元测试
func TestCountSubarrays(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		k    int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{3, 2, 1, 4, 5}, 4, 3},
		"exampleTest2": {[]int{2, 3, 1}, 3, 1},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := countSubarrays(ts.nums, ts.k)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
