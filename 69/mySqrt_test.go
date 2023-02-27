package mySqrt

import (
	"reflect"
	"testing"
)

// TestMySqrt 单元测试
func TestMySqrt(t *testing.T) {
	type test struct { // 定义test结构体
		x    int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {4, 2},
		"exampleTest2": {8, 2},
		"exampleTest3": {0, 0},
		"exampleTest4": {100, 10},
		"exampleTest5": {10000, 100},
		"exampleTest6": {22802, 151},
	}
	for name, ts := range tests {
		t.Run(name, func(s *testing.T) { // 使用t.Run()执行子测试
			got := mySqrt(ts.x)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
