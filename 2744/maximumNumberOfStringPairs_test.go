package maximumNumberOfStringPairs

import (
	"reflect"
	"testing"
)

// TestMaximumNumberOfStringPairs
func TestMaximumNumberOfStringPairs(t *testing.T) {
	type test struct { // 定义test结构体
		words []string
		want  int
	}
	const s = "$problem$"
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]string{"cd", "ac", "dc", "ca", "zz"}, 2},
		"exampleTest2": {[]string{"ab", "ba", "cc"}, 1},
		"exampleTest3": {[]string{"aa", "ab"}, 0},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := maximumNumberOfStringPairs(ts.words)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
