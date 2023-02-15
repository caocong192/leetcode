package isGoodArray

import (
	"reflect"
	"testing"
)

// TestIsGoodArray 单元测试
func TestIsGoodArray(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		want bool
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{12, 5, 7, 23}, true},
		"exampleTest2": {[]int{29, 6, 10}, true},
		"exampleTest3": {[]int{3, 6}, false},
		"exampleTest4": {[]int{2, 6, 10, 12, 18, 100, 1000}, false},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := isGoodArray(ts.nums)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
