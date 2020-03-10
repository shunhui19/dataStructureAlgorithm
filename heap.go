package dataStructureAlgorithm

import (
	"log"
)

// Heap 堆，堆需要满足两个条件：
// 1. 必需为一棵完全二叉树，即除了叶子节点一层外，其它层所有节点都是满的，并且叶子节点这一层靠左填满。
// 2. [大堆]节点值大于等于(或[小堆]小于等于)子树中每个节点的值.
// 因此，堆可以用数组来表示，因而不需要指针, 对于数组中任一位置i上的元素，
// 其左儿子在位置 2i 上， 右儿子在左儿子后的单元 2i+1 中，它的父亲则在位置 i/2 上
type Heap struct {
	// size 当前元素个数
	size int
	// capacity 数组大小
	capacity int
	// element 数组切片
	element []int
}

// MakeEmpty 清空堆元素.
func (h *Heap) MakeEmpty() {
	h.size = 0
	for i := 1; i <= h.size; i++ {
		h.element[i] = 0
	}
}

// Insert 插入一个元素.
// 从下标为1开始存储数据，方便计算
// 插入流程：
// 1.在下一个空闲的位置创建一个空穴，是为了满足完成二叉树性质
// 2.如果元素X放在空穴，满足堆的序，那么插入完成.
// 3.否则，把空穴的父节点的元素移入到该空穴中
// 4.继续步骤3操作直到元素X可以放入正常的位置，满足堆的性质.
func (h *Heap) Insert(data int) {
	if h.IsFull() {
		log.Println("heap is full")
		return
	}

	// 下标从1开始存储数据
	h.size++
	h.element[h.size] = data

	// 最后个位置是 h.size, 创建空穴的位置是 h.size+1, 并且空穴父节点位置为 h.size+1 / 2
	// 如果待插入元素X放到最后个位置，满足堆序要求，直接返回
	for i := h.size; h.element[i>>1] > data; i >>= 1 {
		// 把父节点元素与空穴位置元素交换
		h.element[i], h.element[i>>1] = h.element[i>>1], h.element[i]
	}
}

// FindMin 查找最小元素.
// 最小堆中第一个元素为最小元素
func (h *Heap) FindMin() int {
	return h.element[1]
}

// DeleteMin 删除小最元素.
// 流程：
// 1. 找出最小单元，在小堆中第一个元素为最小值，此时根节点处产生一个空穴
// 2. 由于现在堆中少了一个元素，需要把堆中最后一个元素X移动到堆顶位置
// 3. 比较空穴(堆顶)与左右子树的值, 将空穴的两个儿子中较小者移入空穴，这样就把空穴向下推了一层。
// 4. 重复3步骤，直到X可以被放入空穴中
func (h *Heap) DeleteMin() {
	if h.IsEmpty() {
		return
	}

	// 最后一个元素移动到堆顶
	h.element[1] = h.element[h.size]
	// 删除最后一个元素
	h.element[h.size] = 0
	h.size--

	for i := 1; 2*i <= h.size; i <<= 1 {
		minPos := i
		// 找出左右子树中较小的节点
		if h.element[2*i] < h.element[2*i+1] {
			minPos = 2 * i
		} else {
			minPos = 2*i + 1
		}
		// 较小节点与空穴节点交换值
		if h.element[minPos] < h.element[i] {
			h.element[minPos], h.element[i] = h.element[i], h.element[minPos]
		}
	}
}

// IsFull 堆是否已经满了.
func (h *Heap) IsFull() bool {
	return h.size == h.capacity
}

// IsEmpty 堆是否为空.
func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

func NewHeap(size int) *Heap {
	return &Heap{
		size:     0,
		capacity: size,
		element:  make([]int, size),
	}
}
