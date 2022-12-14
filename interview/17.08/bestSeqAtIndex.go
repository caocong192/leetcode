package bestSeqAtIndex

import (
	"sort"
)

func bestSeqAtIndex(height []int, weight []int) int {
	n := len(height)
	// 如果长度为0, 返回0
	if n == 0 {
		return 0
	}
	person := make([]Person, n)

	for i := range person {
		person[i].height = height[i]
		person[i].weight = weight[i]
	}

	// 按身高 高->低排序, 再按 体重 轻->重 排序
	sort.Slice(person, func(i, j int) bool {
		// 身高高的在前边，身高相等则体重轻的在前边
		if person[i].height == person[j].height {
			return person[i].weight < person[j].weight
		}
		return person[i].height > person[j].height
	})

	var res []Person

	for _, p := range person{
		// 在结果中找到第一个p不能叠在上面的人, 二分法
		j := sort.Search(len(res), func(i int) bool {
			c := res[i]
			return c.height <= p.height || c.weight <= p.weight
		})

		// 如果相等 追加
		if j == len(res) {
			res = append(res, p)
		} else {
			res[j] = p
		}
	}

	return len(res)
}

type Person struct {
	height int
	weight int
}
