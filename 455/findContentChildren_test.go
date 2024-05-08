package findContentChildren

import (
	"reflect"
	"testing"
)

// TestFindContentChildren
func TestFindContentChildren(t *testing.T) {
	type test struct { // 定义test结构体
		g    []int
		s    []int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{1, 2, 3}, []int{1, 1}, 1},
		"exampleTest2": {[]int{1, 2}, []int{1, 2, 3}, 2},
		"exampleTest3": {[]int{4, 1, 2, 3}, []int{1, 3}, 2},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := findContentChildren(ts.g, ts.s)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
