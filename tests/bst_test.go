package tests

import (
	"slices"
	"testing"

	"github.com/caleb-mwasikira/intro_dsa/ds"
)

func TestBSTInsert(t *testing.T) {
	items := []int{11, 6, 19}

	/*
		expected binary search tree

				11
			   /  \
			  6   19
	*/

	var rootNode ds.TreeNode
	for index, item := range items {
		if index == 0 {
			rootNode = ds.TreeNode{Data: item}
			continue
		}

		ds.BSTInsert(&rootNode, item)
	}

	leftSubtree := rootNode.Left
	rightSubtree := rootNode.Right

	expectedLeftNodeData := 6
	if leftSubtree.Data != expectedLeftNodeData {
		t.Errorf("expected %v as left node in BST but got %v", expectedLeftNodeData, leftSubtree.Data)
	}

	expectedRightNodeData := 19
	if rightSubtree.Data != expectedRightNodeData {
		t.Errorf("expected %v as right node in BST but got %v", expectedRightNodeData, rightSubtree.Data)
	}

}

func TestNewBST(t *testing.T) {
	items := []int{11, 6, 8, 19}

	bst := ds.NewBST(items...)

	/*
		expected binary search tree

						11
					   /  \
			          6   19
				       \
					   8
	*/

	// check if root is set up correctly
	if bst.Root == nil || bst.Root.Left == nil || bst.Root.Right == nil {
		t.Errorf("binary search tree not formed correctly; either due to missing root node or one or both children")
	}

	// best way to test the presence of all nodes is
	// to do a traversal on the binary tree
	expectedInorderItems := []int{6, 8, 11, 19}
	inorderItems := ds.InorderTraversal(bst.Root)
	if !slices.Equal(inorderItems, expectedInorderItems) {
		t.Errorf("expected items %#v after inorder traversal on binary search tree but got %#v", expectedInorderItems, inorderItems)
	}
}

func TestPreorderTraversal(t *testing.T) {
	items := []int{15, 68, 78, 86, 38, 58, 4, 37, 35, 87}
	expectedPreorder := []int{15, 4, 68, 38, 37, 35, 58, 78, 86, 87}

	bst := ds.NewBST(items...)
	preorder := ds.PreorderTraversal(bst.Root)

	if !slices.Equal(preorder, expectedPreorder) {
		t.Errorf("expected preorder %#v but got %#v", expectedPreorder, preorder)
	}
}

func TestInorderTraversal(t *testing.T) {
	items := []int{15, 68, 78, 86, 38, 58, 4, 37, 35, 87}
	expectedInorder := []int{4, 15, 35, 37, 38, 58, 68, 78, 86, 87}

	bst := ds.NewBST(items...)
	preorder := ds.InorderTraversal(bst.Root)

	if !slices.Equal(preorder, expectedInorder) {
		t.Errorf("expected inorder %#v but got %#v", expectedInorder, preorder)
	}
}

func TestPostorderTraversal(t *testing.T) {
	items := []int{15, 68, 78, 86, 38, 58, 4, 37, 35, 87}
	expectedPostorder := []int{4, 35, 37, 58, 38, 87, 86, 78, 68, 15}

	bst := ds.NewBST(items...)
	preorder := ds.PostorderTraversal(bst.Root)

	if !slices.Equal(preorder, expectedPostorder) {
		t.Errorf("expected postorder %#v but got %#v", expectedPostorder, preorder)
	}
}
