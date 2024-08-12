package minRectanglesToCoverPoints

import (
	"reflect"
	"testing"
)

func TestMinRectanglesToCoverPoints(t *testing.T) {
	type test struct { // 定义test结构体
		points [][]int
		want   int64
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[][]int{{2, 1}, {1, 9}, {1, 4}, {1, 8}, {3, 5}, {4, 6}}, 1, 2},
		"exampleTest2": {[][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}, 2, 3},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := minRectanglesToCoverPoints(ts.points, ts.w)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
