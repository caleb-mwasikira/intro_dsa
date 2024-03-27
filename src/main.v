module main

import data_structures as ds

fn main() {
	numbers := [10, 20, 30]

	mut ll := ds.LinkedList{
		head: unsafe { nil }
	}
	println('Initial linked list:\n ${ll.repr()}\r\n')

	ll.insert_at_head(...numbers)!
	println('After head insertion:\n ${ll.repr()}\r\n')

	prime_numbers := [1, 2, 3]
	ll.insert_at_end(...prime_numbers)!
	println('After rear insertion:\n ${ll.repr()}\r\n')

	new_primes := [5, 7, 11]
	mid_index := 3
	ll.insert_at_middle(mid_index, ...new_primes)!
	println('After inserting middle at index ${mid_index}:\n ${ll.repr()}\r\n')

	last_node := ll.delete_last_node()!
	println('After deleting last node:\n ${ll.repr()}\r\n')
	println('Deleted node:\n ${last_node}')

	index := 2
	del_node := ll.delete_at_index(index)!
	println('After deleting node at index ${index}:\n ${ll.repr()}\r\n')
	println('Deleted node:\n ${del_node}')

	data := 42
	node := ll.search(data) or {
		eprintln(err)
		return
	}
	println('Node ${node} found')
}
