package isPalindrome

import (
	"reflect"
	"testing"
)

//  TestReverse 单元测试
func TestIsPalindrome(t *testing.T) {
	type test struct { // 定义test结构体
		x    int
		want bool
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {121, true},
		"exampleTest2": {-123, false},
		"exampleTest3": {10, false},
		"exampleTest4": {-101, false},
		"exampleTest5": {1234321, true},
		"exampleTest6": {92544529, true},
		"exampleTest7": {5, true},
		"exampleTest8": {0, true},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := isPalindrome(ts.x)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
