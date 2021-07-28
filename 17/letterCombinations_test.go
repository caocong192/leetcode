package letterCombinations

import (
	"reflect"
	"testing"
)

//  TestLetterCombinations 单元测试
func TestLetterCombinations(t *testing.T) {
	type test struct { // 定义test结构体
		digits string
		want   []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		"exampleTest2": {"58", []string{"jt", "ju", "jv", "kt", "ku", "kv", "lt", "lu", "lv"}},
		"exampleTest3": {"", []string{}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := letterCombinations(ts.digits)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
