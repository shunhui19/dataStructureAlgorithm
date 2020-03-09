package main

import (
	"fmt"

	"github.com/shunhui19/dataStructureAlgorithm"
)

func main() {
	ht := dataStructureAlgorithm.NewHashTable(15)

	// 插入元素
	ht.Insert("china")
	ht.Insert("goes")
	ht.Insert("frank")
	ht.Insert("jak")
	fmt.Println(ht)

	// 查找元素
	fmt.Println(ht.Get("china"))

	// 删除元素
	ht.Delete("china")
	fmt.Println(ht.Get("china"))
}
