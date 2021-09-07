package isValid

import (
	"reflect"
	"testing"
)

// TestIsValid 单元测试
func TestIsValid(t *testing.T) {
	type test struct { // 定义test结构体
		str  string
		want bool
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"()", true},
		"exampleTest2": {"(){}[]", true},
		"exampleTest3": {"(}", false},
		"exampleTest4": {"({[]})", true},
		"exampleTest5": {"(((((", false},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := isValid(ts.str)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
