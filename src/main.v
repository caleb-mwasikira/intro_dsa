module main

import data_structures as ds

fn main() {
	prime_numbers := [1, 2, 3]
	mut cll := ds.CircularLinkedList.from(...prime_numbers) or {
		eprintln(err)
		exit(1)
	}

	println('initial: ${prime_numbers}\n ${cll.repr()}\n')
	println('size: ${cll.size}\n')

	odd_numbers := [11, 13, 15]
	cll.insert_at_beginning(...odd_numbers)
	println('after insert_at_beginning: ${odd_numbers}\n ${cll.repr()}')
	println('size: ${cll.size}\n')
}
