package isPalindrome

import (
	"reflect"
	"testing"
)

// TestIsPalindrome 单元测试
func TestIsPalindrome(t *testing.T) {
	type test struct { // 定义test结构体
		s    string
		want bool
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"0P", false},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := isPalindrome(ts.s)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
