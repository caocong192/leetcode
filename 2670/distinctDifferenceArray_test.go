package distinctDifferenceArray

import (
	"reflect"
	"testing"
)

// TestDistinctDifferenceArray
func TestDistinctDifferenceArray(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		want []int
	}
	const s = "$problem$"
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{1, 2, 3, 4, 5}, []int{-3, -1, 1, 3, 5}},
		"exampleTest2": {[]int{3, 2, 3, 4, 2}, []int{-2, -1, 0, 2, 3}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := distinctDifferenceArray(ts.nums)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
