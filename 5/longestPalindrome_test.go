package longestPalindrome

import (
	"reflect"
	"testing"
)

//  TestLongestPalindrome 单元测试
func TestLongestPalindrome(t *testing.T) {
	type test struct { // 定义test结构体
		s    string
		want string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"babad", "bab"},
		"exampleTest2": {"cbbd", "bb"},
		"exampleTest3": {"a", "a"},
		"exampleTest4": {"ac", "a"},
		"exampleTest5": {"12321", "12321"},
		"exampleTest6": {"9082123321sf3r", "123321"},
		"exampleTest7": {"aa", "aa"},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := longestPalindrome(ts.s)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
