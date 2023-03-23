package checkArithmeticSubarrays

import (
	"reflect"
	"testing"
)

// TestCheckArithmeticSubarrays 单元测试
func TestCheckArithmeticSubarrays(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		l    []int
		r    []int
		want []bool
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}, []bool{true, false, true}},
		"exampleTest2": {[]int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10}, []int{0, 1, 6, 4, 8, 7}, []int{4, 4, 9, 7, 9, 10}, []bool{false, true, false, false, true, true}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := checkArithmeticSubarrays(ts.nums, ts.l, ts.r)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
