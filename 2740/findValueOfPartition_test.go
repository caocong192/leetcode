package findValueOfPartition

import (
	"reflect"
	"testing"
)

// TestFindValueOfPartition
func TestFindValueOfPartition(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		want int
	}
	const s = "$problem$"
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{1, 3, 2, 4}, 1},
		"exampleTest2": {[]int{100, 1, 10}, 9},
		"exampleTest3": {[]int{84, 11, 100, 100, 75}, 0},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试go
			got := findValueOfPartition(ts.nums)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
