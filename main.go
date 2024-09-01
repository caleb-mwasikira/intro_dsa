package main

import (
	"fmt"

	"github.com/caleb-mwasikira/intro_dsa/ds"
	"github.com/caleb-mwasikira/intro_dsa/utils"
)

func main() {
	items := utils.GenRandomArray(10, 0, 100)

	ll := ds.NewLinkedList(items...)

	fmt.Printf("%#v\n", items)
	fmt.Println("linked list size: ", ll.Size())
	fmt.Printf("%#v\n", ll.GetItems())
}
