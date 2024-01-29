package splitWordsBySeparator

import (
	"reflect"
	"testing"
)

// TestSplitWordsBySeparator
func TestSplitWordsBySeparator(t *testing.T) {
	type test struct { // 定义test结构体
		words     []string
		separator byte
		want      []string
	}
	const s = "$problem$"
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]string{"one.two.three", "four.five", "six"}, '.', []string{"one", "two", "three", "four", "five", "six"}},
		"exampleTest2": {[]string{"$easy$", s}, '$', []string{"easy", "problem"}},
		"exampleTest3": {[]string{"|||"}, '|', nil},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := splitWordsBySeparator(ts.words, ts.separator)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
