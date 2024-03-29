module main

import data_structures as ds

fn main() {
	mut numbers := [50, 30, 20, 40, 70, 60, 80]

	mut bst := ds.BinarySearchTree{}
	mut root := unsafe { nil }

	// set BST's head
	first := numbers.first()
	numbers.delete(0)

	root = bst.insert(mut root, first)

	// insert the rest of the numbers
	for _, num in numbers {
		bst.insert(mut root, num)
	}

	// list all values on BST using different traversal algorithms
	bst.inorder(root)
	println("\n")

	bst.preorder(root)
	println("\n")

	bst.postorder(root)
	println("\n")
}
