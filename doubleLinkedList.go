package dataStructureAlgorithm

import (
	"log"
)

// nodes 双链表节点.
type doubleNode struct {
	// Value 节点存储的值
	Value interface{}
	// next 指向后继节点指针
	next *doubleNode
	// Prev 指向前驱节点指针
	Prev *doubleNode
}

// 关键点：
// 1. 插入操作，需要指出该节点的前驱和后继, A->B-C, 把新节点D插入BC之间, 则BDC三者之间的关系:
// 		B.next = D // B的后继节点为D
//		B.next.Prev = D // C的前驱节点为D
// 		D.Prev = C.Prev // D的前驱节点为B
// 		D.next = B.next // D的后继节点为C
// 每个节点拥有一个指向前驱节点指针(头节点除外)和一个指向后继节点指针(最后一个节点除外)

// DoubleLinkedList 双链表.
// 和单链表区别不大，增加了一个指向前驱的指针，在插入和删除多了一倍开销，也增加了空间奢求.
type DoubleLinkedList struct {
	header *doubleNode
}

// MakeEmpty 清空链表.
func (dl *DoubleLinkedList) MakeEmpty() {
	dl.header.next = nil
	// NOTE: 这里可能需要遍历每个节点，并设置节点存储的值和指向下一个节点的指针设置为nil,
	// 防止内存leak.
}

// IsEmpty 链表是否为空表.
func (dl *DoubleLinkedList) IsEmpty() bool {
	return dl.header.next == nil
}

// IsLast 判断节点是否为最后一个节点.
func (dl *DoubleLinkedList) IsLast(node *doubleNode) bool {
	return node.next == nil
}

// Find 获取元素节点.
func (dl *DoubleLinkedList) Find(data interface{}) *doubleNode {
	current := dl.header
	for current != nil && current.Value != data {
		current = current.next
	}
	return current
}

// FindPrevious 查找元素的前驱节点.
func (dl *DoubleLinkedList) FindPrevious(data interface{}) *doubleNode {
	current := dl.header
	for current.next != nil && current.next.Value != data {
		current = current.next
	}
	return current
}

// Delete 删除元素节点.
func (dl *DoubleLinkedList) Delete(data interface{}) {
	node := dl.Find(data)
	if node == nil {
		return
	}

	if !dl.IsLast(node) {
		node.next.Prev = node.Prev // data元素的后继节点的前驱指针指向data元素的前驱节点
	}
	node.Prev.next = node.next // data元素的前驱节点的后继指针指向data元素的后继节点
}

// Insert 新元素newData插入到旧元素oldData节点之后.
func (dl *DoubleLinkedList) Insert(oldData, newData interface{}) {
	oldNode := dl.Find(oldData)
	// 旧元素节点不存在.
	if oldNode == nil {
		log.Printf("the element %v does not exist", oldData)
		return
	}

	newNode := &doubleNode{
		Value: newData,
		Prev:  nil,
		next:  nil,
	}
	if !dl.IsLast(oldNode) {
		oldNode.next.Prev = newNode
	}
	newNode.next = oldNode.next
	newNode.Prev = oldNode
	oldNode.next = newNode
}

// Range 遍历链表节点值.
// NOTE: 遍历的是链表的副本.
func (dl *DoubleLinkedList) Range(f func(value interface{}) bool) {
	current := *dl.header
	for current.next != nil {
		if !f(current.next.Value) {
			break
		}
		current.next = current.next.next
	}
}

// NewLinkedList 初始化空链表.
func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{header: &doubleNode{
		Value: "header",
		next:  nil,
		Prev:  nil,
	}}
}
