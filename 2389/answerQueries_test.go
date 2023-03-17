package answerQueries

import (
	"reflect"
	"testing"
)

// TestAnswerQueries 单元测试
func TestAnswerQueries(t *testing.T) {
	type test struct { // 定义test结构体
		nums    []int
		queries []int
		want    []int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{4, 5, 2, 1}, []int{3, 10, 21}, []int{2, 3, 4}},
		"exampleTest2": {[]int{2, 3, 4, 5}, []int{1}, []int{0}},
		"exampleTest3": {[]int{736411, 184882, 914641, 37925, 214915}, []int{331244, 273144, 118983, 118252, 305688, 718089, 665450}, []int{0}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := answerQueries(ts.nums, ts.queries)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
