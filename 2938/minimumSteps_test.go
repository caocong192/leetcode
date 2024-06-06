package minimumSteps

import (
	"reflect"
	"testing"
)

// TestMinimumSteps 单元测试
func TestMinimumSteps(t *testing.T) {
	type test struct { // 定义test结构体
		s    string
		want int64
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"101", 1},
		"exampleTest2": {"100", 2},
		"exampleTest3": {"0001", 0},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := minimumSteps(ts.s)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
