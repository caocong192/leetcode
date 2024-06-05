package arrangeCoins

import (
	"reflect"
	"testing"
)

// TestArrangeCoins 单元测试
func TestArrangeCoins(t *testing.T) {
	type test struct { // 定义test结构体
		n    int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {5, 2},
		"exampleTest2": {8, 3},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := arrangeCoins(ts.n)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
