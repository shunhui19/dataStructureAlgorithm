package dataStructureAlgorithm

import "errors"

// stackNode 栈节点.
type stackNode struct {
	// Value 节点存储值.
	Value interface{}
	// next 下一个节点指针.
	next *stackNode
}

// Stack 栈是限制插入和删除只能在一个位置上进行的表，该位置是表的末端，叫栈顶(top).
// 这个栈顶可理解为第一个节点，这样做的好处是，入栈(push)和出栈(pop)操作只需要O(1)时间复杂度.
// 实现栈有两种数据结构，链表和数组，区别是栈会频繁分配和回收内存，而数组是固定大小，只需要分配内存一次.
// 下面是链表结构的栈.
type Stack struct {
	header *stackNode
}

// IsEmpty 判断栈是否为空.
func (s *Stack) IsEmpty() bool {
	return s.header.next == nil
}

// MakeEmpty 清空栈.
func (s *Stack) MakeEmpty() {
	for !s.IsEmpty() {
		s.Pop()
	}
}

// Push 入栈.
func (s *Stack) Push(data interface{}) {
	node := &stackNode{
		Value: data,
		next:  nil,
	}
	node.next = s.header.next
	s.header.next = node
}

// Pop 出栈.
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("empty stack")
	}

	element := s.header.next.Value
	s.header.next = s.header.next.next
	return element, nil
}

// NewStack 创始一个空栈.
func NewStack() *Stack {
	return &Stack{header: &stackNode{
		Value: "header",
		next:  nil,
	}}
}
