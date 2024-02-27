package isUgly

import (
	"reflect"
	"testing"
)

// TestIsUgly
func TestIsUgly(t *testing.T) {
	type test struct { // 定义test结构体
		n    int
		want bool
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {6, true},
		"exampleTest2": {1, true},
		"exampleTest3": {14, false},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := isUgly(ts.n)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
