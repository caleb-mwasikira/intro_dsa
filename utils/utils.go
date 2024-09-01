package utils

import (
	"fmt"
	"math/rand"
	"time"
)

/*
generates a random array of size n with items between
min and max
*/
func GenRandomArray(length int, min, max int) []int {
	if length <= 0 || max <= 0 {
		return []int{}
	}

	// if min > max, min is the new max
	if min > max {
		tmp := min
		min = max
		max = tmp
	}

	// seed random number generator with current unix timestamp
	// to prevent generation of the same random numbers every time
	// this function is run
	rand.Seed(time.Now().UnixNano())

	array := make([]int, length)

	for i := 0; i < length; i++ {
		randomNumber := rand.Intn(max-min+1) + min
		array[i] = randomNumber
	}
	return array
}

/*
returns the index of a random element in an array
*/
func RandomIndexChoice(array []int) (int, error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("empty array")
	}
	randomIndex := rand.Intn(len(array))
	return randomIndex, nil
}
