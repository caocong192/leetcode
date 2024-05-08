package wateringPlants

import (
	"reflect"
	"testing"
)

// TestWateringPlants 单元测试
func TestWateringPlants(t *testing.T) {
	type test struct { // 定义test结构体
		plants   []int
		capacity int
		want     int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{2, 2, 3, 3}, 5, 14},
		"exampleTest2": {[]int{1, 1, 1, 4, 2, 3}, 4, 30},
		"exampleTest3": {[]int{7, 7, 7, 7, 7, 7, 7}, 8, 49},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := wateringPlants(ts.plants, ts.capacity)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
