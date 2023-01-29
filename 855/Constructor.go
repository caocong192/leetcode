package Constructor

type ExamRoom struct {
	seats []bool
}

func Constructor(n int) ExamRoom {
	var s ExamRoom

	s.seats = make([]bool, n)
	return s
}

func (this *ExamRoom) Seat() int {
	
}

func (this *ExamRoom) Leave(p int) {
	this.seats[p] = false
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
