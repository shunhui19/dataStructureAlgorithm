package dataStructureAlgorithm

import (
	"log"
)

// Queue 采用循环数组(circular array)来实现队列, 目的：使用比较小的数组实现队列，节省内存.
// 原理：Front标记为队头下标，Rear标记为队尾下标,
// 元素入队时，size+1, rear+1(有新元素入队了，队尾标记移动下一个位置)
// 元素出队时, size-1, front+1(队头元素已经出队了，队头标记移动下一个位置)
// 特殊情况(Front或Rear到达数组最一个位置时)，当rear下标到达数组最一个位置时，设置rear=0，同时Q[rear]=newElement, Front标记也是一样的
type Queue struct {
	// q 队列数组.
	q []interface{}
	// front 队列头的位置.
	// 默认值为1, 入队的第一元素下标为1.
	front int
	// rear 队列尾的位置.
	rear int
	// size 队列中元素个数.
	size int
	// capacity 队列容量大小.
	capacity int
}

// IsEmpty 队列是否为空.
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// IsFull 队列是否已满.
func (q *Queue) IsFull() bool {
	return q.size == q.capacity
}

// Enqueue 元素入队.
func (q *Queue) Enqueue(data interface{}) {
	if q.IsFull() {
		log.Println("full queue")
		return
	}

	q.size++
	q.rear = q.position(q.rear)
	q.q[q.rear] = data
}

// Dequeue 元素出队.
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		log.Println("empty queue")
		return nil
	}

	element := q.q[q.front]

	q.front = q.position(q.front)
	q.size--

	return element
}

// position 队头Front或队尾Rear下标是否到达数组最一个位置
func (q *Queue) position(position int) int {
	if position++; position == q.capacity {
		position = 0
	}
	return position
}

// MakeEmpty 清空队列.
func (q *Queue) MakeEmpty() {
	q.size, q.front, q.rear = 0, 1, 0
}

func NewQueue(size int) *Queue {
	return &Queue{
		q:        make([]interface{}, size),
		front:    1,
		rear:     0,
		size:     0,
		capacity: size,
	}
}
