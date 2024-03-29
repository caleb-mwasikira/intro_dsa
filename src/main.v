module main

fn iterative_fib(n int) !int {
	if n < 0 {
		return error('expected n to be a positive non-zero integer')
	}

	mut prev2 := 0
	mut prev := 1
	mut result := 0

	if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	}

	for i := 2; i <= n; i++ {
		result = prev + prev2
		prev2 = prev
		prev = result
	}
	return result
}

fn recursive_fib(n int) !int {
	if n <= 0 {
		return 0
	} else if n <= 2 {
		return 1
	} else {
		return recursive_fib(n - 1)! + recursive_fib(n - 2)!
	}
}

fn main() {
	number := 8

	println("iterative fibonacci: ${iterative_fib(number)!}")
	println("recursive fibonacci: ${recursive_fib(number)!}")
}
