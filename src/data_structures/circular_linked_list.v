module data_structures

pub struct CircularLinkedList {
mut:
	tail &LinearNode = unsafe { nil }
pub mut:
	size int
}

pub fn CircularLinkedList.from(_items ...int) !CircularLinkedList {
	mut items := _items.clone()

	if items.len == 0 {
		return error('failed to create circular linked list, len of items is zero')
	}

	mut cll := CircularLinkedList{}
	cll.insert_at_beginning(...items)

	return cll
}

pub fn (mut cll CircularLinkedList) insert_at_beginning(_items ...int) {
	mut items := _items.clone()

	if items.len == 0 {
		return
	}

	if cll.is_empty() {
		// get last element and set it as tail node
		last := items.last()
		items.delete_last()

		cll.tail = &LinearNode{
			data: last
		}
		cll.size++
	}

	// insert elements from head to tail
	mut current_node := cll.tail

	// unlink old chain
	mut old_chain := if current_node.next == unsafe { nil } {
		cll.tail
	} else {
		current_node.next
	}

	for _, item in items {
		new_node := &LinearNode{
			data: item
		}

		current_node.next = new_node
		current_node = new_node
		cll.size++
	}

	// link old chain to last newest node
	current_node.next = old_chain
}

// TODO: delete_node, delete_at_index(index), delete_n_items_at_index(index, num_items)

fn (cll CircularLinkedList) is_empty() bool {
	return cll.tail == unsafe { nil }
}

pub fn (cll CircularLinkedList) repr() string {
	mut repr_str := ''

	// traverse circular linked list from head to tail
	mut current_node := cll.tail.next

	for current_node != unsafe { nil } && current_node != cll.tail {
		repr_str += '${current_node.data}-> '
		current_node = current_node.next
	}

	return repr_str + '${current_node.data} ↩ '
}
