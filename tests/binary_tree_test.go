package tests

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/caleb-mwasikira/intro_dsa/ds"
)

var (
	btItems           = []int{17, 43, 49, 31, 6, 11, 8}
	expectedBTInorder = []int{31, 43, 6, 17, 11, 49, 8}
	bt                *ds.BinaryTree
)

func init() {
	/*
		expected binary tree
				17
			   /  \
			 43    49
			/ \    / \
		   31  6  11  8
	*/

	var err error
	var t *testing.T = &testing.T{}
	bt, err = ds.NewBinaryTree(btItems...)
	if err != nil {
		t.Errorf("unexpected error creating binary tree with valid arguments: %v", err)
		return
	}
}

func TestBTInsert(t *testing.T) {
	root, err := ds.BTInsert(nil, btItems...)
	if err != nil {
		t.Errorf("unexpected error creating binary tree with valid arguments: %v", err)
		return
	}

	// best way to test the presence of all nodes is
	// to do a traversal on the binary tree
	inorderItems := ds.InorderTraversal(root)

	if !slices.Equal(inorderItems, expectedBTInorder) {
		t.Errorf("expected inorder traversal to return nodes %#v in binary tree but got %#v", expectedBTInorder, inorderItems)
	}
}

func TestNewBinaryTree(t *testing.T) {
	// best way to test the presence of all nodes is
	// to do a traversal on the binary tree
	inorderItems := ds.InorderTraversal(bt.Root)

	if !slices.Equal(inorderItems, expectedBTInorder) {
		t.Errorf("expected inorder traversal to return nodes %#v in binary tree but got %#v", expectedBTInorder, inorderItems)
	}
}

func TestBTIncludes(t *testing.T) {
	// pick random item from btItems array
	randIndex := rand.Intn(len(btItems))
	item := btItems[randIndex]

	// test btIncludes() on item in items array
	includes := ds.BTIncludes(bt.Root, item)
	if !includes {
		t.Errorf("expected item %v to be in binary tree but btIncludes() returned false", item)
	}

	// test btIncludes() on item NOT in items array
	item = 69
	includes = ds.BTIncludes(bt.Root, item)
	if includes {
		t.Errorf("expected item %v to NOT be in binary tree but btIncludes() returned true", item)
	}
}

func TestBTSearch(t *testing.T) {
	// pick random item from btItems array
	randIndex := rand.Intn(len(btItems))
	item := btItems[randIndex]

	// test search() on item in btItems array
	node := ds.BTSearch(bt.Root, item)
	if node == nil {
		t.Errorf("expected item %v to be in binary tree but search() returned nil", item)
	}

	// test search() on item NOT in items array
	item = 69
	node = ds.BTSearch(bt.Root, item)
	if node != nil {
		t.Errorf("expected item %v to NOT be in binary tree but search() found returned non-nil item", item)
	}
}

func TestTreeSum(t *testing.T) {
	calcSum := 0
	for _, item := range btItems {
		calcSum += item
	}

	treeSum := ds.TreeSum(bt.Root)
	if treeSum != calcSum {
		t.Errorf("expected treeSum() to return %v but got %v", calcSum, treeSum)
	}
}

func TestTreeMin(t *testing.T) {
	calcMin := btItems[0]
	for _, item := range btItems {
		if item < calcMin {
			calcMin = item
		}
	}

	treeMin := ds.TreeMin(bt.Root, btItems[0])
	if treeMin != calcMin {
		t.Errorf("expected treeMin() to return %v but got %v", calcMin, treeMin)
	}
}

func TestTreeMax(t *testing.T) {
	calcMax := btItems[0]
	for _, item := range btItems {
		if item > calcMax {
			calcMax = item
		}
	}

	treeMax := ds.TreeMax(bt.Root, btItems[0])
	if treeMax != calcMax {
		t.Errorf("expected treeMax() to return %v but got %v", calcMax, treeMax)
	}
}

func TestBreadthFirstTraversal(t *testing.T) {
	expectedOrderOfbtItems := []int{17, 43, 49, 31, 6, 11, 8}
	orderOfbtItems, err := ds.BreadthFirstTraversal(bt.Root)
	if err != nil {
		t.Errorf("unexpected error calling breadthFirstTraversal() algorithm: %v", err)
	}

	if !slices.Equal(orderOfbtItems, expectedOrderOfbtItems) {
		t.Errorf("expected btItems %#v on breadthFirstTraversal() but got %#v", expectedOrderOfbtItems, orderOfbtItems)
	}
}
