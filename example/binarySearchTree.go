package main

import (
	"fmt"

	"github.com/shunhui19/dataStructureAlgorithm"
)

func main() {
	bst := dataStructureAlgorithm.NewBinarySearchTree()
	// 插入数据
	bst.Insert(5)
	bst.Insert(2)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(4)
	bst.Insert(8)

	// 查找树中的最小值
	fmt.Println(bst.FindMin())
	// 查找树中的最大值
	fmt.Println(bst.FindMax())
	// 查找树中值为3的节点
	fmt.Println(bst.Find(3))
	// 查找树中值为6的节点
	fmt.Println(bst.Find(6))

	// 递归先序遍历
	bst.PreOrderTraversalRecursion()
	fmt.Println()
	// 递归中序遍历
	bst.MidOrderTraversalRecursion()
	fmt.Println()
	// 递归后序遍历
	bst.PostOrderTraversalRecursion()
	fmt.Println()
	bst.PreOrderTraversalRecursion()
	fmt.Println()

	bst.MakeEmpty()
	fmt.Println(bst.FindMin())
	fmt.Println(bst.FindMax())
}
