package hammingWeight

import (
	"reflect"
	"testing"
)

//  TestHammingWeight 单元测试
func TestHammingWeight(t *testing.T) {
	type test struct { // 定义test结构体
		x    uint32
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {00000000000000000000000000001011, 3},
		"exampleTest2": {00000000000000000000000010000000, 1},
		"exampleTest3": {11111111111111111111111111111101, 31},
		"exampleTest4": {00000000000000000000000000000000, 0},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := hammingWeight(ts.x)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
