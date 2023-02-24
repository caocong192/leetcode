package plusOne

import (
	"reflect"
	"testing"
)

// TestPlusOne 单元测试
func TestPlusOne(t *testing.T) {
	type test struct { // 定义test结构体
		digits []int
		want   []int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{1, 2, 3}, []int{1, 2, 4}},
		"exampleTest2": {[]int{1, 2, 9}, []int{1, 3, 0}},
		"exampleTest3": {[]int{0}, []int{1}},
		"exampleTest4": {[]int{9, 9, 9}, []int{1, 0, 0, 0}},
	}
	for name, ts := range tests {
		t.Run(name, func(s *testing.T) { // 使用t.Run()执行子测试
			got := plusOne(ts.digits)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
