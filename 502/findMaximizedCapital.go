package findMaximizedCapital

import (
	"container/heap"
	"sort"
)

func findMaximizedCapital(k, w int, profits, capital []int) int {
	n := len(profits)
	type pair struct{ c, p int }
	arr := make([]pair, n)

	for i, p := range profits {
		arr[i] = pair{capital[i], p}
	}

	// 对购入资本从小到大排序
	sort.Slice(arr, func(i, j int) bool { return arr[i].c < arr[j].c })

	h := &hp{}
	for cur :=0; k > 0; k-- {
		// cur 小于元素个数,  并且 投入资本arr[cur].c 小于 手里的资本w
		for cur < n && arr[cur].c <= w {
			// 压入堆中
			heap.Push(h, arr[cur].p)
			cur++
		}
		// 如果h的长度为0, 则表明手里的资本w 买不了任何资源了
		if h.Len() == 0 {
			break
		}
		w += heap.Pop(h).(int)
	}
	return w
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool {  // 绑定less方法
	return h.IntSlice[i] > h.IntSlice[j]  // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}

func (h *hp) Pop() interface{} {  // 绑定pop方法，从最后拿出一个元素并返回
	old := h.IntSlice
	n := len(old)
	x := old[n-1]
	h.IntSlice = old[0 : n-1]
	return x
}

func (h *hp) Push(x interface{}) {  // 绑定push方法，插入新元素
	h.IntSlice = append(h.IntSlice, x.(int))
}


