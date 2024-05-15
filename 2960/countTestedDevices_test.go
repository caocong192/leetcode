package countTestedDevices

import (
	"reflect"
	"testing"
)

// TestCountTestedDevices 单元测试
func TestCountTestedDevices(t *testing.T) {
	type test struct { // 定义test结构体
		batteryPercentages []int
		want               int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{1, 1, 2, 1, 3}, 3},
		"exampleTest2": {[]int{0, 1, 2}, 2},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := countTestedDevices(ts.batteryPercentages)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
