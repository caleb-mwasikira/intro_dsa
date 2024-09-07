package ds

import "fmt"

type BinaryTree struct {
	Root *TreeNode
}

func BTInsert(root *TreeNode, items ...int) (*TreeNode, error) {
	if len(items) == 0 {
		return nil, fmt.Errorf("cannot create new node; empty array")
	}

	if root == nil {
		item := items[0]
		items = items[1:]
		root = &TreeNode{Data: item}
	}

	queue, err := NewCircularQueue[*TreeNode](len(items))
	if err != nil {
		return nil, err
	}
	err = queue.Push(root)
	if err != nil {
		return nil, err
	}

	for !queue.IsEmpty() {
		currentNode, err := queue.Front()
		if err != nil {
			return nil, err
		}

		if currentNode.Left == nil {
			if len(items) == 0 {
				break
			}

			item := items[0]
			items = items[1:]
			newNode := TreeNode{Data: item}

			currentNode.Left = &newNode
			queue.Push(&newNode)
		}

		if currentNode.Right == nil {
			if len(items) == 0 {
				break
			}

			item := items[0]
			items = items[1:]
			newNode := TreeNode{Data: item}

			currentNode.Right = &newNode
			queue.Push(&newNode)
		}
	}
	return root, nil
}

func NewBinaryTree(items ...int) (*BinaryTree, error) {
	rootNode, err := BTInsert(nil, items...)
	if err != nil {
		return nil, err
	}

	return &BinaryTree{
		Root: rootNode,
	}, nil
}

func TreeIncludes(root *TreeNode, item int) bool {
	if root == nil {
		return false
	}

	if root.Data == item {
		return true
	}

	return TreeIncludes(root.Left, item) || TreeIncludes(root.Right, item)
}

func TreeSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Data + TreeSum(root.Left) + TreeSum(root.Right)
}

func TreeMin(root *TreeNode, min int) int {
	if root == nil {
		return min
	}

	if root.Data < min {
		return root.Data
	}

	minLeft := TreeMin(root.Left, min)
	minRight := TreeMin(root.Right, min)

	if minLeft < minRight {
		return minLeft
	} else {
		return minRight
	}
}

func TreeMax(root *TreeNode, max int) int {
	if root == nil {
		return max
	}

	if root.Data > max {
		return root.Data
	}

	maxLeft := TreeMax(root.Left, max)
	maxRight := TreeMax(root.Right, max)

	if maxLeft > maxRight {
		return maxLeft
	} else {
		return maxRight
	}
}
