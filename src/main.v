module main

fn iterative_fib(n int) !i64 {
	if n < 0 {
		return error('expected n to be a positive non-zero integer')
	}

	mut prev2 := i64(0)
	mut prev := i64(1)
	mut result := i64(0)

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

fn recursive_fib(n int, mut memo map[int]i64) !i64 {
	mut result := i64(0)

	if n in memo {
		return memo[n]
	} else {
		if n == 0 || n == 1 {
			result = n
		} else {
			result = recursive_fib(n - 1, mut memo)! + recursive_fib(n - 2, mut memo)!
		}

		memo[n] = result
	}

	return result
}

fn recursive_trib(n int, mut trib_memo map[int]i64) i64 {
	mut result := i64(0)

	if n in trib_memo {
		result = trib_memo[n]
	} else {
		result = match n {
			0, 1 {
				i64(0)
			}
			2 {
				i64(1)
			}
			else {
				recursive_trib(n - 1, mut trib_memo) + recursive_trib(n - 2, mut trib_memo) +
					recursive_trib(n - 3, mut trib_memo)
			}
		}

		trib_memo[n] = result
	}

	return result
}

fn main() {
	number := 40

	println('iterative fibonacci: ${iterative_fib(number)!}')

	// recursive fibonacci is a very slow algorithm for large input values. O(2^n)
	// to optimise its performance, we employ a technique used in dynamic programming
	// called memoization - where we store already-computed duplicate subproblems into a
	// hashmap and retrieve them whenever needed
	mut memo := map[int]i64{}
	println('recursive fibonacci: ${recursive_fib(number, mut memo)!}')

	mut trib_memo := map[int]i64{}
	println('recursive tribonacci: ${recursive_trib(number, mut trib_memo)}')
}
