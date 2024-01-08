package numOfBurgers

import (
	"reflect"
	"testing"
)

// TestNumOfBurgers 单元测试
func TestNumOfBurgers(t *testing.T) {
	type test struct { // 定义test结构体
		tomatoSlices int
		cheeseSlices int
		want         []int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {16, 7, []int{1, 6}},
		"exampleTest2": {17, 4, []int{}},
		"exampleTest3": {4, 17, []int{}},
		"exampleTest4": {0, 0, []int{0, 0}},
		"exampleTest5": {2, 1, []int{0, 1}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := numOfBurgers(ts.tomatoSlices, ts.cheeseSlices)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
