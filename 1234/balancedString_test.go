package balancedString

import (
	"reflect"
	"testing"
)

// TestBalancedString 单元测试
func TestBalancedString(t *testing.T) {
	type test struct { // 定义test结构体
		s    string
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		//"exampleTest1": {"QWER", 0},
		//"exampleTest2": {"QQWE", 1},
		//"exampleTest3": {"QQQW", 2},
		//"exampleTest4": {"QQQQ", 3},
		//"exampleTest5": {"RQQQQWEQ", 3},
		"exampleTest6": {"QWERQWERQWERQQQQ", 3},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := balancedString(ts.s)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
