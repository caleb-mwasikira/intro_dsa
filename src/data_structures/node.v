module data_structures

// Node for linear data structures
struct LinearNode {
mut:
	data int    @[required]
	next &LinearNode = unsafe { nil }
}
