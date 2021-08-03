package networkDelayTime

import (
	"reflect"
	"testing"
)

// TestNetworkDelayTime 单元测试
func TestNetworkDelayTime(t *testing.T) {
	type test struct { // 定义test结构体
		times [][]int
		n     int
		k     int
		want  int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, 2},
		"exampleTest2": {[][]int{{1, 2, 1}}, 2, 1, 1},
		"exampleTest3": {[][]int{{1, 2, 1}}, 2, 2, -1},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := networkDelayTime(ts.times, ts.n, ts.k)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
