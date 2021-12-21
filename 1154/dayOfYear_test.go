package dayOfYear

import (
	"reflect"
	"testing"
)

// TestDayOfYear 单元测试
func TestDayOfYear(t *testing.T) {
	type test struct { // 定义test结构体
		str  string
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"2019-01-09", 9},
		"exampleTest2": {"2019-02-10", 41},
		"exampleTest3": {"2003-03-01", 60},
		"exampleTest4": {"2004-03-01", 61},
		"exampleTest5": {"2000-12-04", 339},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := dayOfYear(ts.str)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}