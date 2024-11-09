package ds

import "fmt"

type BinaryTree struct {
	Root *Node
}

func BTInsert(root *Node, items ...int) (*Node, error) {
	if len(items) == 0 {
		return nil, fmt.Errorf("cannot create new node; empty array")
	}

	if root == nil {
		item := items[0]
		items = items[1:]
		root = &Node{Data: item}
	}

	queue, err := NewCircularQueue[*Node](len(items))
	if err != nil {
		return nil, err
	}
	err = queue.AddItems(root)
	if err != nil {
		return nil, err
	}

	for !queue.IsEmpty() {
		currentNode, err := queue.GetFrontItem()
		if err != nil {
			return nil, err
		}

		if currentNode.Left == nil {
			if len(items) == 0 {
				break
			}

			item := items[0]
			items = items[1:]
			newNode := Node{Data: item}

			currentNode.Left = &newNode
			queue.AddItems(&newNode)
		}

		if currentNode.Right == nil {
			if len(items) == 0 {
				break
			}

			item := items[0]
			items = items[1:]
			newNode := Node{Data: item}

			currentNode.Right = &newNode
			queue.AddItems(&newNode)
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

func BTIncludes(root *Node, item int) bool {
	if root == nil {
		return false
	}

	if root.Data == item {
		return true
	}

	return BTIncludes(root.Left, item) || BTIncludes(root.Right, item)
}

func BTSearch(root *Node, item int) *Node {
	if root == nil {
		return nil
	}

	// preorder traversal; root->left->right
	if root.Data == item {
		return root
	}

	leftNode := BTSearch(root.Left, item)
	if leftNode != nil {
		return leftNode
	}

	rightNode := BTSearch(root.Right, item)
	if rightNode != nil {
		return rightNode
	}

	return nil
}

func TreeSum(root *Node) int {
	if root == nil {
		return 0
	}
	return root.Data + TreeSum(root.Left) + TreeSum(root.Right)
}

func TreeMin(root *Node, min int) int {
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

func TreeMax(root *Node, max int) int {
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

func BreadthFirstTraversal(root *Node) ([]int, error) {
	if root == nil {
		return []int{}, nil
	}

	// create queue and AddItems root node onto the queue.
	// we don't know how many nodes are going to be on the tree,
	// so we pick an arbitrary size for our queue.
	size := 10
	queue, err := NewCircularQueue[*Node](size)
	if err != nil {
		return []int{}, err
	}
	err = queue.AddItems(root)
	if err != nil {
		return []int{}, err
	}

	items := []int{}
	for !queue.IsEmpty() {
		currentNode, err := queue.GetFrontItem()
		if err != nil {
			return []int{}, err
		}

		items = append(items, currentNode.Data)

		if currentNode.Left != nil {
			err = queue.AddItems(currentNode.Left)
			if err != nil {
				return []int{}, err
			}
		}

		if currentNode.Right != nil {
			err = queue.AddItems(currentNode.Right)
			if err != nil {
				return []int{}, err
			}
		}
	}
	return items, nil
}
