package readBinaryWatch

import (
	"reflect"
	"testing"
)

//  TestReadBinaryWatch 单元测试
func TestReadBinaryWatch(t *testing.T) {
	type test struct { // 定义test结构体
		turnedOn int
		want     []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {1, []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"}},
		"exampleTest2": {9, []string{}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := readBinaryWatch(ts.turnedOn)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
