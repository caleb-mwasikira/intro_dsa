package ds

import (
	"fmt"
)

type FSNodeType int

const (
	File FSNodeType = iota
	Directory
)

type FSNode struct {
	Name     string
	Type     FSNodeType
	Children []*FSNode // only used if Type == Directory
}

func (n *FSNode) AddChild(child *FSNode) {
	if n.Type == Directory {
		n.Children = append(n.Children, child)
	}
}

func PrintTree(n *FSNode, prefix string, isLast bool) {
	var connector string
	if isLast {
		connector = "└── "
	} else {
		connector = "├── "
	}

	fmt.Println(prefix + connector + n.Name)

	// Determine the new prefix for children
	var newPrefix string
	if isLast {
		newPrefix = prefix + "    "
	} else {
		newPrefix = prefix + "│   "
	}

	for i, child := range n.Children {
		PrintTree(child, newPrefix, i == len(n.Children)-1)
	}
}
