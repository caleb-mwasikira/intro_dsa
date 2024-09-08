package tests

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/caleb-mwasikira/intro_dsa/ds"
)

var (
	expectedBSTPreorder  = []int{17, 6, 11, 8, 43, 31, 49}
	expectedBSTInorder   = []int{6, 8, 11, 17, 31, 43, 49}
	expectedBSTPostorder = []int{8, 11, 6, 31, 49, 43, 17}
)

func init() {
	/*
		expected binary search tree
				17
			   /  \
			 6     43
			  \    / \
		      11  31  49
			  /
			 8
	*/
}

func TestBSTInsert(t *testing.T) {
	var rootNode ds.TreeNode
	for index, item := range items {
		if index == 0 {
			rootNode = ds.TreeNode{Data: item}
			continue
		}

		ds.BSTInsert(&rootNode, item)
	}

	// best way to test the presence of all nodes is
	// to do a traversal on the binary tree
	inorderItems := ds.InorderTraversal(&rootNode)
	if !slices.Equal(inorderItems, expectedBSTInorder) {
		t.Errorf("expected items %#v after inorder traversal on binary search tree but got %#v", expectedBSTInorder, inorderItems)
	}
}

func TestNewBST(t *testing.T) {
	bst := ds.NewBST(items...)

	// check if root is set up correctly
	if bst.Root == nil || bst.Root.Left == nil || bst.Root.Right == nil {
		t.Errorf("binary search tree not formed correctly; either due to missing root node or one or both children")
	}

	// best way to test the presence of all nodes is
	// to do a traversal on the binary tree
	inorderItems := ds.InorderTraversal(bst.Root)
	if !slices.Equal(inorderItems, expectedBSTInorder) {
		t.Errorf("expected items %#v after inorder traversal on binary search tree but got %#v", expectedBSTInorder, inorderItems)
	}
}

func TestPreorderTraversal(t *testing.T) {
	bst := ds.NewBST(items...)
	preorder := ds.PreorderTraversal(bst.Root)

	if !slices.Equal(preorder, expectedBSTPreorder) {
		t.Errorf("expected preorder %#v but got %#v", expectedBSTPreorder, preorder)
	}
}

func TestInorderTraversal(t *testing.T) {
	bst := ds.NewBST(items...)
	preorder := ds.InorderTraversal(bst.Root)

	if !slices.Equal(preorder, expectedBSTInorder) {
		t.Errorf("expected inorder %#v but got %#v", expectedBSTInorder, preorder)
	}
}

func TestPostorderTraversal(t *testing.T) {
	bst := ds.NewBST(items...)
	preorder := ds.PostorderTraversal(bst.Root)

	if !slices.Equal(preorder, expectedBSTPostorder) {
		t.Errorf("expected postorder %#v but got %#v", expectedBSTPostorder, preorder)
	}
}

func TestBSTIncludes(t *testing.T) {
	// pick random item from items array
	randIndex := rand.Intn(len(items))
	item := items[randIndex]

	bst := ds.NewBST(items...)

	// test bstIncludes() on item in items array
	includes := ds.BSTIncludes(bst.Root, item)
	if !includes {
		t.Errorf("expected item %v to be in binary search tree but bstIncludes() returned false", item)
	}

	// test bstIncludes() on item NOT in items array
	item = 69
	includes = ds.BSTIncludes(bst.Root, item)
	if includes {
		t.Errorf("expected item %v to NOT be in binary search tree but bstIncludes() returned true", item)
	}
}

func TestBSTSearch(t *testing.T) {
	// pick random item from items array
	randIndex := rand.Intn(len(items))
	item := items[randIndex]

	bst := ds.NewBST(items...)

	// test bstSearch() on item in items array
	node := ds.BSTSearch(bst.Root, item)
	if node == nil {
		t.Errorf("expected item %v to be in binary search tree but bstSearch() returned nil", item)
	}

	// test bstSearch() on item NOT in items array
	item = 69
	node = ds.BSTSearch(bst.Root, item)
	if node != nil {
		t.Errorf("expected item %v to NOT be in binary search tree but bstSearch() returned non-nil item", item)
	}
}
