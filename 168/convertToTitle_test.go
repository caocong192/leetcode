package convertToTitle

import (
	"reflect"
	"testing"
)

//  TestConvertToTitle 单元测试
func TestConvertToTitle(t *testing.T) {
	type test struct { // 定义test结构体
		columnNumber int
		want         string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {1, "A"},
		"exampleTest2": {28, "AB"},
		"exampleTest3": {701, "ZY"},
		"exampleTest4": {2147483647, "FXSHRXW"},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := convertToTitle(ts.columnNumber)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
