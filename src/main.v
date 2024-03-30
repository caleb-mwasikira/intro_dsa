module main

// Answers the question: is it possible to build the amount from
// a set of numbers?
// eg: amount = 5, numbers = [1, 2, 3] returns true because the amount
// 5 can be built from the combinations of 2+3, 3+1+1, 2+1+2 or 1+1+1+1+1
fn sum_possible(amount int, numbers []u32, mut memo map[int]bool) bool {
	if amount == 0 {
		return true
	} else if amount < 0 {
		return false
	} else {
		if amount in memo {
			return true
		}

		for _, num in numbers {
			sub_amount := amount - num
			if sum_possible(sub_amount, numbers, mut memo) {
				memo[amount] = true
				return true
			}
		}

		memo[amount] = false
		return false
	}
}

fn main() {
	amount := 4
	numbers := [u32(1), 2, 3]
	mut memo := map[int]bool{}

	println('Is amount ${amount} sum_possible from the numbers ${numbers}? ${sum_possible(amount, numbers, mut memo)}')
}
