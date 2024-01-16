package main

func binarySearch(arr []int, x int, left int, right int) int {
	if right <= left {
		// промежуток пуст
		return -1
	}
	// промежуток не пуст
	mid := (left + right) / 2
	if arr[mid] == x {
		// центральный элемент — искомый
		return mid
	} else if x < arr[mid] {
		// искомый элемент меньше центрального значит следует искать в левой половине
		return binarySearch(arr, x, left, mid)
	} else {
		// иначе следует искать в правой половине
		return binarySearch(arr, x, mid+1, right)
	}
}

func brokenSearch(arr []int, k int) int {
	return binarySearch(arr, k, 0, len(arr))
}
