package main

import (
	"fmt"

	"github.com/shunhui19/dataStructureAlgorithm"
)

func main() {
	dl := dataStructureAlgorithm.NewDoubleLinkedList()

	fmt.Println("**************** insert *****************")
	// 插入节点
	dl.Insert("header", "china")
	dl.Insert("china", "usa")
	dl.Insert("usa", "india")
	dl.Insert("china", "england")

	dl.Range(func(value interface{}) bool {
		fmt.Printf("%v->", value)
		return true
	})

	fmt.Println("**************** delete *****************")
	// 删除节点
	dl.Delete("china")
	dl.Range(func(value interface{}) bool {
		fmt.Printf("%v->", value)
		return true
	})
	fmt.Println()
}
