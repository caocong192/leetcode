package garbageCollection

import (
	"reflect"
	"testing"
)

// TestGarbageCollection 单元测试
func TestGarbageCollection(t *testing.T) {
	type test struct { // 定义test结构体
		garbage []string
		travel  []int
		want    int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]string{"G", "P", "GP", "GG"}, []int{2, 4, 3}, 21},
		"exampleTest2": {[]string{"MMM", "PGM", "GP"}, []int{3, 10}, 37},
		"exampleTest3": {[]string{"G", "G", "G"}, []int{1, 1}, 5},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := garbageCollection(ts.garbage, ts.travel)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
