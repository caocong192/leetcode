package MyQueue

// 先入先出队列

type MyQueue struct {
	stack1, stack2 []int
}

type MyQueue struct {
	stack1, stack2 []int
}

func Constructor() (q MyQueue) {
	return
}

func (this *MyQueue) Push(x int) {
	l := len(this.stack2)

	// 如果stack2不为空, 则先将栈弹出到栈stack1
	if l > 0 {
		for l > 0 {
			this.stack1 = append(this.stack1, this.stack2[l-1])
			this.stack2 = this.stack2[:l-1]
			l--
		}
		// 后来的数据x入栈stack1
		this.stack1 = append(this.stack1, x)
		l = len(this.stack1)
		// 再将stack1栈数据 弹出到栈stack2
		for l > 0 {
			this.stack2 = append(this.stack2, this.stack1[l-1])
			this.stack1 = this.stack1[:l-1]
			l--
		}
	} else {
		// 如果栈stack2为空, 直接入栈
		this.stack2 = append(this.stack2, x)
	}

}

func (this *MyQueue) Pop() int {
	l := len(this.stack2)
	res := this.stack2[l-1]
	this.stack2 = this.stack2[:l-1]
	return res
}

func (this *MyQueue) Peek() int {
	return this.stack2[len(this.stack2)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
