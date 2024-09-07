package tests

import (
	"slices"
	"testing"

	"github.com/caleb-mwasikira/intro_dsa/ds"
)

func TestBTInsert(t *testing.T) {
	items := []int{17, 43, 49, 31, 6, 11, 8}

	/*
		expected binary tree
				17
			   /  \
			 43    49
			/ \    / \
		   31  6  11  8
	*/

	root, err := ds.BTInsert(nil, items...)
	if err != nil {
		t.Errorf("unexpected error creating binary tree with valid arguments: %v", err)
		return
	}

	// best way to test the presence of all nodes is
	// to do a traversal on the binary tree
	expectedInorderItems := []int{31, 43, 6, 17, 11, 49, 8}
	inorderItems := ds.InorderTraversal(root)

	if !slices.Equal(inorderItems, expectedInorderItems) {
		t.Errorf("expected inorder traversal to return nodes %#v in binary tree but got %#v", expectedInorderItems, inorderItems)
	}
}

func TestNewBinaryTree(t *testing.T) {
	items := []int{17, 43, 49, 31, 6, 11, 8}

	/*
		expected binary tree
				17
			   /  \
			 43    49
			/ \    / \
		   31  6  11  8
	*/

	bt, err := ds.NewBinaryTree(items...)
	if err != nil {
		t.Errorf("unexpected error creating binary tree with valid arguments: %v", err)
		return
	}

	// best way to test the presence of all nodes is
	// to do a traversal on the binary tree
	expectedInorderItems := []int{31, 43, 6, 17, 11, 49, 8}
	inorderItems := ds.InorderTraversal(bt.Root)

	if !slices.Equal(inorderItems, expectedInorderItems) {
		t.Errorf("expected inorder traversal to return nodes %#v in binary tree but got %#v", expectedInorderItems, inorderItems)
	}
}
