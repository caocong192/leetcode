package strStr

import (
	"reflect"
	"testing"
)

// TestStrStr 单元测试
func TestStrStr(t *testing.T) {
	type test struct { // 定义test结构体
		haystack string
		needle string
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"hello", "ll", 2},
		"exampleTest2": {"aaaaa", "bba", -1},
		"exampleTest3": {"", "", 0},
		"exampleTest4": {"a", "a", 0},
		"exampleTest5": {"abc", "c", 2},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := strStr(ts.haystack, ts.needle)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
