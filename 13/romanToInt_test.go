package romanToInt

import (
	"reflect"
	"testing"
)

//  TestRomanToInt 单元测试
func TestRomanToInt(t *testing.T) {
	type test struct { // 定义test结构体
		x    string
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"III", 3},
		"exampleTest2": {"IV", 4},
		"exampleTest3": {"IX", 9},
		"exampleTest4": {"LVIII", 58},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := romanToInt(ts.x)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
