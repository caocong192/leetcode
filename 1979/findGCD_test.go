package findGCD

import (
	"reflect"
	"testing"
)

// TestFindGCD 单元测试
func TestFindGCD(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{7, 5, 6, 8, 3}, 1},
		"exampleTest2": {[]int{3, 3}, 3},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := findGCD(ts.nums)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
