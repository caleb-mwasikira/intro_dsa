package main

import (
	"fmt"

	"github.com/caleb-mwasikira/intro_dsa/alg"
	"github.com/caleb-mwasikira/intro_dsa/utils"
)

func main() {
	items := utils.GenRandomArray(10, 0, 100)
	fmt.Printf("before sort: %#v\n", items)

	sorted := alg.SelectionSort(items)
	fmt.Printf("sorted array: %#v\n", sorted)

	fmt.Printf("after sort: %#v\n", items)
}
