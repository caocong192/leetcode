package divide

import (
	"reflect"
	"testing"
)

//  TestDivide 单元测试
func TestDivide(t *testing.T) {
	type test struct { // 定义test结构体
		t    int
		n    int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {7, -3, -2},
		"exampleTest3": {-2147483648, 1, 2147483647},
		"exampleTest4": {10, 3, 3},

	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := divide(ts.t, ts.n)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
