package alg

/*
selection sorts an array in place
*/
func SelectionSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		smallestIndex := i

		for j := i; j < n; j++ {
			if arr[j] < arr[smallestIndex] {
				smallestIndex = j
			}
		}

		// swap elements
		arr[i], arr[smallestIndex] = arr[smallestIndex], arr[i]
	}
	return arr
}
