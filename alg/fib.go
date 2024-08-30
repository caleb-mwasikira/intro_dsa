package alg

func IterativeFib(n int) []int {
	items := []int{}
	if n < 0 {
		// invalid fib index
		return items
	} else if n == 0 {
		return []int{0}
	} else if n == 1 {
		return []int{0, 1}
	}

	prev2 := int(0)
	prev := int(1)
	result := prev + prev2
	items = append(items, prev2, prev, result)

	for i := int(2); i < n; i++ {
		prev2 = prev
		prev = result
		result = prev + prev2
		items = append(items, result)
	}

	return items
}

func RecursiveFib(n int) []int {
	items := []int{}
	if n < 0 {
		return items
	}

	for i := int(0); i <= n; i++ {
		res := recFib(i)
		items = append(items, res)
	}
	return items
}

func recFib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return recFib(n-1) + recFib(n-2)
}

func CachedRecursiveFib(n int) []int {
	items := []int{}
	cache := map[int]int{}

	if n < 0 {
		return items
	}

	for i := int(0); i <= n; i++ {
		res := cachedRecFib(i, cache)
		items = append(items, res)
	}
	return items
}

func cachedRecFib(n int, cache map[int]int) int {
	if n == 0 || n == 1 {
		return n
	}

	// check if value in cache already
	if val, ok := cache[n]; ok {
		return val
	}

	prev := cachedRecFib(n-1, cache)
	prev2 := cachedRecFib(n-2, cache)

	// place calculated fib values onto cache
	cache[n-1] = prev
	cache[n-2] = prev2
	return prev + prev2
}
