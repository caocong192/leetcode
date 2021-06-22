package convert

import "bytes"

func convert(s string, numRows int) string {

	// 特殊情况: 如果只排列成1行, 或者s只有一个数  则排列后还是本身, 直接返回结果
	if numRows == 1 || len(s) == 1 {
		return s
	}

	// 构造二维数组, 第一个维度为行, 第二个维度保存对应行的值
	m := make([][]byte, numRows, len(s)/numRows+numRows)

	for i := 0; i < len(s); i++ {
		// 从第一行第一个数开始(含) 到第二个数(不含)之间.  一共为 numRow * 2 -2 个数
		// 对i 取模, 得当前i的位置指向第几个数  赋值给pos
		pos := i % (numRows*2 - 2)

		// row 保存当前 i 对应哪一行
		row := 0

		// 如果 pos 小于 行数, 那么是正向排列, pos 即为行数, 直接赋值
		if pos < numRows {
			row = pos
		} else {
			// 否者是逆序排列,
			row = 2*numRows - 2 - pos
		}

		// 将s[i], 追加到对应行的切片中
		m[row] = append(m[row], s[i])
	}

	// 拼接后返回结果
	ret := bytes.Join(m, []byte(""))
	return string(ret)
}

/*
PAYPALISHIRING
P   A   H   N    1 5 9 13     每次+n+1
A P L S I I G    2 4 6 8 19 12 14
Y   I   R        3 7 11

PAYPALISHIRING

PINALSIGYAHRPI

P     I      N
A  L  S   I  G
Y  A  H   R
P     I

*/
