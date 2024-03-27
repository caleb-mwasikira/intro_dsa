module data_structures

pub struct Stack {
	size int @[required]
mut:
	top   int = -1 // -1 means the stack is empty
	items []int
}

// Build a new stack with fixed size
pub fn Stack.new(size int) !Stack {
	if size < 0 {
		return error('cannot allocate stack of size less than zero')
	}

	println("Creating new stack of size ${size}...")
	return Stack{
		size: size
	}
}

// Build a stack from an existing array
pub fn Stack.from(items ...int) Stack {
	println("Creating stack from existing array: ${items}")
	new_top := items.len - 1

	return Stack{
		size: items.len
		items: items
		top: new_top
	}
}

pub fn (mut stack Stack) push(items ...int) ! {
	for _, item in items {
		if stack.is_full() {
			return error('cannot push any more items onto stack; stack already full')
		}

		println("Pushing ${item} into stack...")
		stack.items << item
		stack.top++
	}
}

pub fn (mut stack Stack) pop() !int {
	if stack.is_empty() {
		return error('cannot pop any more items out of stack; stack empty')
	}

	println("Popping last item from stack...")
	last := stack.items.last()
	stack.items.delete_last()
	stack.top--
	return last
}

fn (stack Stack) is_empty() bool {
	return stack.top == -1
}

fn (stack Stack) is_full() bool {
	return stack.top == (stack.size - 1)
}
