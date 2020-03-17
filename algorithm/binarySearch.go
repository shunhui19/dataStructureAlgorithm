package algorithm

// BinarySearch 二分查找法.
// 查找指定数字是否在数组中.
func BinarySearch(data []int, find int) bool {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := (low + high) >> 1
		if find == data[mid] {
			return true
		}
		if find > data[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}

// BinarySearchFirst 二分查找第一个出现值的位置.
// data中存储多个重复的值
func BinarySearchFirst(data []int, find int) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if find < data[mid] {
			high = mid - 1
		} else if find > data[mid] {
			low = mid + 1
			// 下面是等于指定元素的值情况:
			// 如果数组第一个元素等于指定值，肯定是第一个
			// 如果当前位置mid的前一个元素等于指定值，那么肯定在[low,mid-1]区间中
		} else {
			if mid == 0 || data[mid-1] != find {
				return mid
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

// BinarySearchFirstBig 查找第一个大于等于指定值的元素
func BinarySearchFirstGreater(data []int, find int) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if data[mid] >= find {
			// 如果是第一个元素，或者mid前一个元素小于指定值
			if mid == 0 || data[mid-1] < find {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

// BinarySearchLastLess 查找最后一个小于等于指定元素值
func BinarySearchLastLess(data []int, find int) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if find >= data[mid] {
			// 如果是最后一个元素，或者mid后一个元素大于指定值
			if mid == len(data)-1 || data[mid+1] > find {
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

// BinarySearchLast 查找指定元素最后一个
func BinarySearchLast(data []int, find int) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if find < data[mid] {
			high = mid - 1
		} else if find > data[mid] {
			low = mid + 1
		} else {
			if mid == len(data)-1 || data[mid+1] != find {
				return mid
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}
