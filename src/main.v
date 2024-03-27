module main

import data_structures as ds

fn main() {
	numbers := [10, 20, 30]

	mut stack := ds.Stack.from(...numbers)
	for _ in 0 .. 5 {
		item := stack.pop() or {
			eprintln("${err}\r\n\r\n")
			break
		}
		println(item)
	}

	mut stack_2 := ds.Stack.new(numbers.len)!
	for _, num in numbers {
		stack_2.push(num)!
	}

	// adding more items to already full stack
	stack_2.push(42)!
}
