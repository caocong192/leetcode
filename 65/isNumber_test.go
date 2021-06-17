package isNumber

import (
	"reflect"
	"testing"
)

//  TestIsNumber 单元测试
func TestIsNumber(t *testing.T) {
	type test struct { // 定义test结构体
		s    string
		want bool
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"1", true},
		"exampleTest2": {"0089", true},
		"exampleTest3": {"-0.1", true},
		"exampleTest4": {"-.9", true},
		"exampleTest5": {"3e+7", true},
		"exampleTest6": {"-123.456e789", true},
		"exampleTest7": {"3e+7", true},
		"exampleTest8": {"99e2.5", false},
		"exampleTest9": {"--6", false},
		"exampleTest10": {"1e9", true},
		"exampleTest11": {"1E9", true},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := isNumber(ts.s)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
