package minimumRecolors

import (
	"reflect"
	"testing"
)

// TestMinimumRecolors 单元测试
func TestMinimumRecolors(t *testing.T) {
	type test struct { // 定义test结构体
		blocks string
		k      int
		want   int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"WBWBBBW", 2, 0},
		"exampleTest2": {"WBBWWBBWBW", 7, 3},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := minimumRecolors(ts.blocks, ts.k)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
