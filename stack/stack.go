package stack

import "container/list"

type Stack struct {
	list *list.List
}

func NewStack() Stack {
	return Stack{
		list: list.New(),
	}
}

func (this Stack) Push(v interface{}) {
	this.list.PushBack(v)
}

// 查看栈顶的元素
func (this Stack) Peek() interface{} {
	if this.Len() == 0 {
		return nil
	}
	return this.list.Back().Value
}

func (this Stack) Pop() interface{} {
	if this.Len() == 0 {
		return nil
	}
	e := this.list.Back()
	this.list.Remove(e)
	return e.Value
}

func (this Stack) Len() int {
	return this.list.Len()
}
