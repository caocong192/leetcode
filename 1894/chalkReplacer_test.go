package chalkReplacer

import (
	"reflect"
	"testing"
)

// TestChalkReplacer 单元测试
func TestChalkReplacer(t *testing.T) {
	type test struct { // 定义test结构体
		chalk []int
		k     int
		want  int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{5, 1, 5}, 22, 0},
		"exampleTest2": {[]int{3, 4, 1, 2}, 25, 1},
		"exampleTest3": {[]int{1}, 100, 0},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := chalkReplacer(ts.chalk, ts.k)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
