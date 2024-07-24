package discountPrices

import (
	"reflect"
	"testing"
)

// TestDiscountPrices 单元测试
func TestDiscountPrices(t *testing.T) {
	type test struct { // 定义test结构体
		sentence string
		discount int
		want     string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {"there are $1 $2 and 5$ candies in the shop", 50, "there are $0.50 $1.00 and 5$ candies in the shop"},
		"exampleTest2": {"1 2 $3 4 $5 $6 7 8$ $9 $10$", 100, "1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$"},
		"exampleTest3": {"$76111 ab $6 $", 48, "$39577.72 ab $3.12 $"},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := discountPrices(ts.sentence, ts.discount)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
