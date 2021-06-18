package reverse

import "testing"
import "reflect"

//  TestReverse 单元测试
func TestReverse(t *testing.T) {
	type test struct { // 定义test结构体
		x    int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {123, 321},
		"exampleTest2": {-123, -321},
		"exampleTest3": {120, 21},
		"exampleTest4": {1534236469, 0},
		"exampleTest5": {-2147483412, -2143847412},

	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := reverse(ts.x)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
