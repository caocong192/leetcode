package isArraySpecial

import (
	"reflect"
	"testing"
)

// TestIsArraySpecial
func TestIsArraySpecial(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		want bool
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{1}, true},
		"exampleTest2": {[]int{2, 1, 4}, true},
		"exampleTest3": {[]int{4, 3, 1, 6}, false},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := isArraySpecial(ts.nums)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
