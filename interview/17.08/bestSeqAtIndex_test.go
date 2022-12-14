package bestSeqAtIndex


import (
"reflect"
"testing"
)

// TestBestSeqAtIndex 单元测试
func TestBestSeqAtIndex(t *testing.T) {
	type test struct { // 定义test结构体
		height []int
		weight []int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{65,70,56,75,60,68}, []int{100,150,90,190,95,110}, 6},
		"exampleTest2": {[]int{65,70,56,75,60,68,65}, []int{100,150,90,190,95,110,101}, 6},

	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := bestSeqAtIndex(ts.height,ts.weight)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
