package isHappy

import (
	"reflect"
	"testing"
)

// TestIsHappy 单元测试
func TestIsHappy(t *testing.T) {
	type test struct { // 定义test结构体
		n    int
		want bool
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {19, true},
		"exampleTest2": {2, false},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := isHappy(ts.n)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
