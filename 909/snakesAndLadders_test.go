package snakesAndLadders

import (
	"reflect"
	"testing"
)

// TestSnakesAndLadders 单元测试
func TestSnakesAndLadders(t *testing.T) {
	type test struct { // 定义test结构体
		board [][]int
		want  int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[][]int{
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, 35, -1, -1, 13, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, 15, -1, -1, -1, -1}}, 4},
		"exampleTest2": {[][]int{
			{-1, -1, 19, 10, -1},
			{2, -1, -1, 6, -1},
			{-1, 17, -1, 19, -1},
			{25, -1, 20, -1, -1},
			{-1, -1, -1, -1, 15}}, 2},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := snakesAndLadders(ts.board)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
