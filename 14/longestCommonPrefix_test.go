package longestCommonPrefix

import (
	"reflect"
	"testing"
)

//  TestLongestCommonPrefix 单元测试
func TestLongestCommonPrefix(t *testing.T) {
	type test struct { // 定义test结构体
		x    []string
		want string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]string{"flower", "flow", "flight"}, "fl"},
		"exampleTest2": {[]string{"forever", "foreverlove", "foreve"}, "foreve"},
		"exampleTest3": {[]string{"dog", "racecar", "car"}, ""},
		"exampleTest4": {[]string{"a"}, "a"},
		"exampleTest5": {[]string{"", ""}, ""},
		"exampleTest6": {[]string{"flower","flower","flower","flower"}, "flower"},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := longestCommonPrefix(ts.x)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
