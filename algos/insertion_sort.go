package algos

func InsertionSort(items []int) []int {
	if len(items) <= 1 {
		return items
	}

	for index := 0; index < len(items); index++ {
		key := items[index]
		if index == 0 {
			// assume first element as already sorted
			continue
		}

		for j := index; j >= 0; j-- {
			if key < items[j] {
				// shift j to the right by 1
				items[j+1] = items[j]

				// place key where j was
				items[j] = key
			}
		}
	}
	return items
}
