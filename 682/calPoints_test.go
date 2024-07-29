package calPoints

import (
	"reflect"
	"testing"
)

// TestCalPoints 单元测试
func TestCalPoints(t *testing.T) {
	type test struct { // 定义test结构体
		operations []string
		want       int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]string{"5", "2", "C", "D", "+"}, 30},
		"exampleTest2": {[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}, 27},
		"exampleTest3": {[]string{"1"}, 1},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := calPoints(ts.operations)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
