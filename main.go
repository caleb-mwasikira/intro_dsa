package main

import (
	"fmt"

	"github.com/caleb-mwasikira/intro_dsa/algos"
)

func main() {
	// iterativeFib(n)
	var n int = 6
	result := algos.IterativeFib(n)
	fmt.Printf("iterativeFib(%v) = %#v\n", n, result)

	// recursiveFib(n)
	result = algos.RecursiveFib(n)
	fmt.Printf("recursiveFib(%v) = %#v\n", n, result)

	// cachedRecursiveFib(n)
	result = algos.CachedRecursiveFib(n)
	fmt.Printf("cachedRecursiveFib(%v) = %#v\n", n, result)
}
