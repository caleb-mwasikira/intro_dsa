package utils

import (
	"crypto/rand"
	"log"
	"math/big"
)

func GenRandomArray(size, min, max int) []int {
	if size <= 0 {
		log.Fatal("invalid length for random array generator")
	}

	if min > max {
		log.Fatal("min cannot be greater than max")
	}

	items := make([]int, size)

	for i := 0; i < size; i++ {
		randNumber, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
		if err != nil {
			return []int{}
		}
		items[i] = int(randNumber.Int64()) + min
	}
	return items
}
