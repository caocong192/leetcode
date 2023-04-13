package mostFrequentEven

import (
	"reflect"
	"testing"
)

// TestMostFrequentEven
func TestMostFrequentEven(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{8154, 9139, 8194, 3346, 5450, 9190, 133, 8239, 4606, 8671, 8412, 6290}, 3346},
		"exampleTest2": {[]int{0, 0, 0, 0}, 0},
		"exampleTest3": {[]int{0, 1, 2, 2, 4, 4, 1}, 2},
		"exampleTest4": {[]int{0, 1, 2, 2, 4, 4, 1}, 2},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := mostFrequentEven(ts.nums)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
