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
