module data_structures

pub struct LinkedList {
mut:
	head &LinearNode = unsafe { nil }
	size int
}

pub fn LinkedList.from(_items ...int) !LinkedList {
	mut ll := LinkedList{}
	ll.insert_at_beginning(..._items)
	return ll
}

pub fn (mut ll LinkedList) insert_at_beginning(_items ...int) {
	mut items := _items.clone()

	if ll.is_empty() {
		first := items.first()
		items.delete(0)

		new_node := &LinearNode{
			data: first
		}
		ll.head = new_node
		ll.size++
	}

	mut current_node := ll.head

	for _, item in items {
		new_node := &LinearNode{
			data: item
		}
		current_node.next = new_node
		current_node = new_node
		ll.size++
	}
}

pub fn (mut ll LinkedList) insert_at_end(_items ...int) {
	mut items := _items.clone()

	if ll.is_empty() {
		ll.insert_at_beginning(..._items)
	} else {
		mut current_node := ll.head

		// traverse linked list until last node
		for current_node.next != unsafe { nil } {
			current_node = current_node.next
		}

		// begin rear insertion now that we are at the last node
		for _, item in items {
			new_node := &LinearNode{
				data: item
			}
			current_node.next = new_node
			current_node = new_node
			ll.size++
		}
	}
}

fn (mut ll LinkedList) node_at_index(index int) !&LinearNode {
	if !ll.is_valid_index(index) {
		return error('invalid index provided for middle insertion')
	}

	// traverse to the node at the index
	mut count := 0
	mut current_node := ll.head

	for current_node != unsafe { nil } {
		if index == count {
			break
		}

		current_node = current_node.next
		count++
	}

	return current_node
}

pub fn (mut ll LinkedList) insert_at_middle(index int, _items ...int) ! {
	if ll.is_empty() || ll.size == 1 {
		return error('linked list has no middle to insert to')
	}

	// (index - 1) gets node just before the index
	mut prev_node := ll.node_at_index(index - 1)!
	next_node := prev_node.next

	items := _items.clone()

	// add new nodes into a chain as prev_node's next
	for _, item in items {
		new_node := &LinearNode{
			data: item
		}
		prev_node.next = new_node
		prev_node = new_node
		ll.size++
	}

	// link back up to the old chain
	prev_node.next = next_node
}

pub fn (mut ll LinkedList) delete_head_node() ! {
	if ll.is_empty() {
		return error('cannot delete nodes from an empty linked list')
	}

	ll.head = ll.head.next
	ll.size--
}

pub fn (mut ll LinkedList) delete_tail_node() ! {
	if ll.is_empty() {
		return error('cannot delete nodes from an empty linked list')
	}

	// traverse till 2nd last node
	mut current_node := ll.head

	for current_node.next.next != unsafe { nil } {
		current_node = current_node.next
	}

	// delete last node by changing the 2nd last node's pointer to null
	current_node.next = unsafe { nil }
	ll.size--
}

pub fn (mut ll LinkedList) delete_from_middle(index int) ! {
	if ll.is_empty() || ll.size == 1 {
		return error('linked list has no middle to delete from')
	}

	// (index - 1) gets the node just before the index
	mut prev_node := ll.node_at_index(index - 1)!

	// delete node from chain
	prev_node.next = prev_node.next.next
	ll.size--
}

pub fn (mut ll LinkedList) delete_n_items_from_middle(index int, num_items int) ! {
	if ll.is_empty() || ll.size == 1 {
		return error('linked list has no middle to delete from')
	}

	// (index - 1) gets the node just before the index
	mut prev_node := ll.node_at_index(index - 1)!

	for _ in 0 .. num_items {
		// delete node from chain
		prev_node.next = prev_node.next.next
		ll.size--
	}
}

pub fn (mut ll LinkedList) delete_node(data int) !&LinearNode {
	if ll.is_empty() {
		return error('cannot delete_node from empty linked list')
	}

	// search for node with data provided
	mut current_node := ll.head
	mut prev_node := current_node

	for current_node != unsafe { nil } {
		if current_node.data == data {
			break
		}

		prev_node = current_node
		current_node = current_node.next
	}

	return if current_node == unsafe { nil } {
		// we did not find the node we were looking for
		error('node with data item: ${data} not found')
	} else {
		// we found the node we were looking for, delete it
		prev_node.next = prev_node.next.next
		ll.size--
		current_node
	}
}

pub fn (ll LinkedList) size() int {
	mut count := 0
	mut current_node := ll.head

	// traverse linked list counting the number of nodes
	for current_node != unsafe { nil } {
		current_node = current_node.next
		count++
	}
	return count
}

fn (ll LinkedList) is_valid_index(index int) bool {
	return index >= 0 && index < ll.size
}

fn (ll LinkedList) is_empty() bool {
	return ll.head == unsafe { nil }
}

pub fn (ll LinkedList) repr() string {
	mut repr_str := ''

	if ll.is_empty() {
		return '[head: ?]'
	}

	// traverse through each node on linked list
	mut current_node := ll.head

	for current_node != unsafe { nil } {
		repr_str += '${current_node.data}-> '
		current_node = current_node.next
	}

	return repr_str + '?'
}
