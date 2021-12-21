package removeElement

import (
	"reflect"
	"testing"
)

// TestRemoveElement 单元测试
func TestRemoveElement(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		val int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{3,2,2,3}, 3, 2},
		"exampleTest2": {[]int{0,1,2,2,3,0,4,2}, 2, 5},
		"exampleTest3": {[]int{0,1,2,2,3,2,2,2,3,2,3,4,5,6,2}, 2, 8},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := removeElement(ts.nums, ts.val)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
