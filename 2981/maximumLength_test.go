package maximumLength

import (
	"reflect"
	"testing"
)

// TestMaximumLength 单元测试
func TestMaximumLength(t *testing.T) {
	type test struct { // 定义test结构体
		s    string
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"aaaa", 2},
		"exampleTest2": {"abcdef", -1},
		"exampleTest3": {"abcaba", 1},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := maximumLength(ts.s)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
