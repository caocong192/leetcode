package lengthOfLongestSubstring

import "testing"
import "reflect"

//  TestLengthOfLongestSubstring 单元测试
func TestLengthOfLongestSubstring(t *testing.T) {
	type test struct { // 定义test结构体
		s    string
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"abcabcbb", 3},
		"exampleTest2": {"bbbbb", 1},
		"exampleTest3": {"pwwkew", 3},
		"exampleTest4": {"au", 2},
		"exampleTest5": {"a12345/.+p", 10},
		"exampleTest6": {"", 0},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := lengthOfLongestSubstring(ts.s)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
