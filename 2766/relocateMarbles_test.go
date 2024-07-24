package relocateMarbles

import (
	"reflect"
	"testing"
)

// TestRelocateMarbles
func TestRelocateMarbles(t *testing.T) {
	type test struct { // 定义test结构体
		nums     []int
		moveFrom []int
		moveTo   []int
		want     []int
	}
	const s = "$problem$"
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{1, 6, 7, 8}, []int{1, 7, 2}, []int{2, 9, 5}, []int{5, 6, 8, 9}},
		"exampleTest2": {[]int{1, 1, 3, 3}, []int{1, 3}, []int{2, 2}, []int{2}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := relocateMarbles(ts.nums, ts.moveFrom, ts.moveTo)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
