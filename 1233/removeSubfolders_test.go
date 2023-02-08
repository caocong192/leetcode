package removeSubfolders

import (
	"reflect"
	"testing"
)

// TestRemoveSubfolders 单元测试
func TestRemoveSubfolders(t *testing.T) {
	type test struct { // 定义test结构体
		folder []string
		want   []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {
			[]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"},
			[]string{"/a", "/c/d", "/c/f"},
		},

		"exampleTest2": {
			[]string{"/a", "/a/b/c", "/a/b/d"},
			[]string{"/a"},
		},

		"exampleTest3": {
			[]string{"/a/b/c", "/a/b/ca", "/a/b/d"},
			[]string{"/a/b/c", "/a/b/ca", "/a/b/d"},
		},

		"exampleTest4": {
			[]string{"/a/b/c", "/a/b/ca", "/a/b/d", "/a/b/c"},
			[]string{"/a/b/c", "/a/b/ca", "/a/b/d"},
		},

		"exampleTest5": {
			[]string{"a/b/cde", "/a/b/c", "/a/b/ca", "/a/b/d", "/a/b/c"},
			[]string{"a/b/cde", "/a/b/c", "/a/b/ca", "/a/b/d"},
		},

		"exampleTest6": {
			[]string{"/aa/ab/ac/ae","/aa/ab/af/ag","/ap/aq/ar/as","/ap/aq/ar","/ap/ax/ay/az","/ap","/ap/aq/ar/at","/aa/ab/af/ah","/aa/ai/aj/ak","/aa/ai/am/ao"},
			[]string{"/aa/ab/ac/ae","/aa/ab/af/ag","/aa/ab/af/ah","/aa/ai/aj/ak","/aa/ai/am/ao","/ap"},
		},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := removeSubfolders(ts.folder)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
