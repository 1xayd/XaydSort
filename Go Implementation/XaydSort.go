package main

func XaydSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	xaydSortRecursive(arr, 0, len(arr)-1)
}

func xaydSortRecursive(arr []int, low, high int) {
	if high-low < 40 {
		xaydSortInsert(arr, low, high)
		return
	}

	pivotIdx := xaydPartition(arr, low, high)

	xaydSortRecursive(arr, low, pivotIdx-1)
	xaydSortRecursive(arr, pivotIdx+1, high)
}

func xaydPartition(arr []int, low, high int) int {
	mid := low + ((high - low) >> 1)

	if arr[mid] < arr[low] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	if arr[high] < arr[low] {
		arr[low], arr[high] = arr[high], arr[low]
	}
	if arr[mid] < arr[high] {
		arr[mid], arr[high] = arr[high], arr[mid]
	}

	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func xaydSortInsert(arr []int, low, high int) {
	for i := low + 1; i <= high; i++ {
		key := arr[i]

		pos := exponentialSearch(arr, low, i-1, key)

		if pos != i {
			copy(arr[pos+1:i+1], arr[pos:i])
			arr[pos] = key
		}
	}
}

func exponentialSearch(arr []int, left, right int, key int) int {
	if right < left || key <= arr[left] {
		return left
	}
	if key > arr[right] {
		return right + 1
	}

	bound := 1
	for left+bound <= right && arr[left+bound] < key {
		bound <<= 1
	}

	start := left + (bound >> 1)
	end := right
	if left+bound <= right {
		end = left + bound
	}

	return binarySearchInsertPos(arr, start, end, key)
}

func binarySearchInsertPos(arr []int, left, right int, key int) int {
	for left <= right {
		mid := left + ((right - left) >> 1)

		if arr[mid] < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
