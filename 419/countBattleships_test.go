package countBattleships

import (
	"reflect"
	"testing"
)

// TestCountBattleships 单元测试
func TestCountBattleships(t *testing.T) {
	type test struct { // 定义test结构体
		board [][]byte
		want  int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}, 2},
		"exampleTest2": {[][]byte{{'X', '.', 'X'}, {'X', '.', 'X'}}, 2},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := countBattleships(ts.board)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
