package findRepeatedDnaSequences

import (
	"reflect"
	"testing"
)

//  TestFindRepeatedDnaSequences 单元测试
func TestFindRepeatedDnaSequences(t *testing.T) {
	type test struct { // 定义test结构体
		s    string
		want []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", []string{"AAAAACCCCC", "CCCCCAAAAA"}},
		"exampleTest2": {"AAAAAAAAAAAAA", []string{"AAAAAAAAAA"}},
		"exampleTest3": {"AAAAAAAAAA", []string{}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := findRepeatedDnaSequences(ts.s)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
