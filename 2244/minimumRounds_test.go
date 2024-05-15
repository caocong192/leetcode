package minimumRounds

import (
	"reflect"
	"testing"
)

// TestMinimumRounds 单元测试
func TestMinimumRounds(t *testing.T) {
	type test struct { // 定义test结构体
		tasks []int
		want  int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}, 4},
		"exampleTest2": {[]int{2, 3, 3}, -1},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := minimumRounds(ts.tasks)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
