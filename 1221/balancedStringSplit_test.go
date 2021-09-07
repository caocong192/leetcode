package balancedStringSplit

import (
	"reflect"
	"testing"
)

// TestBalancedStringSplit 单元测试
func TestBalancedStringSplit(t *testing.T) {
	type test struct { // 定义test结构体
		str  string
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"RLRRLLRLRL", 4},
		"exampleTest2": {"RLLLLRRRLR", 3},
		"exampleTest3": {"LLLLRRRR", 1},
		"exampleTest4": {"RLRRRLLRLL", 2},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := balancedStringSplit(ts.str)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
