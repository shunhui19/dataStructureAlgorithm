package dataStructureAlgorithm

import (
	"log"
)

// 线性表, 是用一段连续的内存存储数据, 存储地址是连续在一起的。
// 时间复杂度为O(n)
// 插入元素时，需要从尾部开始依次向后移动元素
// 删除元素时，从删除元素位置下一个元素开始依次向前移动一个元素
type LineList struct {
	// list 采用切片存储, 只支持整数.
	list []int
	// maxLength 最大长度.
	maxLength int
	// length 当前线性表中元素个数
	length int
}

// NewLineList return instance of LineList.
func NewLineList(size int) *LineList {
	return &LineList{
		list:      make([]int, size),
		maxLength: size,
		length:    0,
	}
}

// Insert insert a value to appoint position.
func (l *LineList) Insert(position int, value int) {
	if l.length >= l.maxLength {
		log.Printf("线性表已经满了!")
		return
	}

	if position < 1 || position > l.maxLength {
		log.Printf("插入位置不在范围内!")
		return
	}

	for i := l.length - 1; i >= position-1; i-- {
		log.Printf("第 {%v} 个位置开始移动元素[%v]", i, l.list[i])
		l.list[i+1] = l.list[i]
	}

	l.list[position-1] = value
	l.length++
}

// Delete 删除指定位置的元素，并返回该位置元素的值
func (l *LineList) Delete(position int) interface{} {
	if l.length == 0 {
		return false
	}

	if l.length < position || position > l.maxLength {
		log.Printf("删除的元素位置超出了范围")
		return false
	}

	element := l.list[position-1]

	// 如果元素不在尾部
	if position < l.length {
		for i := position - 1; i < l.length-1; i++ {
			l.list[i] = l.list[i+1]
		}
	}
	// 这里只是移动元素，相当于后面的元素覆盖前面的元素，所以最后个元素删除
	l.list[l.length-1] = 0
	l.length--

	return element
}

// Clear 清空线性表.
func (l *LineList) Clear() {
	l.list = l.list[:0]
	l.length = 0
}

// Find 查找指定位置的元素
func (l *LineList) Find(position int) int {
	if position < 1 || position > l.maxLength {
		return 0
	}
	return l.list[position-1]
}
