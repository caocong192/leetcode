package repeatedStringMatch

import (
	"reflect"
	"testing"
)

// TestRepeatedStringMatch 单元测试
func TestRepeatedStringMatch(t *testing.T) {
	type test struct { // 定义test结构体
		a    string
		b    string
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"abcd", "cdabcdab", 3},
		"exampleTest2": {"a", "aa", 2},
		"exampleTest3": {"a", "a", 1},
		"exampleTest4": {"a", "", 0},
		"exampleTest5": {"abc", "wxyz", -1},
		"exampleTest6": {"abcdef", "ef", 1},
		"exampleTest7": {"abcdef", "", 0},
		"exampleTest8": {"abc", "cabcabca", 4},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := repeatedStringMatch(ts.a, ts.b)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
