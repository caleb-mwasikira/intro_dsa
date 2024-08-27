package main

import (
	"fmt"
	"log"

	"github.com/caleb-mwasikira/intro_dsa/ds"
	"github.com/caleb-mwasikira/intro_dsa/utils"
)

func main() {
	size := 4
	arr := utils.GenRandomArray(size, 0, 100)

	fmt.Printf("original array: %#v\n", arr)

	// add half of items now
	q, err := ds.NewCircularQueue(size, arr[0:len(arr)/2]...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", q.GetItems())

	// and half of items later
	err = q.Enqueue(arr[len(arr)/2:]...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", q.GetItems())

	// consume items on the queue; we shd see FIFO rule apply here
	for {
		item, err := q.Dequeue()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(item)
	}
}
