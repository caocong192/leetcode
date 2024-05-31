package findMissingAndRepeatedValues

import (
	"reflect"
	"testing"
)

// TestFindMissingAndRepeatedValues 单元测试
func TestFindMissingAndRepeatedValues(t *testing.T) {
	type test struct { // 定义test结构体
		grid [][]int
		want []int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[][]int{{1, 3}, {2, 2}}, []int{2, 4}},
		"exampleTest2": {[][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}, []int{9, 5}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := findMissingAndRepeatedValues(ts.grid)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
