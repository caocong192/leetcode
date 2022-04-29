package Constructor

import "container/list"

type CQueue struct {
	stack1, stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	// 如果栈2 为空,  栈1 的元素 倒序放入栈2
	if this.stack2.Len() == 0 {
		for this.stack1.Len() != 0 {
			e := this.stack1.Back()
			this.stack1.Remove(e)
			this.stack2.PushBack(e.Value)
		}
	}
	if this.stack2.Len() != 0 {
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
