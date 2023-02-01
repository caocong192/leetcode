package decodeMessage

import (
	"reflect"
	"testing"
)

// TestDecodeMessage 单元测试
func TestDecodeMessage(t *testing.T) {
	type test struct { // 定义test结构体
		key     string
		message string
		want    string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv", "this is a secret"},
		"exampleTest2": {"eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb", "the five boxing wizards jump quickly"},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := decodeMessage(ts.key, ts.message)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
