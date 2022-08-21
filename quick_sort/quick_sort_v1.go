package quick_sort

/*
	快速排序
*/

func QuickSort1(src []int) {
	_quickSort1(src, 0, len(src)-1)
}

func _quickSort1(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition1(arr, left, right)
		_quickSort1(arr, left, partitionIndex-1)
		_quickSort1(arr, partitionIndex+1, right)
	}
	return arr
}

func partition1(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index += 1
		}
	}
	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	return index - 1
}
