package getLucky

import (
	"reflect"
	"testing"
)

// TestGetLucky 单元测试
func TestGetLucky(t *testing.T) {
	type test struct { // 定义test结构体
		s    string
		k    int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"iiii", 1, 36},
		"exampleTest2": {"leetcode", 2, 6},
		"exampleTest3": {"hvmhoasabaymnmsd", 1, 79},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := getLucky(ts.s, ts.k)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
