package algorithm

import (
	"fmt"
	"math"
)

// bubbleSort 冒泡排序.
// 原理：比较相邻的两个数，直到最后一个数，这样一趟就完成了，分为N趟
// 第一趟：0~N
// 第二趟：0~N-1
// 第三趟：0~N-2
// ...
// 时间复杂度为：O(n*n)
func BubbleSort(data []int) []int {
	length := len(data)

	for i := length; i >= 1; i-- {
		for j := 0; j < i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}

	return data
}

// InsertSort 插入排序.
// 原理：每一步将一个待排序的数据插入到前面已经排好序的有序序列中, 直到所有元素插入完毕为止.
// 流程：
// 1. 将第一个元素看作是一个有序序列，把第二个元素到最后一个元素当成是未排序序列
// 2. 从头到尾依次扫描未排序序列，将扫描到的每个元素插入到有序序列的适当位置，
// 如果待插入的元素与有序序列中的某个元素相等，则将待插入的元素插入到相等元素后面
// 时间复杂度为：O(n*n)
func InsertSort(data []int) []int {
	length := len(data)
	for i := 0; i < length-1; i++ {
		// 第i个的下一个元素与前面 0~i 有序序列个元素比较，并插入到适当的位置
		for j := i + 1; j > 0; j-- {
			if data[j] < data[j-1] {
				data[j-1], data[j] = data[j], data[j-1]
			}
		}
	}

	return data
}

// ShellSort 希尔排序.
// 原理：是递减增量排序算法，是插入排序的一种更高效的改进版本.
// 流程：
// 1. 取增量，一般为数组长度的1/2
// 2. 按增量取得子序列，对子序列进行排序
// 3. 将增量递减，重复1，2步骤
// 4. 直至增量均为0
func shellSort(data []int) []int {
	length := len(data)
	for gap := int(math.Floor(float64(length / 2))); gap > 0; gap = int(math.Floor(float64(gap / 2))) {
		for i := 0; i+gap < length; i++ {
			if data[i] > data[i+gap] {
				data[i], data[i+gap] = data[i+gap], data[i]
			}
		}
		fmt.Println(data)
		//for i := gap; i < length-1; i++ {
		//	for k := i - gap; k < gap; k++ {
		//		if data[k] > data[k+gap] {
		//			data[k], data[k+gap] = data[k+gap], data[k]
		//		}
		//	}
		//}
	}

	return data
}

// QuickSort 快速排序.
// 原理：采用分治思想实现, 步骤如下：
// 1. 选择基准值
// 2. 将数组分为两个子数组， 小于基准值的元素和大于基准值的元素
// 3. 对这两个子数组进行快速排序
func QuickSort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	base := data[0]
	var result, left, right []int
	for i := 1; i < len(data); i++ {
		if data[i] < base {
			left = append(left, data[i])
		} else {
			right = append(right, data[i])
		}
	}

	// 继续对子数组快速排序
	left = QuickSort(left)
	right = QuickSort(right)

	result = append(left, base)
	return append(result, right...)
}
