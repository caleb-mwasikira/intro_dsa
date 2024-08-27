package tests

import (
	"fmt"
	"slices"
	"testing"

	"github.com/caleb-mwasikira/intro_dsa/algos"
)

type FibTest struct {
	n              int
	expectedResult []int
}

var (
	MAX_N int = 40
	N     int = 40

	tests []FibTest = []FibTest{
		{n: -1, expectedResult: []int{}},
		{n: 0, expectedResult: []int{0}},
		{n: 1, expectedResult: []int{0, 1}},
		{n: 6, expectedResult: []int{0, 1, 1, 2, 3, 5, 8}},
	}
)

/*
tests all our fib(n) functions for accuracy
*/
func TestFib(t *testing.T) {
	t.Run("test iterativeFib(n int) []int", func(t *testing.T) {
		testFibFunc(t, algos.IterativeFib)
	})

	t.Run("test recursiveFib(n int) []int", func(t *testing.T) {
		testFibFunc(t, algos.RecursiveFib)
	})

	t.Run("test cachedRecursiveFib(n int) []int", func(t *testing.T) {
		testFibFunc(t, algos.CachedRecursiveFib)
	})
}

func testFibFunc(t *testing.T, fn func(int) []int) {
	for _, test := range tests {
		result := fn(test.n)

		if !slices.Equal(result, test.expectedResult) {
			t.Errorf("expected result %#v from func iterativeFib(%v) but got %#v instead\n", test.expectedResult, test.n, result)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	fmt.Println()
	b.Run("bench iterativeFib(n int)[]int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			algos.IterativeFib(N)
		}
	})
	fmt.Println()

	if N > MAX_N {
		b.Logf("skipping bench recursiveFib(n int)[]int as it is too slow for values > %v\n\n", MAX_N)
	} else {
		b.Run("bench recursiveFib(n int)[]int", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				algos.RecursiveFib(N)
			}
		})
		fmt.Println()
	}

	b.Run("bench cachedRecursiveFib(n int)[]int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			algos.CachedRecursiveFib(N)
		}
	})
}
