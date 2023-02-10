package dieSimulator

import (
	"reflect"
	"testing"
)

// TestDieSimulator 单元测试
func TestDieSimulator(t *testing.T) {
	type test struct { // 定义test结构体
		n       int
		rollMax []int
		want    int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {
			2,
			[]int{1, 1, 2, 2, 2, 3},
			34,
		},

		"exampleTest2": {
			2,
			[]int{1, 1, 1, 1, 1, 1},
			30,
		},

		"exampleTest3": {
			3,
			[]int{1, 1, 1, 2, 2, 3},
			181,
		},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := dieSimulator(ts.n, ts.rollMax)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
