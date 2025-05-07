package main

import (
	"log"
	"os"
	"path/filepath"
	"slices"

	"github.com/caleb-mwasikira/intro_dsa/ds"
)

var (
	ignoredFiles []string = []string{".git"}
)

// WalkDir recursively builds the tree starting from rootPath
func WalkDir(path string) (*ds.FSNode, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	node := &ds.FSNode{
		Name: filepath.Base(path),
		Type: ds.Directory,
	}

	if !info.IsDir() {
		// Base case: it's a file
		node.Type = ds.File
		return node, nil
	}

	if slices.Contains(ignoredFiles, filepath.Base(path)) {
		// We are not going to list files, folders in ignoredFiles
		return node, nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		childPath := filepath.Join(path, entry.Name())
		childNode, err := WalkDir(childPath)
		if err != nil {
			return nil, err
		}
		node.Children = append(node.Children, childNode)
	}

	return node, nil
}

func main() {
	dir := "/home/jobava/Work/golang"
	rootNode, err := WalkDir(dir)
	if err != nil {
		log.Fatalf("error walking dir; %v", err)
	}

	ds.PrintTree(rootNode, "", true)
}
