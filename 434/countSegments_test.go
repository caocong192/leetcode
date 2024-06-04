package countSegments

import (
	"reflect"
	"testing"
)

// TestCountSegments 单元测试
func TestCountSegments(t *testing.T) {
	type test struct { // 定义test结构体
		s    string
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"Hello, my name is John", 5},
		"exampleTest2": {" ", 0},
		"exampleTest3": {"          ", 0},
		"exampleTest4": {"", 0},
		"exampleTest5": {"a", 1},
		"exampleTest6": {"Of all the gin joints in all the towns in all the world,   ", 13},
		"exampleTest7": {", , , ,        a, eaefa", 6},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := countSegments(ts.s)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
