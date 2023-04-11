package isRobotBounded

import (
	"reflect"
	"testing"
)

// TestIsRobotBounded 单元测试
func TestIsRobotBounded(t *testing.T) {
	type test struct { // 定义test结构体
		s    string
		want bool
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"GGLLGG", true},
		"exampleTest2": {"GG", false},
		"exampleTest3": {"GL", true},
		"exampleTest4": {"GLL", true},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := isRobotBounded(ts.s)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
