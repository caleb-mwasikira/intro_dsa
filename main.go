package main

import (
	"fmt"
	"math/rand"

	"github.com/caleb-mwasikira/intro_dsa/alg"
)

func main() {
	items := []int{1, 2, 3, 4, 5, 6}
	expectedIndex := rand.Intn(len(items))
	item := items[expectedIndex]

	fmt.Printf("%#v\n", items)

	index := alg.BinarySearch(items, item)
	fmt.Printf("binarySearch(%v)\n", item)
	fmt.Println("expected index: ", expectedIndex)
	fmt.Println("index: ", index)

	fmt.Println()
	index = alg.RecursiveBinarySearch(items, item, 0, len(items)-1)
	fmt.Printf("recursiveBinarySearch(%v)\n", item)
	fmt.Println("expected index: ", expectedIndex)
	fmt.Println("index: ", index)
}
