package ds

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func BSTInsert(root *TreeNode, item int) *TreeNode {
	if root == nil {
		newNode := TreeNode{Data: item}
		return &newNode
	}

	if item > root.Data {
		// recursively insert on right subtree
		root.Right = BSTInsert(root.Right, item)
	} else if item < root.Data {
		// recursively insert on left subtree
		root.Left = BSTInsert(root.Left, item)
	} else {
		// item == root.Data; do nothing
	}

	return root
}

type BST struct {
	Root *TreeNode
}

func NewBST(items ...int) BST {
	bst := BST{}

	for index, item := range items {
		if index == 0 {
			bst.Root = &TreeNode{Data: item}
			continue
		}

		BSTInsert(bst.Root, item)
	}
	return bst
}

func BSTIncludes(root *TreeNode, item int) bool {
	if root == nil {
		return false
	}

	if item < root.Data {
		return BSTIncludes(root.Left, item)
	} else if item > root.Data {
		return BSTIncludes(root.Right, item)
	} else {
		// item == root.Data
		return true
	}
}

func BSTSearch(root *TreeNode, item int) *TreeNode {
	if root == nil {
		return nil
	}

	if item < root.Data {
		return BSTSearch(root.Left, item)
	} else if item > root.Data {
		return BSTSearch(root.Right, item)
	} else {
		// item == root.Data
		return root
	}
}

/* Note: binary search trees do NOT guarentee any form of node balancing.
In the worst case scenario, a binary search tree could be skewed linearly
to the left or right like so;

input: [1,2,3,4]	     or          [4,3,2,1]
			1							4
		     \						   /
			  2						  3
			   \					 /
			    3					2
				 \				   /
				  4				  1

Having a BST skewed like this reduces the performance of BST search algorithm
from O(log n) O(n), as our BST 'tree' is now just a fancy linked-list.

TODO: to prevent this learn about self-balancing trees
*/

func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	leftValues := PreorderTraversal(root.Left)
	rightValues := PreorderTraversal(root.Right)

	// [ root, leftValues..., rightValues... ]
	return append(append([]int{root.Data}, leftValues...), rightValues...)
}

func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	leftValues := InorderTraversal(root.Left)
	rightValues := InorderTraversal(root.Right)

	// [ leftValues..., root, rightValues... ]
	return append(append(leftValues, root.Data), rightValues...)
}

func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	leftValues := PostorderTraversal(root.Left)
	rightValues := PostorderTraversal(root.Right)

	// [leftValues..., rightValues..., root ]
	return append(append(leftValues, rightValues...), root.Data)
}
