package countWords

import (
	"reflect"
	"testing"
)

// TestCountWords 单元测试
func TestCountWords(t *testing.T) {
	type test struct { // 定义test结构体
		words1 []string
		words2 []string
		want   int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]string{"leetcode", "is", "amazing", "as", "is"}, []string{"amazing", "leetcode", "is"}, 2},
		"exampleTest2": {[]string{"b", "bb", "bbb"}, []string{"a", "aa", "aaa"}, 0},
		"exampleTest3": {[]string{"a", "ab"}, []string{"a", "a", "a", "ab"}, 1},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := countWords(ts.words1, ts.words2)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
