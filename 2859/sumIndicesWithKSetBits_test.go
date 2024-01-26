package sumIndicesWithKSetBits

import (
	"reflect"
	"testing"
)

// TestSumIndicesWithKSetBits 单元测试
func TestSumIndicesWithKSetBits(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		k    int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{5, 10, 1, 5, 2}, 1, 13},
		"exampleTest2": {[]int{4, 3, 2, 1}, 2, 1},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := sumIndicesWithKSetBits(ts.nums, ts.k)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
