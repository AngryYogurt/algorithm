package quick_sort

/*
	快速排序
*/

func QuickSort3(src []int) {
	_quickSort3(src, 0, len(src)-1)
}

func _quickSort3(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition3(arr, left, right)
		_quickSort3(arr, left, partitionIndex-1)
		_quickSort3(arr, partitionIndex+1, right)
	}
	return arr
}

func partition3(arr []int, left, right int) int {
	pivot := arr[left]
	for left < right {
		for left < right && arr[right] > pivot {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= pivot {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = pivot
	return left
}
