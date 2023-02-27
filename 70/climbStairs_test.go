package climbStairs

import (
	"reflect"
	"testing"
)

// TestClimbStairs 单元测试
func TestClimbStairs(t *testing.T) {
	type test struct { // 定义test结构体
		n    int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {2, 2},
		"exampleTest2": {3, 3},
		"exampleTest3": {44, 1134903170},
	}
	for name, ts := range tests {
		t.Run(name, func(s *testing.T) { // 使用t.Run()执行子测试
			got := climbStairs(ts.n)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
