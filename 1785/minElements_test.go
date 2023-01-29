package minElements

import (
	"reflect"
	"testing"
)

// TestMinElements 单元测试
func TestMinElements(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		limit int
		goal int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{1,-10,9,1},100,0, 1},
		"exampleTest2": {[]int{1,-1,1},3,-4, 2},

	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := minElements(ts.nums, ts.limit, ts.goal)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
