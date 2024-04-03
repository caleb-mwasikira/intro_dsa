module main

fn iterative_fib(number int) !i64 {
	if number < 0 {
		return error("cannot calculate fib of negative numbers")
	}

	if number == 0 || number == 1 {
		return number
	} else {
		mut prev := i64(1)
		mut prev2 := i64(0)
		mut result := i64(prev + prev2)

		for i := 2; i < number; i++ {
			prev2 = prev
			prev = result
			result = prev + prev2
		}

		return result
	}
}

fn recursive_fib(number int, mut memo map[int]i64) !i64 {
	if number < 0 {
		return error("cannot calculate fib of negative numbers")
	}

	if number in memo {
		return memo[number]
	}

	if number == 0 || number == 1 {
		return number
	} else {
		result := recursive_fib(number - 1, mut memo)! + recursive_fib(number - 2, mut memo)!
		memo[number] = result
		return result
	}
}

fn main() {
	index := 100
	// it_result := iterative_fib(index)!

	mut memo := map[int]i64{}
	rec_result := recursive_fib(index, mut memo)!

	// println("iterative fib(${index}) = ${it_result}")
	println("recursive fib(${index}) = ${rec_result}")
}
