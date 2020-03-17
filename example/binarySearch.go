package main

import (
	"fmt"

	"github.com/shunhui19/dataStructureAlgorithm/algorithm"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(algorithm.BinarySearch(arr, 9))
	//fmt.Println(algorithm.BinarySearch(arr, 19))

	// 查找第一个值等于给定值的元素
	arr = []int{1, 2, 5, 6, 6, 6, 7, 8, 9}
	fmt.Println(algorithm.BinarySearchFirst(arr, 6))

	// 查找最后一个值等于给定值的元素
	arr = []int{1, 2, 5, 6, 6, 6, 6, 7, 8, 9}
	fmt.Println(algorithm.BinarySearchLast(arr, 6))

	// 查找第一个大于等于等于给定值的元素
	arr = []int{1, 2, 4, 6, 6, 6, 7, 8, 9}
	fmt.Println(algorithm.BinarySearchFirstGreater(arr, 6))

	// 查找最后一个小于等于等于给定值的元素
	arr = []int{1, 2, 4, 6, 6, 6, 7, 8, 9}
	fmt.Println(algorithm.BinarySearchLastLess(arr, 6))
}
