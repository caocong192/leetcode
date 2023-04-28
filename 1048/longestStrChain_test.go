package longestStrChain

import (
	"reflect"
	"testing"
)

// TestLongestStrChain
func TestLongestStrChain(t *testing.T) {
	type test struct { // 定义test结构体
		words []string
		want  int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]string{"a", "b", "ba", "bca", "bda", "bdca"}, 4},
		"exampleTest2": {[]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}, 5},
		"exampleTest3": {[]string{"abcd", "dbqca"}, 1},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := longestStrChain(ts.words)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
