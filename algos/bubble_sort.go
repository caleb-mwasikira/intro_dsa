package algos

func BubbleSort(items []int) []int {
	if len(items) <= 1 {
		return items
	}

	for index, item := range items {
		if index == 0 {
			continue
		}

		for j := index; j > 0; j-- {
			if item < items[j-1] {
				// swap the two elements
				items[j] = items[j-1]
				items[j-1] = item
			}
		}
	}
	return items
}
