package Constructor

import "container/list"

type MinStack struct {
	stack *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: list.New(),
	}
}

func (this *MinStack) Push(x int) {
	this.stack.PushBack(x)
}

func (this *MinStack) Pop() {
	this.stack.Remove(this.stack.Back())
}

func (this *MinStack) Top() int {
	return this.stack.Back().Value.(int)
}

func (this *MinStack) Min() (ans int) {
	if this.stack.Len() == 0 {
		return
	}

	ans = this.stack.Front().Value.(int)
	for i := this.stack.Front().Next(); i != nil; i = i.Next() {
		if i.Value.(int) < ans {
			ans = i.Value.(int)
		}
	}
	return
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
