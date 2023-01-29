package maximumScore

import (
	"reflect"
	"testing"
)

// TestMaximumScore 单元测试
func TestMaximumScore(t *testing.T) {
	type test struct { // 定义test结构体
		a, b, c int
		want    int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {2, 4, 6, 6},
		"exampleTest2": {4, 4, 6, 7},
		"exampleTest3": {1, 8, 8, 8},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := maximumScore(ts.a, ts.b, ts.c)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
