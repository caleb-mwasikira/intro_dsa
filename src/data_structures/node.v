module data_structures

struct Node {
mut:
	data int    @[required]
	next &Node = unsafe { nil }
}
