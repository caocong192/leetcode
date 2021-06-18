package smallestGoodBase

import (
	"reflect"
	"testing"
)

// TestSmallestGoodBase 单元测试
func TestSmallestGoodBase(t *testing.T) {
	type test struct { // 定义test结构体
		n    string
		want string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"13", "3"},
		"exampleTest2": {"4681", "8"},
		"exampleTest3": {"1000000000000000000", "999999999999999999"},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := smallestGoodBase(ts.n)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
