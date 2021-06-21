package convert

import (
	"reflect"
	"testing"
)

// TestConvert 单元测试
func TestConvert(t *testing.T) {
	type test struct { // 定义test结构体
		arr     string
		numRows int
		want    string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		"exampleTest2": {"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		"exampleTest3": {"A", 1, "A"},
		"exampleTest4": {"A", 3, "A"},
		"exampleTest5": {"AB", 3, "AB"},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := convert(ts.arr, ts.numRows)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
