package fillCups

import (
	"reflect"
	"testing"
)

// TestFillCups 单元测试
func TestFillCups(t *testing.T) {
	type test struct { // 定义test结构体
		amount []int
		want   int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{1, 4, 2}, 4},
		"exampleTest2": {[]int{5, 4, 4}, 7},
		"exampleTest3": {[]int{5, 0, 0}, 5},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := fillCups(ts.amount)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
