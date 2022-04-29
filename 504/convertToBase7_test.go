package convertToBase7

import (
	"reflect"
	"testing"
)

// TestConvertToBase7 单元测试
func TestConvertToBase7(t *testing.T) {
	type test struct { // 定义test结构体
		num  int
		want string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {10, "13"},
		"exampleTest2": {-7, "-10"},
		"exampleTest3": {100, "202"},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := convertToBase7(ts.num)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
