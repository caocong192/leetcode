package main

import (
	"bytes"
	"fmt"
)

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	m := make([][]byte, numRows, len(s)/numRows+numRows)

	for i := 0; i < len(s); i++ {
		// 把所有排序分解成 两列.  一共为 numRow * 2 -2 个数
		// 对i 取模, 得当前i的位置指向第几个数  赋值给pos
		pos := i % (numRows*2 - 2)
		// row 保存当前 i 对应哪一行
		row := 0
		// 如果 pos 小于 行数, 那么是正向排列
		if pos < numRows {
			// pos 即为行数, 直接赋值
			row = pos
		} else {
			// 否者是逆序排列,
			//row = (numRows-1) - (pos - numRows + 1)
			row = 2*numRows - pos - 2
		}
		//fmt.Println(s[i])
		m[row] = append(m[row], s[i])

		//fmt.Println(2)
	}

	ret := bytes.Join(m, []byte(""))
	return string(ret)
}

func main() {
	s := "AB"
	numRows := 3
	want := "AB"

	got := convert(s, numRows)
	fmt.Println(got)
	fmt.Println(want)

}
