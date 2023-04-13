package longestDecomposition

import (
	"reflect"
	"testing"
)

// TestLongestDecomposition
func TestLongestDecomposition(t *testing.T) {
	type test struct { // 定义test结构体
		s    string
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"ghiabcdefhelloadamhelloabcdefghi", 7},
		"exampleTest2": {"merchant", 1},
		"exampleTest3": {"antaprezatepzapreanta", 11},
		"exampleTest4": {"a", 1},
		"exampleTest5": {"elvtoelvto", 2},
		"exampleTest6": {"aaa", 3},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := longestDecomposition(ts.s)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
