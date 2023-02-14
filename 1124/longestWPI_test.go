package longestWPI

import (
	"reflect"
	"testing"
)

// TestLongestWPI 单元测试
func TestLongestWPI(t *testing.T) {
	type test struct { // 定义test结构体
		hours []int
		want  int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{9, 9, 6, 0, 6, 6, 9}, 3},
		"exampleTest2": {[]int{6, 6, 6}, 0},
		"exampleTest3": {[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 10},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := longestWPI(ts.hours)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
