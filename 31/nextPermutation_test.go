package nextPermutation

import (
	"reflect"
	"testing"
)

//  TestNextPermutation 单元测试
func TestNextPermutation(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		want []int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{9, 5, 8, 6, 2}, []int{9, 6, 2, 5, 8}},
		"exampleTest2": {[]int{1, 2, 3}, []int{1, 3, 2}},
		"exampleTest3": {[]int{1, 1, 5}, []int{1, 5, 1}},
		"exampleTest4": {[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		"exampleTest5": {[]int{1}, []int{1}},
		"exampleTest6": {[]int{1, 2}, []int{2, 1}},
		"exampleTest7": {[]int{2, 1}, []int{1, 2}},
		"exampleTest8": {[]int{1, 5, 1}, []int{5, 1, 1}},
		"exampleTest9": {[]int{5, 1, 1}, []int{1, 1, 5}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := nextPermutation(ts.nums)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
