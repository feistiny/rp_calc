package stack

import (
	"log"
	"testing"
)

func TestPushAndPop(t *testing.T) {
	s := NewStack()
	if s.Len() != 0 {
		log.Fatal("空栈长度不是 0")
	}
	elem := 1
	s.Push(elem)
	ePeek := s.Peek()
	if elem != ePeek {
		log.Fatal("peek 的元素跟 push 进去的不一样")
	}
	if s.Len() != 1 {
		log.Fatal("空栈 push 后长度不是 1")
	}
	eRet := s.Pop()
	if elem != eRet {
		log.Fatal("pop 出来的跟 push 进去的不一样")
	}
	if s.Len() != 0 {
		log.Fatal("pop 出来后栈不是空")
	}
}

func TestPopFromEmptyStack(t *testing.T) {
	s := NewStack()
	ePeek := s.Peek()
	if nil != ePeek {
		log.Fatal("空栈 peek 出来的应该是 nil")
	}
	eRet := s.Pop()
	if nil != eRet {
		log.Fatal("空栈 pop 出来的应该是 nil")
	}
	if s.Len() != 0 {
		log.Fatal("空栈 pop 出来后栈的长度不是 0")
	}
}