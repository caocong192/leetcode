package triangleNumber

import (
	"reflect"
	"testing"
)

// TestTriangleNumber 单元测试
func TestTriangleNumber(t *testing.T) {
	type test struct { // 定义test结构体
		n []int
		want  int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{2,2,3,4}, 3},
		"exampleTest2": {[]int{2,4,4,4,5}, 10},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := triangleNumber(ts.n)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}