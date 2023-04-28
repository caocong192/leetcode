package DinnerPlates

import (
	"sort"
)

type DinnerPlates struct {
	stack         map[int][]int // hash表 存储stack对应的值
	unFilledStack []int         // 保存不满的stack
	maxStack      int           // 标记使用的最大stack
	capacity      int           // 每个stack的存储能力
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{
		stack:         map[int][]int{},
		unFilledStack: []int{0},
		maxStack:      0,
		capacity:      capacity,
	}
}

func (this *DinnerPlates) Push(val int) {
	// 从小打到排序 未满的栈切片
	sort.Ints(this.unFilledStack)

	// 遍历未满的栈
	for i, s := range this.unFilledStack {
		this.stack[s] = append(this.stack[s], val)
		// 如果满了, 剔除
		if len(this.stack[s]) == this.capacity {
			this.unFilledStack = append(this.unFilledStack[:i], this.unFilledStack[i+1:]...)
		}
		return
	}

	// 如果没有可用的栈, 最大栈的值自增1
	this.maxStack++
	this.stack[this.maxStack] = append(this.stack[this.maxStack], val)
	// 如果该栈未满, 加入到未满的栈中
	if len(this.stack[this.maxStack]) < this.capacity {
		this.unFilledStack = append(this.unFilledStack, this.maxStack)
	}
	return
}

func (this *DinnerPlates) Pop() (ans int) {
	// 从最大的栈 开始向下遍历
	for ; this.maxStack >= 0; this.maxStack-- {
		n := len(this.stack[this.maxStack])
		// 如果栈长度大于1, 剔除一个元素后该栈还是不为空
		if n > 1 {
			ans = this.stack[this.maxStack][n-1]
			this.stack[this.maxStack] = this.stack[this.maxStack][0 : n-1]
			// 判断该位置 是否在 未满的栈中, 如果在则跳过
			for _, v := range this.unFilledStack {
				if v == this.maxStack {
					return
				}
			}
			// 加入未满的栈
			this.unFilledStack = append(this.unFilledStack, this.maxStack)
			return
		} else if n == 1 {
			// 剔除元素后 栈为空
			ans = this.stack[this.maxStack][0]
			for index, v := range this.unFilledStack {
				// 如果该位置 在不满的栈中, 则剔除
				if v == this.maxStack {
					this.unFilledStack = append(this.unFilledStack[0:index], this.unFilledStack[index+1:]...)
				}
			}
			// 最大的栈标记 自减
			this.maxStack--
			return
		}
	}
	return -1
}

func (this *DinnerPlates) PopAtStack(index int) int {
	if this.maxStack < index {
		return -1
	}

	n := len(this.stack[index])
	// 如果栈不为空
	if n > 0 {
		// 取值最后一位, 并减少该栈这个元素
		ans := this.stack[index][n-1]
		this.stack[index] = this.stack[index][:n-1]
		// 如果该位置 已经在不满的栈中, 则跳过
		for _, v := range this.unFilledStack {
			if v == index {
				return ans
			}
		}
		// 否则加入不满的栈中
		this.unFilledStack = append(this.unFilledStack, index)
		return ans
	}
	return -1
}

/**
 * Your DinnerPlates object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAtStack(index);
 */
