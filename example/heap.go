package main

import (
	"fmt"

	"github.com/shunhui19/dataStructureAlgorithm"
)

func main() {
	h := dataStructureAlgorithm.NewHeap(15)

	// 插入
	h.Insert(4)
	h.Insert(2)
	h.Insert(5)
	h.Insert(6)
	h.Insert(9)
	h.Insert(3)
	h.Insert(15)
	h.Insert(1)
	h.Insert(100)
	h.Insert(11)
	h.Insert(12)

	// 查找最小元素
	fmt.Printf("Min: %d\n", h.FindMin())
	// 删除最小元素
	h.DeleteMin()
}
