package numTilePossibilities

import (
	"reflect"
	"testing"
)

// TestNumTilePossibilities
func TestNumTilePossibilities(t *testing.T) {
	type test struct { // 定义test结构体
		tiles string
		want  int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"AAB", 8},
		//"exampleTest2": {"AAABBC", 188},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := numTilePossibilities(ts.tiles)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
