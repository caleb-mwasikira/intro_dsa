package tests

import (
	"testing"

	"github.com/caleb-mwasikira/intro_dsa/alg"
)

func TestBinarySearch(t *testing.T) {
	// binary search requires a sorted array to work
	items := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// test for elements in array
	for expectedIndex, item := range items {
		index := alg.BinarySearch(items, item)

		if index != expectedIndex {
			t.Errorf("expected index=%v from binarySearch(%v) but got index=%v", expectedIndex, item, index)
		}
	}

	// test for elements NOT in array
	item := 69
	expectedIndex := -1
	index := alg.BinarySearch(items, item)

	if index != expectedIndex {
		t.Errorf("expected index=%v from binarySearch(%v) but got index=%v", expectedIndex, item, index)
	}
}

func TestRecursiveBinarySearch(t *testing.T) {
	// binary search requires a sorted array to work
	items := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// test for elements in array
	for expectedIndex, item := range items {
		index := alg.RecursiveBinarySearch(items, item, 0, len(items)-1)

		if index != expectedIndex {
			t.Errorf("expected index=%v from recursiveBinarySearch(%v) but got index=%v", expectedIndex, item, index)
		}
	}

	// test for elements NOT in array
	item := 69
	expectedIndex := -1
	index := alg.RecursiveBinarySearch(items, item, 0, len(items)-1)

	if index != expectedIndex {
		t.Errorf("expected index=%v from binarySearch(%v) but got index=%v", expectedIndex, item, index)
	}
}
