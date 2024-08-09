package main

import (
	"fmt"

	"github.com/caleb-mwasikira/intro_dsa/algos"
	"github.com/caleb-mwasikira/intro_dsa/utils"
)

func main() {
	arr := utils.GenRandomArray(10, 0, 100)

	fmt.Printf("original array: %#v\n", arr)

	sortedArr := algos.InsertionSort(arr)
	fmt.Printf("insertion sort: %#v\n", sortedArr)

	sortedArr = algos.BubbleSort(arr)
	fmt.Printf("bubble sort: %#v\n", sortedArr)
}
