package main

import (
	"fmt"
	DinnerPlates "github.com/leetcood/1172"
)

func main() {

	//str := []string{"DinnerPlates", "push", "push", "push", "push", "push", "popAtStack", "push", "push", "popAtStack", "popAtStack", "pop", "pop", "pop", "pop", "pop"}

	// CASE 1
	c := DinnerPlates.Constructor(2)
	c.Push(1)
	c.Push(2)
	c.Push(3)
	c.Push(4)
	c.Push(5)
	fmt.Println(c.PopAtStack(0))
	c.Push(20)
	c.Push(21)
	fmt.Println(c.PopAtStack(0))
	fmt.Println(c.PopAtStack(2))
	fmt.Println(c.Pop())
	fmt.Println(c.Pop())
	fmt.Println(c.Pop())
	fmt.Println(c.Pop())
	fmt.Println(c.Pop())

	// CASE 2
	//c := DinnerPlates.Constructor(1)
	//c.Push(1)
	//c.Push(2)
	//println(c.PopAtStack(1))
	//println(c.Pop())
	//c.Push(1)
	//c.Push(2)
	//println(c.Pop())
	//println(c.Pop())

	// CASE 3
	//c := DinnerPlates.Constructor(2)
	//c.Push(1)
	//c.Push(2)
	//c.Push(3)
	//c.Push(4)
	//c.Push(5)
	//println(c.Pop())
	//println(c.Pop())
	//println(c.Pop())
	//println(c.Pop())
	//c.Push(2)
	//c.Push(3)
	//println(c.Pop())

}
