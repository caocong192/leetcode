package maxLength

import (
	"reflect"
	"testing"
)

// TestMaxLength 单元测试
func TestMaxLength(t *testing.T) {
	type test struct { // 定义test结构体
		arr  []string
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]string{"un", "iq", "ue"}, 4},
		"exampleTest2": {[]string{"cha", "r", "act", "ers"}, 6},
		"exampleTest3": {[]string{"abcdefghijklmnopqrstuvwxyz"}, 26},
		"exampleTest4": {[]string{""}, 0},
		"exampleTest5": {[]string{"a"}, 1},
		"exampleTest6": {[]string{"a", "a", "a", "a"}, 1},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := maxLength(ts.arr)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
