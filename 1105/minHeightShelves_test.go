package minHeightShelves

import (
	"reflect"
	"testing"
)

// TestMinHeightShelves
func TestMinHeightShelves(t *testing.T) {
	type test struct { // 定义test结构体
		books      [][]int
		shelfWidth int
		want       int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, 4, 6},
		"exampleTest2": {[][]int{{1, 3}, {2, 4}, {3, 2}}, 6, 4},
		"exampleTest3": {[][]int{{1, 1}, {1, 1}, {3, 3}, {3, 4}, {3, 2}, {1, 1}}, 3, 11},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := minHeightShelves(ts.books, ts.shelfWidth)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
