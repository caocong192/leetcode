package groupThePeople

import (
	"reflect"
	"testing"
)

// TestGroupThePeople 单元测试
func TestGroupThePeople(t *testing.T) {
	type test struct { // 定义test结构体
		groupSizes []int
		want  [][]int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{3,3,3,3,3,1,3}, [][]int{{5}, {0,1,2},{3,4,6}}},
		"exampleTest2": {[]int{2,1,3,3,3,2}, [][]int{{1}, {0,5},{2,3,4}}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := groupThePeople(ts.groupSizes)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
