package scholar

import (
	"fmt"
	"leetcode/utils"
)

func Solution() {
	obj := Constructor()
	obj.Push(3)
	obj.Push(2)
	obj.Push(-3)
	obj.Push(-3)
	obj.Push(-3)
	obj.Push(5)
	obj.Push(6)
	obj.Push(7)
	obj.Push(-5)
	obj.Push(-5)
	obj.Push(-5)
	obj.Push(-5)

	utils.End()
	fmt.Println(obj.Show())
	utils.End()

	fmt.Println("Top:", obj.Top())
	fmt.Println("Min:", obj.GetMin())

	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()

	utils.End()
	fmt.Println(obj.Show())
	fmt.Println("Min", obj.GetMin())
	utils.End()
}
