package maximumSumOfHeights

import (
	"reflect"
	"testing"
)

// TestMaximumSumOfHeights 单元测试
func TestMaximumSumOfHeights(t *testing.T) {
	type test struct { // 定义test结构体
		maxHeights []int
		want       int64
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{5, 3, 4, 1, 1}, 13},
		"exampleTest2": {[]int{6, 5, 3, 9, 2, 7}, 22},
		"exampleTest3": {[]int{3, 2, 5, 5, 2, 3}, 18},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := maximumSumOfHeights(ts.maxHeights)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
