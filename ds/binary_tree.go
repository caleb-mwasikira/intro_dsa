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
