package numberOfBoomerangs

import (
	"reflect"
	"testing"
)

// TestNumberOfBoomerangs 单元测试
func TestNumberOfBoomerangs(t *testing.T) {
	type test struct { // 定义test结构体
		points [][]int
		want   int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[][]int{{0, 0}, {1, 0}, {2, 0}}, 2},
		"exampleTest2": {[][]int{{1, 1}, {2, 2}, {3, 3}}, 2},
		"exampleTest3": {[][]int{{1, 1}}, 0},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := numberOfBoomerangs(ts.points)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
