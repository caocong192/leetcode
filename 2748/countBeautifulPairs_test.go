package countBeautifulPairs

import (
	"reflect"
	"testing"
)

// TestCountBeautifulPairs 单元测试
func TestCountBeautifulPairs(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{2, 5, 1, 4}, 5},
		//"exampleTest2": {[]int{11, 21, 12}, 2},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := countBeautifulPairs(ts.nums)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
