package main

import (
	"fmt"

	"github.com/shunhui19/dataStructureAlgorithm"
)

func main() {
	ll := dataStructureAlgorithm.NewLinkedList()

	fmt.Println("**************** insert *****************")
	// 插入节点
	ll.Insert("header", "china")
	ll.Insert("china", "usa")
	ll.Insert("usa", "india")
	ll.Insert("china", "england")
	// 遍历节点
	ll.Range(func(value interface{}) bool {
		fmt.Printf("%v->", value)
		return true
	})
	fmt.Println()

	fmt.Println("**************** delete *****************")
	// 删除节点
	ll.Delete("china")
	ll.Range(func(value interface{}) bool {
		fmt.Printf("%v->", value)
		return true
	})
	fmt.Println()

	fmt.Println("**************** find *****************")
	fmt.Println(ll.Find("usa").Value)

	fmt.Println("**************** makeEmpty *****************")
	ll.MakeEmpty()
	ll.Range(func(value interface{}) bool {
		fmt.Printf("%v->", value)
		return true
	})
	fmt.Println()
}
