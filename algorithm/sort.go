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
	// 是否有数据交换标志，当没有数据交换时，表示已经完成排好序了，不需要继续冒泡了
	var flag bool

	for i := length; i >= 1; i-- {
		flag = false
		for j := 0; j < i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				flag = true
			}
		}
		if !flag {
			break
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

// SelectSort 选择排序.
// 原理：每一步从一个待排序序列选择一个最小的数据插入到已经排好序的有序序列中，直到所有元素插入完毕.
// 和插入排序类似
func SelectSort(data []int) []int {
	length := len(data)
	// 未排序最小元素下标
	var min int
	for i := 0; i < length-1; i++ {
		min = i
		// 从未排序列表中选择最小元素
		for j := i + 1; j < length; j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
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

// HeapSort 堆排序.
// 堆也是可以理解为特殊的二叉查找树, 存储结构采用数组, 可以非常的节省空间的，
// 不需要左右指针，单纯的通过数组下标就可以找到一个节点的左右子节点和父节点.
// 原理，先构造大顶堆或小顶堆，然后再堆排序.
// 构造堆有两种方式：
// 1.从下往上构造堆
// 2.从上往下构造堆
// 就是顺着节点所在路径，向上或向下，对比，然后交换
func HeapSort(data []int) []int {
	length := len(data)
	if length == 0 {
		return data
	}

	for i := 0; i < length; i++ {
		heapifyUp(data[i:])
	}
	return data
}

// heapify 从下往上把数组构成一个大堆
// 节点i的左子树为2*i+1, 右子树为2*i+2
func heapifyUp(data []int) {
	length := len(data)
	for i := length/2 - 1; i >= 0; i-- {
		maxPos := i
		if 2*i+1 <= length-1 && data[2*i+1] > data[i] {
			maxPos = 2*i + 1
		}
		if 2*i+2 <= length-1 && data[maxPos] < data[2*i+2] {
			maxPos = 2*i + 2
		}
		data[i], data[maxPos] = data[maxPos], data[i]
	}
}

// mergeSort 把原问题拆分为多个子问题
func mergeSort(data []int, left, right int, tmpArr []int) {
	if left < right {
		mid := (left + right) / 2
		mergeSort(data, left, mid, tmpArr)    // 左边序列排序
		mergeSort(data, mid+1, right, tmpArr) // 右边序列排序
		merge(data, left, mid, right, tmpArr) // 合并两个有序序列
	}
}

// MergeSort 归并排序
func MergeSort(data []int) []int {
	tmpArr := make([]int, len(data))
	mergeSort(data, 0, len(data)-1, tmpArr)
	return data
}

// merge 合并子序列并按从小到大顺序放入临时数组tmpArr中.
// left是左子序列的起始下标，mid为左序列的最后一个元素的下标，
// mid+1为右子序列的起始下标，right为右子序列的最后一个元素下标
// 如下所示
// [1, 2, 8, 9]		[3, 5, 10, 15]
// left      mid     mid+1     right
func merge(data []int, left, mid, right int, tmpArr []int) []int {
	lStart := left
	rStart := mid + 1
	t := 0

	// 比较两个左右子序列元素大小，并按从小到大顺序写入tmpArr中
	for lStart <= mid && rStart <= right {
		if data[lStart] < data[rStart] {
			tmpArr[t] = data[lStart]

			lStart++
			t++
		} else {
			tmpArr[t] = data[rStart]
			rStart++
			t++
		}
	}

	// 将左子序列剩余元素拷贝到临时数组tmpArr中
	for lStart <= mid {
		tmpArr[t] = data[lStart]
		lStart++
		t++
	}

	// 将右子序列剩余元素拷贝到临时数组tmpArr中
	for rStart <= right {
		tmpArr[t] = data[rStart]
		rStart++
		t++
	}

	// 把有序的元素拷贝到原数组中去
	t = 0
	for left <= right {
		data[left] = tmpArr[t]
		left++
		t++
	}

	return data
}
