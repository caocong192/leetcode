package myAtoi

import (
	"reflect"
	"testing"
)

//  TestMyAtoi 单元测试
func TestMyAtoi(t *testing.T) {
	type test struct { // 定义test结构体
		s    string
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"42", 42},
		"exampleTest2": {"   -42", -42},
		"exampleTest3": {"4193 with words", 4193},
		"exampleTest4": {"words and 987", 0},
		"exampleTest5": {"-91283472332", -2147483648},
		"exampleTest6": {"91283472332", 2147483647},
		"exampleTest7": {"91283.472332", 91283},
		"exampleTest8": {"123 456", 123},
		"exampleTest9": {"+-12", 0},
		"exampleTest10": {"  ", 0},

	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := myAtoi(ts.s)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
