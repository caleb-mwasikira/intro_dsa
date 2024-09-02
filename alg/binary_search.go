package alg

func BinarySearch(items []int, item int) int {
	low := 0
	high := len(items) - 1

	for low <= high {
		mid := (low + high) / 2

		if items[mid] == item {
			return mid
		} else if item > items[mid] {
			// item is on right side
			low = mid + 1
		} else {
			// item is on left side
			high = mid - 1
		}
	}
	return -1
}

func RecursiveBinarySearch(items []int, item, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2

	if items[mid] == item {
		return mid
	} else if item > items[mid] {
		// item is on right side
		return RecursiveBinarySearch(items, item, mid+1, high)
	} else {
		// item is on left side
		return RecursiveBinarySearch(items, item, low, mid-1)
	}
}
