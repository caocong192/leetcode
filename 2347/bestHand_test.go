package bestHand

import (
	"reflect"
	"testing"
)

// TestBestHand 单元测试
func TestBestHand(t *testing.T) {
	type test struct { // 定义test结构体
		ranks []int
		suits []byte
		want  string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{13, 2, 3, 1, 9}, []byte{'a', 'a', 'a', 'a', 'a'}, "Flush"},
		"exampleTest2": {[]int{4, 4, 2, 4, 4}, []byte{'d', 'a', 'a', 'b', 'c'}, "Three of a Kind"},
		"exampleTest3": {[]int{10, 10, 2, 12, 9}, []byte{'a', 'b', 'c', 'a', 'd'}, "Pair"},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := bestHand(ts.ranks, ts.suits)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
