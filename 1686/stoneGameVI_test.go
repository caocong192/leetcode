package stoneGameVI

import (
	"reflect"
	"testing"
)

// TestStoneGameVI 单元测试
func TestStoneGameVI(t *testing.T) {
	type test struct { // 定义test结构体
		aliceValues []int
		bobValues   []int
		want        int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{1, 3}, []int{2, 1}, 1},
		"exampleTest2": {[]int{1, 2}, []int{3, 1}, 0},
		"exampleTest3": {[]int{2, 4, 3}, []int{1, 6, 7}, -1},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := stoneGameVI(ts.aliceValues, ts.bobValues)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
