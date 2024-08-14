package isArraySpecial

import (
	"reflect"
	"testing"
)

// TestIsArraySpecial
func TestIsArraySpecial(t *testing.T) {
	type test struct { // 定义test结构体
		nums    []int
		queries [][]int
		want    []bool
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{3, 4, 1, 2, 6}, [][]int{{0, 4}}, []bool{false}},
		"exampleTest2": {[]int{4, 3, 1, 6}, [][]int{{0, 2}, {2, 3}}, []bool{false, true}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := isArraySpecial(ts.nums, ts.queries)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
