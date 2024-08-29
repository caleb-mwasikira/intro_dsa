package tests

import (
	"fmt"
	"slices"
	"testing"

	"github.com/caleb-mwasikira/intro_dsa/ds"
	"github.com/caleb-mwasikira/intro_dsa/utils"
)

var (
	size      int   = 5
	testArray []int = utils.GenRandomArray(size, 0, 100)
)

func TestNewStack(t *testing.T) {
	testNewStack_WithInvalidSize_Errors(t, 0)
	testNewStack_WithInvalidSize_Errors(t, -1)
}

func TestStackPush(t *testing.T) {
	stack, _ := ds.NewStack[int](size)
	numItemsAdded, err := stack.Push(testArray...)
	if err != nil {
		t.Error("stack push threw error when it shouldnt")
	}

	if numItemsAdded != len(testArray) {
		t.Errorf("numItemsAdded= %v to stack is not congruent to len(testArray)= %v\n", numItemsAdded, len(testArray))
	}

	// test stack overflow
	_, err = stack.Push(testArray...)
	if err == nil {
		t.Error("stack.Push() does not return (stackoverflow) error when it should")
	}
}

func TestStackPop(t *testing.T) {
	stack, _ := ds.NewStack(size, testArray...)

	// pop last item from stack
	item, err := stack.Pop()
	if err != nil {
		t.Error("stack.Pop() returned error when it shouldn't")
	}

	if item != testArray[len(testArray)-1] {
		t.Error("expected stack.Pop() to return last item in testArray but it didn't")
	}

	// pop all items in stack
	stackItems := stack.GetItems()
	for range stackItems {
		stack.Pop()
	}

	// test stack undeflow
	_, err = stack.Pop()
	if err == nil {
		t.Error("expected stack.Pop() to return error in stackunderflow scenario")
	}
}

func TestStackGetItems(t *testing.T) {
	myTestArray := testArray
	stack, _ := ds.NewStack[int](size, myTestArray...)

	items := stack.GetItems()
	slices.Reverse(myTestArray)

	if !slices.Equal(items, myTestArray) {
		t.Errorf("expected stack.GetItems() to return %#v but got %#v instead", myTestArray, items)
	}
}

func testNewStack_WithInvalidSize_Errors(t *testing.T, size int) {
	testName := fmt.Sprintf("NewStack(%v)", size)

	t.Run(testName, func(t *testing.T) {
		// test if NewStack() return an error with zero
		_, err := ds.NewStack[int](size)
		if err == nil {
			t.Errorf("%v does not return error when it should\n", testName)
		}

		fmt.Printf("    %v returns error\n", testName)
	})
	println()
}
