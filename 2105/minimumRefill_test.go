package minimumRefill

import (
	"reflect"
	"testing"
)

// TestMinimumRefill 单元测试
func TestMinimumRefill(t *testing.T) {
	type test struct { // 定义test结构体
		plants    []int
		capacityA int
		capacityB int
		want      int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{2, 2, 3, 3}, 5, 5, 1},
		"exampleTest2": {[]int{2, 2, 3, 3}, 3, 4, 2},
		"exampleTest3": {[]int{5}, 10, 8, 0},
		"exampleTest4": {[]int{2, 1, 1}, 2, 2, 0},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := minimumRefill(ts.plants, ts.capacityA, ts.capacityB)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
