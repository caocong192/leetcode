package pathInZigZagTree

import (
	"reflect"
	"testing"
)

//  TestPathInZigZagTree 单元测试
func TestPathInZigZagTree(t *testing.T) {
	type test struct { // 定义test结构体
		label int
		want  []int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {14, []int{1, 3, 4, 14}},
		"exampleTest2": {26, []int{1, 2, 6, 10, 26}},
		"exampleTest3": {11, []int{1, 2, 6, 11}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := pathInZigZagTree(ts.label)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
