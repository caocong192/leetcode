package maxProfit

import (
	"reflect"
	"testing"
)

// TestMaxProfit 单元测试
func TestMaxProfit(t *testing.T) {
	type test struct { // 定义test结构体
		prices []int
		want   int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{7, 1, 5, 3, 6, 4}, 5},
		"exampleTest2": {[]int{7, 6, 4, 3, 1}, 0},
		"exampleTest3": {[]int{1}, 0},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := maxProfit(ts.prices)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
