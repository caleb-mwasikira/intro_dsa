module data_structures

struct LinearNode {
	data int @[required]
mut:
	next &LinearNode = unsafe { nil }
}

