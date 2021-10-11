package generateParenthesis

import (
	"reflect"
	"testing"
)

//  TestGenerateParenthesis 单元测试
func TestGenerateParenthesis(t *testing.T) {
	type test struct { // 定义test结构体
		n    int
		want []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {1, []string{"()"}},
		"exampleTest2": {3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		"exampleTest3": {2, []string{"(())", "()()"}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := generateParenthesis(ts.n)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
