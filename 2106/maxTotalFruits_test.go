package maxTotalFruits

import (
	"reflect"
	"testing"
)

// TestMaxTotalFruits 单元测试
func TestMaxTotalFruits(t *testing.T) {
	type test struct { // 定义test结构体
		fruits   [][]int
		startPos int
		k        int
		want     int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[][]int{{2, 8}, {6, 3}, {8, 6}}, 5, 4, 9},
		"exampleTest2": {[][]int{{0, 9}, {4, 1}, {5, 7}, {6, 2}, {7, 4}, {10, 9}}, 5, 4, 14},
		"exampleTest3": {[][]int{{0, 3}, {6, 4}, {8, 5}}, 3, 2, 0},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := maxTotalFruits(ts.fruits, ts.startPos, ts.k)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
