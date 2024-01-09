package containsNearbyDuplicate

import (
	"reflect"
	"testing"
)

// TestContainsNearbyDuplicate
func TestContainsNearbyDuplicate(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		k    int
		want bool
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{1, 2, 3, 1}, 3, true},
		"exampleTest2": {[]int{1, 0, 1, 1}, 1, true},
		"exampleTest3": {[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := containsNearbyDuplicate(ts.nums, ts.k)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
