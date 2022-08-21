package quick_sort

/*
	快速排序
*/

func QuickSort2(src []int) {
	_quickSort2(src, 0, len(src)-1)
}

func _quickSort2(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition2(arr, left, right)
		_quickSort2(arr, left, partitionIndex-1)
		_quickSort2(arr, partitionIndex+1, right)
	}
	return arr
}

func partition2(arr []int, left, right int) int {
	start, _ := left, right
	pivot := arr[left]
	for {
		for arr[left] <= pivot && left < right {
			left++
		}
		for arr[right] > pivot && right > left {
			right--
		}
		if left >= right {
			break
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	if arr[left] > pivot {
		left--
	}
	arr[start], arr[left] = arr[left], arr[start]
	return left
}
