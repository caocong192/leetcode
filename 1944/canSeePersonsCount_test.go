package canSeePersonsCount

import (
	"reflect"
	"testing"
)

// TestCanSeePersonsCount 单元测试
func TestCanSeePersonsCount(t *testing.T) {
	type test struct { // 定义test结构体
		heights []int
		want    []int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{10, 6, 8, 5, 11, 9}, []int{3, 1, 2, 1, 1, 0}},
		"exampleTest2": {[]int{5, 1, 2, 3, 10}, []int{4, 1, 1, 1, 0}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := canSeePersonsCount(ts.heights)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
