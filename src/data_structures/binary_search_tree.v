module data_structures

pub struct BinarySearchTree {
	head &TreeNode = unsafe { nil }
}

pub fn (mut bst BinarySearchTree) insert(mut root TreeNode, item int) &TreeNode {
	if root == unsafe { nil } {
		new_node := &TreeNode{
			data: item
		}
		return new_node
	}

	if item < root.data {
		// proceed to left subtree
		root.left = bst.insert(mut root.left, item)
	} else if item > root.data {
		// proceed to right subtree
		root.right = bst.insert(mut root.right, item)
	} else {
		// item == root.data, do nothin...
	}

	return unsafe { root }
}

// left-> root-> right
pub fn (bst BinarySearchTree) inorder(root &TreeNode) {
	if root != unsafe { nil } {
		// traverse left subtree
		bst.inorder(root.left)

		// traverse root
		print('${root.data}-> ')

		// traverse right subtree
		bst.inorder(root.right)
	}
}

// root-> left-> right
pub fn (bst BinarySearchTree) preorder(root &TreeNode) {
	if root != unsafe { nil } {
		// traverse root
		print('${root.data}-> ')

		// traverse left subtree
		bst.preorder(root.left)

		// traverse right subtree
		bst.preorder(root.right)
	}
}

// left-> right-> root
pub fn (bst BinarySearchTree) postorder(root &TreeNode) {
	if root != unsafe {nil} {
		// traverse left subtree
		bst.postorder(root.left)

		// traverse right subtree
		bst.postorder(root.right)

		// traverse root
		print("${root.data}-> ")
	}
}
