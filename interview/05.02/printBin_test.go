package printBin

import (
	"reflect"
	"testing"
)

// TestPrintBin 单元测试
func TestPrintBin(t *testing.T) {
	type test struct { // 定义test结构体
		num  float64
		want string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {0.625, "0.101"},
		"exampleTest2": {0.1, "ERROR"},
		"exampleTest3": {0.75, "0.11"},
		"exampleTest4": {0.752323, "ERROR"},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := printBin(ts.num)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
