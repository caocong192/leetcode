package sortArrayByParity

import (
	"reflect"
	"testing"
)

//  TestSortArrayByParity 单元测试
func TestSortArrayByParity(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		want  []int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{3,1,2,4}, []int{2,4,3,1}},
		"exampleTest2": {[]int{0}, []int{0}},
		"exampleTest3": {[]int{3,1,5,4}, []int{4,1,5,3}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := sortArrayByParity(ts.nums)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
