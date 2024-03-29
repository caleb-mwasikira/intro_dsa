module data_structures

// Node for linear data structures
struct LinearNode {
mut:
	data int         @[required]
	next &LinearNode = unsafe { nil }
}

struct TreeNode {
	data int @[required]
mut:
	left  &TreeNode = unsafe { nil }
	right &TreeNode = unsafe { nil }
}
