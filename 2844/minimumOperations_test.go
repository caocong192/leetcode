package minimumOperations

import (
	"reflect"
	"testing"
)

// TestMinimumOperations
func TestMinimumOperations(t *testing.T) {
	type test struct { // 定义test结构体
		num  string
		want int
	}
	const s = "$problem$"
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"2245047", 2},
		"exampleTest2": {"2908305", 3},
		"exampleTest3": {"10", 1},
		"exampleTest4": {"111110", 5},
		"exampleTest5": {"1", 1},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := minimumOperations(ts.num)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
