package dataStructureAlgorithm

import (
	"log"
)

// nodes 单链表节点.
type node struct {
	// Value 节点存储的值
	Value interface{}
	// next 指向下一个节点指针
	next *node
}

// 链表是由一系列节点组成的集合，每个节点使用一个对象的引用指向它的后继节点。
// 设置头接点的目的：当需要删除第一个节点时，链表的起始端就需要改变，会增加代码复杂度.
// Header->Node1->Node2->Node3->....
// 关键点：
// 1. 插入操作，A->B->C, 把新节点D插入到BC之间，则BDC三者之间的关系为:
//  B.next = D // B节点的后继节点指向新节点D
// 	D.next = B.next // 新节点D的后继节点指向B的后继节点C
// 2. 删除操作，A->B->D->C, 删除节点D，则BDC三者之间的关系为:
// 	B.next = C // B的后继节点指向节点C
// DoubleLinkedList 单链表.
type LinkedList struct {
	// header 头节点
	header *node
}

// MakeEmpty 清空链表.
func (ll *LinkedList) MakeEmpty() {
	ll.header.next = nil
	// NOTE: 这里可能需要遍历每个节点，并设置节点存储的值和指向下一个节点的指针设置为nil,
	// 防止内存leak.
}

// IsEmpty 链表是否为空表.
func (ll *LinkedList) IsEmpty() bool {
	return ll.header.next == nil
}

// IsLast 判断节点是否为最后一个节点.
func (ll *LinkedList) IsLast(node *node) bool {
	return node.next == nil
}

// Find 获取元素节点.
func (ll *LinkedList) Find(data interface{}) *node {
	current := ll.header
	for current != nil && current.Value != data {
		current = current.next
	}
	return current
}

// FindPrevious 查找元素的前驱节点.
func (ll *LinkedList) FindPrevious(data interface{}) *node {
	current := ll.header
	for current.next != nil && current.next.Value != data {
		current = current.next
	}
	return current
}

// Delete 删除元素节点.
func (ll *LinkedList) Delete(data interface{}) {
	p := ll.FindPrevious(data)

	// 判断前驱元素是否为最后一个节点，这里有两种情况:
	// 1.空链表, 则待删除的元素不存在
	// 2.非空链表，若前驱为最后一个节点了，该元素也不存在
	if !ll.IsLast(p) {
		p.next = p.next.next
	}
}

// Insert 新元素newData插入到旧元素oldData节点之后.
func (ll *LinkedList) Insert(oldData, newData interface{}) {
	oldNode := ll.Find(oldData)
	// 旧元素节点不存在.
	if oldNode == nil {
		log.Printf("the element %v does not exist", oldData)
		return
	}

	newNode := &node{
		Value: newData,
		next:  nil,
	}
	newNode.next = oldNode.next
	oldNode.next = newNode
}

// Range 遍历链表节点值.
// NOTE: 遍历的是链表的副本.
func (ll *LinkedList) Range(f func(value interface{}) bool) {
	current := *ll.header
	for current.next != nil {
		if !f(current.next.Value) {
			break
		}
		current.next = current.next.next
	}
}

// NewLinkedList 初始化空链表.
func NewLinkedList() *LinkedList {
	return &LinkedList{header: &node{
		Value: "header",
		next:  nil,
	}}
}
