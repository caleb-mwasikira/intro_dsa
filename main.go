package main

import (
	"fmt"
	"log"
	"slices"

	"github.com/caleb-mwasikira/intro_dsa/ds"
	"github.com/caleb-mwasikira/intro_dsa/utils"
)

func main() {
	size := 5
	arr := utils.GenRandomArray(size, 0, 100)

	fmt.Printf("original array: %#v\n", arr)

	// add half of items now
	stack, err := ds.NewStack(size, arr[0:size/2]...)
	if err != nil {
		log.Fatalf("error creating new stack: %v\n", err)
	}

	// and half of items later
	_, err = stack.Push(arr[size/2:]...)
	if err != nil {
		log.Fatalf("error adding item to stack: %v\n", err)
	}

	// check if items added are on the stack
	items := stack.GetItems()
	if !slices.Equal(arr, items) {
		log.Printf("array %v not equal to stack items %v\n", arr, items)
	}

}
