package minimumSeconds

import (
	"reflect"
	"testing"
)

// TestMinimumSeconds
func TestMinimumSeconds(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		want int
	}
	const s = "$problem$"
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{1, 2, 1, 2}, 1},
		"exampleTest2": {[]int{2, 1, 3, 3, 2}, 2},
		"exampleTest3": {[]int{1, 2, 3, 4, 1, 2, 3, 4}, 2},
		"exampleTest4": {[]int{4, 18}, 1},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := minimumSeconds(ts.nums)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
