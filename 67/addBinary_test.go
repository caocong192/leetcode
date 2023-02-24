package addBinary

import (
	"reflect"
	"testing"
)

// TestAddBinary 单元测试
func TestAddBinary(t *testing.T) {
	type test struct { // 定义test结构体
		a    string
		b    string
		want string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"11", "1", "100"},
		"exampleTest2": {"1010", "1011", "10101"},
	}
	for name, ts := range tests {
		t.Run(name, func(s *testing.T) { // 使用t.Run()执行子测试
			got := addBinary(ts.a, ts.b)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
