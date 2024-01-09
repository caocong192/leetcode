package canConstruct

import (
	"reflect"
	"testing"
)

// TestCanConstruct
func TestCanConstruct(t *testing.T) {
	type test struct { // 定义test结构体
		ransomNote string
		magazine   string
		want       bool
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"a", "b", false},
		"exampleTest2": {"aa", "ab", false},
		"exampleTest3": {"aa", "aab", true},
		"exampleTest4": {"aabcaad", "aabcbdaaa", true},
		"exampleTest5": {"naabcaad", "aabcbdaaa", false},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := canConstruct(ts.ransomNote, ts.magazine)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
