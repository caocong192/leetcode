package mergeSimilarItems

import (
	"reflect"
	"testing"
)

// TestMergeSimilarItems 单元测试
func TestMergeSimilarItems(t *testing.T) {
	type test struct { // 定义test结构体
		items1 [][]int
		items2 [][]int
		want   [][]int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[][]int{{1, 1}, {4, 5}, {3, 8}}, [][]int{{3, 1}, {1, 5}}, [][]int{{1, 6}, {3, 9}, {4, 5}}},
		"exampleTest2": {[][]int{{1, 1}, {3, 2}, {2, 3}}, [][]int{{2, 1}, {3, 2}, {1, 3}}, [][]int{{1, 4}, {2, 4}, {3, 4}}},
		"exampleTest3": {[][]int{{1, 1}, {2, 2}, {3, 3}}, [][]int{{4, 4}, {5, 6}}, [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 6}}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := mergeSimilarItems(ts.items1, ts.items2)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
