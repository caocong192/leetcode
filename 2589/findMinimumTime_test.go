package findMinimumTime

import (
	"reflect"
	"testing"
)

// TestFindMinimumTime 单元测试
func TestFindMinimumTime(t *testing.T) {
	type test struct { // 定义test结构体
		tasks [][]int
		want  int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[][]int{{2, 3, 1}, {4, 5, 1}, {1, 5, 2}}, 2},
		"exampleTest2": {[][]int{{1, 3, 2}, {2, 5, 3}, {5, 6, 2}}, 4},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := findMinimumTime(ts.tasks)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
