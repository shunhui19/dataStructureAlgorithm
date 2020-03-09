package dataStructureAlgorithm

import (
	"log"
)

// MinTableSize 散列表空间最小值
const MinTableSize = 10

// HashTable 散列表.
// 散列是一种用于以常数平均时间执行插入，删除和查找的技术。
// 底层数据结构是一个包含关键字的具有固定大小的数组，所以插入，更新和
// 查找速度非常快.
// 关键点：如何实现一个均匀分配关键字到存储单元的函数, 特点如下：
// 1.如何均匀分配关键字到各个单元
// 2.如何处理冲突(collision), 当两个关键字散列到同一个值时, 可采用"分离链式法" 和 "开放定址法"
// 		链式法，就是每个存储单元对应一个链表，链表中存储具有相同的hash(key)值的元素
// 3.如何确定散列表的大小
type HashTable struct {
	// size 散列表的大小，即底层数组大小.
	size int
	// link 链表，存储相同散列值数据.
	link []*hashNode
}

// hashNode 散列表中链表，用于处理冲突的散列值.
// 将散列到同一个值的所有元素保留到这个链表中去, 并且新元素保留在表头位置(即第一个位置， 不是header节点)
type hashNode struct {
	// Value 链表中的元素值
	Value interface{}
	// next 指向下一个节点的指针
	next *hashNode
}

// Insert 存储一个元素data到散列表中, 类似链表在表头执行插入操作.
func (h *HashTable) Insert(key string) {
	node, pos := h.find(key)
	if node != nil {
		return
	}

	newNode := &hashNode{
		Value: key,
		next:  nil,
	}

	if h.link[pos].next == nil {
		h.link[pos].next = newNode
		return
	}

	// 新节点指向header节点指向的下一个元素节点
	newNode.next = h.link[pos].next
	// header节点指向新插入的元素节点
	h.link[pos].next = newNode
}

// Get 获取指定的值
func (h *HashTable) Get(key string) interface{} {
	node, _ := h.find(key)
	if node != nil {
		return node.Value
	}
	return nil
}

// Delete 从散列表中删除一个元素.
func (h *HashTable) Delete(key string) {
	pos := h.hash(key)
	l := h.link[pos]
	for l.next != nil {
		if l.next.Value == key {
			if l.next.next != nil {
				l.next = l.next.next
			} else {
				l.next = nil
			}
			return
		}
		l = l.next
	}
}

// find 查找一个元素data节点和链表所在位置.
func (h *HashTable) find(key string) (*hashNode, int) {
	pos := h.hash(key)
	l := h.link[pos].next
	for l != nil && l.Value != key {
		l = l.next
	}

	return l, pos
}

// hash 把字符串的每个字符ASCII码值*32再求和, 最后执行 sum % tableSize 求模运算.
func (h *HashTable) hash(key string) int {
	var sum = 0
	for _, s := range key {
		sum += int(s) << 5
	}
	return sum % h.size
}

func NewHashTable(tableSize int) *HashTable {
	if tableSize < MinTableSize {
		log.Println("table size too small, the min is", MinTableSize)
		return nil
	}

	ht := &HashTable{
		size: tableSize,
		link: make([]*hashNode, tableSize),
	}

	// 每个散列存储单元对应一个链表, 以下标作为标记
	for i := 0; i < tableSize; i++ {
		ht.link[i] = &hashNode{
			Value: "header",
			next:  nil,
		}
	}
	return ht
}
