module data_structures

pub struct LinkedList {
mut:
	head &Node
}

pub fn LinkedList.from(_items ...int) !LinkedList {
	mut ll := LinkedList{
		head: unsafe { nil }
	}
	ll.insert_at_head(..._items)!
	return ll
}

pub fn (mut ll LinkedList) insert_at_head(_items ...int) ! {
	mut items := _items.clone()

	if items.len == 0 {
		return error('no items given for head insertion')
	}

	if ll.is_empty() {
		first_item := items.first()
		items.delete(0)

		new_node := Node{
			data: first_item
		}
		ll.head = &new_node
	}

	mut current_node := ll.head

	for _, item in items {
		mut new_node := &Node{
			data: item
		}

		current_node.next = new_node
		current_node = new_node
	}
}

pub fn (mut ll LinkedList) insert_at_middle(index int, _items ...int) ! {
	ll.validate_index(index)!

	if ll.is_empty() || index == 0 {
		ll.insert_at_head(..._items)!
		return
	}

	if _items.len == 0 {
		return error('no items given for middle insertion')
	}

	mut current_index := 0
	mut current_node := ll.head
	mut after_node := unsafe { nil }

	for current_node != unsafe { nil } {
		if current_index == index {
			after_node = current_node.next

			for _, item in _items {
				new_node := &Node{
					data: item
				}
				current_node.next = new_node
				current_node = new_node
			}

			current_node.next = after_node
		}

		current_node = current_node.next
		current_index++
	}
}

pub fn (mut ll LinkedList) insert_at_end(_items ...int) ! {
	if _items.len == 0 {
		return error('no items given for rear insertion')
	}

	if ll.is_empty() {
		ll.insert_at_head(..._items)!
		return
	}

	// traverse to the last node
	mut current_node := ll.head

	for current_node.next != unsafe { nil } {
		current_node = current_node.next
	}

	for _, item in _items {
		new_node := &Node{
			data: item
		}
		current_node.next = new_node
		current_node = new_node
	}
}

pub fn (ll LinkedList) delete_last_node() !&Node {
	if ll.is_empty() {
		return error('attempting to delete from empty linked list')
	}

	// traverse to second last node
	mut current_node := ll.head

	for current_node.next.next != unsafe { nil } {
		current_node = current_node.next
	}

	last_node := current_node.next

	// unlink last node from 2nd last node
	current_node.next = unsafe { nil }
	return last_node
}

pub fn (mut ll LinkedList) delete_at_index(index int) !&Node {
	ll.validate_index(index)!

	if ll.is_empty() {
		return error('attempting to delete from empty linked list')
	}

	mut del_node := unsafe { nil }

	if index == 0 {
		// deleting head of linked list
		del_node = ll.head
		ll.head = ll.head.next
		return del_node
	}

	mut current_node := ll.head
	mut current_index := 0

	for current_node != unsafe { nil } {
		// traverse to the node just before the one at index
		if current_index == index - 1 {
			del_node = current_node.next
			current_node.next = current_node.next.next
		}

		current_node = current_node.next
		current_index++
	}

	return del_node
}

pub fn (ll LinkedList) search(value int) !&Node {
	if ll.is_empty() {
		return error('attempting to search from empty linked list')
	}

	// traverse till end of linked list or until we find value
	mut current_node := ll.head

	for current_node != unsafe { nil } {
		if current_node.data == value {
			return current_node
		}

		current_node = current_node.next
	}

	return error('node with value ${value} not found')
}

fn (mut ll LinkedList) validate_index(index int) ! {
	ll_size := ll.size()

	if index < 0 || index >= ll_size {
		return error('invalid index ${index} for linked list. Acceptable Range is 0..${ll_size}')
	}
}

fn (ll LinkedList) is_empty() bool {
	return ll.head == unsafe { nil }
}

fn (mut ll LinkedList) size() int {
	if ll.is_empty() {
		return 0
	}

	mut count := 0
	mut current_node := ll.head

	for current_node != unsafe { nil } {
		count++
		current_node = current_node.next
	}

	return count
}

pub fn (ll LinkedList) repr() string {
	mut repr_str := ''
	mut current_node := ll.head

	for current_node != unsafe { nil } {
		node_data := current_node.data
		repr_str += '[${node_data}] ->'

		current_node = current_node.next
	}

	return repr_str + '?'
}
