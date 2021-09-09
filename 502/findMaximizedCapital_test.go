package findMaximizedCapital

import (
	"reflect"
	"testing"
)

// TestFindMaximizedCapital 单元测试
func TestFindMaximizedCapital(t *testing.T) {
	type test struct { // 定义test结构体
		k, w int
		profits, capital []int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {2, 0, []int{1,2,3},[]int{0,1,1},4},
		"exampleTest2": {3, 0, []int{1,2,3},[]int{0,1,2},6},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := findMaximizedCapital(ts.k, ts.w,ts.profits,ts.capital)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
