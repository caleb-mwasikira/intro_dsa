module main

import data_structures as ds

fn main() {
	prime_numbers := [1, 2, 3]

	mut ll := ds.LinkedList.from(...prime_numbers)!
	println('begin linked list: ${ll.repr()}\n')

	odd_numbers := [9, 15]
	ll.insert_at_end(...odd_numbers)
	println('after insert_at_end ${odd_numbers}:\n ${ll.repr()}\n')

	even_numbers := [4, 6, 8]
	ll.insert_at_middle(1, ...even_numbers)!
	println('after insert_at_middle ${even_numbers}:\n ${ll.repr()}\n')

	ll.delete_head_node()!
	println('after delete_head_node:\n ${ll.repr()}\n')

	ll.delete_tail_node()!
	println('after delete_tail_node:\n ${ll.repr()}\n')

	index := 2
	ll.delete_from_middle(index)!
	println('after deleting index ${index}:\n ${ll.repr()}\n')

	num_items := 2
	ll.delete_n_items_from_middle(index, num_items)!
	println('after deleting ${num_items} items starting at index ${index}:\n ${ll.repr()}\n')

	data := 6
	println('deleting node with data item ${data}...')
	ll.delete_node(data) or {
		eprintln(err)
		println(ll.repr())

		exit(1)
	}
	println(ll.repr())

	println("final size of linked list: ${ll.size()}\n")
}
