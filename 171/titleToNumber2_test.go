package titleToNumber

import (
	"reflect"
	"testing"
)

//  TestTitleToNumber2 单元测试
func TestTitleToNumber2(t *testing.T) {
	type test struct { // 定义test结构体
		n    string
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"AB", 28},
		"exampleTest2": {"A", 1},
		"exampleTest3": {"ZY", 701},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := titleToNumber2(ts.n)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
