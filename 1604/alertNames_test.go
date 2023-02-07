package alertNames

import (
	"reflect"
	"testing"
)

// TestAlertNames 单元测试
func TestAlertNames(t *testing.T) {
	type test struct { // 定义test结构体
		keyName []string
		keyTime []string
		want    []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {
			[]string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"},
			[]string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"},
			[]string{"daniel"},
		},

		"exampleTest2": {
			[]string{"alice", "alice", "alice", "bob", "bob", "bob", "bob"},
			[]string{"12:01", "12:00", "18:00", "21:00", "21:20", "21:30", "23:00"},
			[]string{"bob"},
		},

		"exampleTest3": {
			[]string{"john", "john", "john"},
			[]string{"23:58", "23:59", "00:01"},
			[]string{},
		},

		"exampleTest4": {
			[]string{"leslie", "leslie", "leslie", "clare", "clare", "clare", "clare"},
			[]string{"13:00", "13:20", "14:00", "18:00", "18:51", "19:30", "19:49"},
			[]string{"clare", "leslie"},
		},

		"exampleTest5": {
			[]string{"a","a","a","a","a","b","b","b","b","b","b"},
			[]string{"04:48","23:53","06:36","07:45","12:16","00:52","10:59","17:16","00:36","01:26","22:42"},
			[]string{"b"},
		},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := alertNames(ts.keyName, ts.keyTime)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
