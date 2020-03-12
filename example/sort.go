package main

import (
	"fmt"

	"github.com/shunhui19/dataStructureAlgorithm/algorithm"
)

func main() {
	arr := []int{30, 3, 12, 9, 29, 100, 66, 22, 15}
	fmt.Println(arr)

	// 冒泡排序
	arr = []int{30, 3, 12, 9, 29, 100, 66, 22, 15}
	fmt.Printf("冒泡排序: %v\n", algorithm.BubbleSort(arr))

	// 插入排序
	arr = []int{30, 3, 12, 9, 29, 100, 66, 22, 15}
	fmt.Printf("插入排序: %v\n", algorithm.InsertSort(arr))

	// 快速排序
	arr = []int{8, 13, 6, 5, 4, 2, 11, 9, 7}
	fmt.Println(arr)
	fmt.Printf("快速排序: %v\n", algorithm.QuickSort(arr))

}
